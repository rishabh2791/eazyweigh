package middlewares

import (
	"eazyweigh/application"
	"eazyweigh/domain/entity"
	"eazyweigh/domain/value_objects"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-hclog"
)

type AuthMiddleware struct {
	Logger   hclog.Logger
	AppStore application.AppStore
}

func NewAuthMiddleware(logger hclog.Logger, appStore application.AppStore) *AuthMiddleware {
	return &AuthMiddleware{
		Logger:   logger,
		AppStore: appStore,
	}
}

func (auth *AuthMiddleware) ValidateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response := value_objects.Response{}
		ctx.Header("Content-Type", "application/json")

		user := entity.User{}
		err := json.NewDecoder(ctx.Request.Body).Decode(&user)
		if err != nil {
			response.Status = false
			response.Message = err.Error()
			response.Payload = nil
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		errs := user.Validate("login")
		if errs != nil {
			response.Status = false
			response.Message = errs.Error()
			response.Payload = errs
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		ctx.Set("user", &user)
		ctx.Next()
	}
}

func (auth *AuthMiddleware) ValidateAccessToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response := value_objects.Response{}

		token, err := extractToken(ctx.Request)
		if err != nil {
			response.Status = false
			response.Message = err.Error()
			response.Payload = "Token Extraction Error."
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		accessDetails, err := auth.AppStore.AuthApp.ValidateAccessToken(token)
		if err != nil {
			response.Status = false
			response.Message = err.Error()
			response.Payload = "Unable to Validate Access Token."
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		uuid, err := auth.AppStore.AuthApp.FetchAuth(accessDetails.AccessUUID)
		if err != nil {
			response.Status = false
			response.Message = err.Error()
			response.Payload = "Token Expired."
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		if accessDetails.Username != uuid {
			response.Status = false
			response.Message = "Invalid Token."
			response.Payload = "Invalid Token."
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		user, err := auth.AppStore.UserApp.Get(accessDetails.Username)
		if err != nil {
			response.Status = false
			response.Message = err.Error()
			response.Payload = "Unable to get User Details from Database."
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		if user == nil {
			response.Status = false
			response.Message = "Invalid User."
			response.Payload = "User Does Not Exist."
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		if !user.Active {
			response.Status = false
			response.Message = "User Inactive."
			response.Payload = "User Inactive."
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		ctx.Set("user", user)
		ctx.Set("access_details", accessDetails)
		ctx.Next()
	}
}

func (auth *AuthMiddleware) ValidateRefreshToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response := value_objects.Response{}
		ctx.Header("Content-Type", "application/json")

		token, err := extractToken(ctx.Request)
		if err != nil {
			response.Status = false
			response.Message = err.Error()
			response.Payload = nil
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		refreshDetails, err := auth.AppStore.AuthApp.ValidateRefreshToken(token)
		if err != nil {
			response.Status = false
			response.Message = err.Error()
			response.Payload = nil
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		user, err := auth.AppStore.UserApp.Get(refreshDetails.Username)
		if err != nil {
			response.Status = false
			response.Message = err.Error()
			response.Payload = nil
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		if !user.Active {
			response.Status = false
			response.Message = "User Inactive."
			response.Payload = nil
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		actualCustomKey := auth.AppStore.AuthApp.GenerateCustomKey(user.Username, user.SecretKey)
		if refreshDetails.CustomKey != actualCustomKey {
			response.Status = false
			response.Message = "Authentication Failed, Invalid Token."
			response.Payload = nil
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		ctx.Set("user", user)
		ctx.Set("refresh_details", refreshDetails)
		ctx.Next()
	}
}

func extractToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	authHeaderContent := strings.Split(authHeader, " ")
	if len(authHeaderContent) != 2 {
		return "", errors.New("Token not provided or malformed")
	}
	return authHeaderContent[1], nil
}
