package persistance

import (
	"crypto/hmac"
	"crypto/sha256"
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/config"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"github.com/hashicorp/go-hclog"
	"golang.org/x/crypto/bcrypt"
)

type AuthRepo struct {
	RedisClient *redis.Client
	Logger      hclog.Logger
	Configs     *config.ServerConfig
}

var _ repository.AuthRepository = &AuthRepo{}

func NewAuthRepo(logger hclog.Logger, conf *config.ServerConfig, redisClient *redis.Client) *AuthRepo {
	return &AuthRepo{
		RedisClient: redisClient,
		Logger:      logger,
		Configs:     conf,
	}
}

func (auth *AuthRepo) Authenticate(reqUser *entity.User, user *entity.User) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqUser.Password))
}

func (auth *AuthRepo) GenerateTokens(user *entity.User) (*value_objects.Token, error) {
	td := value_objects.Token{}
	tokenData := auth.Configs.GetTokenConfig()
	keyData := auth.Configs.GetKeyConfig()
	td.ATExpires = time.Now().Add(time.Minute * time.Duration(tokenData.JWTExpiration)).Unix()
	td.AccessUUID = uuid.New().String()
	td.ATDuration = tokenData.JWTExpiration * 60
	td.RTExpires = time.Now().Add(time.Hour * 24 * time.Duration(tokenData.RefreshExpiration)).Unix()
	td.RefreshUUID = uuid.New().String()
	td.RTDuration = tokenData.RefreshExpiration * 24 * 60 * 60

	accessClaims := jwt.MapClaims{}
	accessClaims["username"] = user.Username
	accessClaims["token_type"] = "access"
	accessClaims["authorized"] = true
	accessClaims["duration"] = tokenData.JWTExpiration
	accessClaims["access_uuid"] = td.AccessUUID
	accessClaims["expiry"] = td.ATExpires

	atPrivateKey, err := ioutil.ReadFile(keyData.AccessTokenPrivateKeyPath)
	if err != nil {
		auth.Logger.Error("Unable to get Private Key String for Access Token.")
		return nil, err
	}

	atSigningKey, err := jwt.ParseRSAPrivateKeyFromPEM(atPrivateKey)
	if err != nil {
		auth.Logger.Error("Unable to get Signing Key String from Private Key for Access Token.")
		return nil, err
	}

	refreshClaims := jwt.MapClaims{}
	refreshClaims["username"] = user.Username
	refreshClaims["token_type"] = "refresh"
	refreshClaims["refresh_uuid"] = td.RefreshUUID
	refreshClaims["expiry"] = td.RTExpires
	refreshClaims["custom_key"] = auth.GenerateCustomKey(user.Username, user.SecretKey)

	rtPrivateKey, err := ioutil.ReadFile(keyData.RefreshTokenPrivateKeyPath)
	if err != nil {
		auth.Logger.Error("Unable to get Private Key String for Refresh Token.")
		return nil, err
	}

	rtSigningKey, err := jwt.ParseRSAPrivateKeyFromPEM(rtPrivateKey)
	if err != nil {
		auth.Logger.Error("Unable to get Signing Key String from Private String for Refresh Token.")
		return nil, err
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodRS256, accessClaims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodRS256, refreshClaims)
	td.Username = user.Username

	td.AccessToken, err = accessToken.SignedString(atSigningKey)
	if err != nil {
		auth.Logger.Error("Unable to sign Access Token.")
		return nil, err
	}
	td.RefreshToken, err = refreshToken.SignedString(rtSigningKey)
	if err != nil {
		auth.Logger.Error("Unable to sign Refresh Token.")
		return nil, err
	}

	return &td, nil
}

func (auth *AuthRepo) ValidateAccessToken(tokenString string) (*value_objects.AccessDetail, error) {
	accessDetails := value_objects.AccessDetail{}
	keyData := auth.Configs.GetKeyConfig()
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("Unexpected Signing Method.")
		}

		publicKey, err := ioutil.ReadFile(keyData.AccessTokenPublicKeyPath)
		if err != nil {
			return nil, err
		}

		verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
		if err != nil {
			return nil, err
		}
		return verifyKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid || claims["username"].(string) == "" || claims["token_type"].(string) != "access" {
		return nil, errors.New("Invalid Token. Authentication Failed.")
	}

	accessDetails.AccessUUID = claims["access_uuid"].(string)
	accessDetails.Username = claims["username"].(string)
	accessDetails.Duration = int(claims["duration"].(float64))

	return &accessDetails, nil
}

func (auth *AuthRepo) ValidateRefreshToken(tokenString string) (*value_objects.RefreshDetail, error) {
	refreshDetails := value_objects.RefreshDetail{}
	keyData := auth.Configs.GetKeyConfig()
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("Unexpected Signing Method.")
		}

		publicKey, err := ioutil.ReadFile(keyData.RefreshTokenPublicKeyPath)
		if err != nil {
			return nil, err
		}

		verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
		if err != nil {
			return nil, err
		}

		return verifyKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid || claims["username"].(string) == "" || claims["token_type"].(string) != "refresh" {
		return nil, errors.New("")
	}

	refreshDetails.RefreshUUID = claims["refresh_uuid"].(string)
	refreshDetails.Username = claims["username"].(string)
	refreshDetails.CustomKey = claims["custom_key"].(string)

	return &refreshDetails, nil
}

func (auth *AuthRepo) GenerateAuth(token *value_objects.Token) error {
	at := time.Unix(token.ATExpires, 0)
	rt := time.Unix(token.RTExpires, 0)
	now := time.Now()
	accessErr := auth.RedisClient.Set(token.AccessUUID, token.Username, at.Sub(now)).Err()
	if accessErr != nil {
		return accessErr
	}
	refreshErr := auth.RedisClient.Set(token.RefreshUUID, token.Username, rt.Sub(now)).Err()
	if refreshErr != nil {
		return refreshErr
	}
	return nil
}

func (auth *AuthRepo) GenerateCustomKey(username string, secretKey string) string {
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(username))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}

func (auth *AuthRepo) FetchAuth(uuid string) (string, error) {
	username, err := auth.RedisClient.Get(uuid).Result()
	if err != nil {
		return "", err
	}
	return username, nil
}

func (auth *AuthRepo) DeleteAuth(uuid string) (int64, error) {
	deleted, err := auth.RedisClient.Del(uuid).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}
