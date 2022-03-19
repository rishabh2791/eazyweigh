package application

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"
	"eazyweigh/domain/value_objects"
)

type AuthApp struct {
	AuthRepo repository.AuthRepository
}

var _ AuthAppInterface = &AuthApp{}

func NewAuthApp(authRepo repository.AuthRepository) *AuthApp {
	return &AuthApp{
		AuthRepo: authRepo,
	}
}

type AuthAppInterface interface {
	Authenticate(reqUser *entity.User, user *entity.User) error
	GenerateTokens(user *entity.User) (*value_objects.Token, error)
	GenerateAuth(token *value_objects.Token) error
	GenerateCustomKey(username string, secretKey string) string
	ValidateAccessToken(token string) (*value_objects.AccessDetail, error)
	FetchAuth(uuid string) (string, error)
	DeleteAuth(uuid string) (int64, error)
	ValidateRefreshToken(token string) (*value_objects.RefreshDetail, error)
}

func (authApp *AuthApp) Authenticate(reqUser *entity.User, user *entity.User) error {
	return authApp.AuthRepo.Authenticate(reqUser, user)
}

func (authApp *AuthApp) GenerateTokens(user *entity.User) (*value_objects.Token, error) {
	return authApp.AuthRepo.GenerateTokens(user)
}

func (authApp *AuthApp) GenerateAuth(token *value_objects.Token) error {
	return authApp.AuthRepo.GenerateAuth(token)
}

func (authApp *AuthApp) GenerateCustomKey(username string, secretKey string) string {
	return authApp.AuthRepo.GenerateCustomKey(username, secretKey)
}

func (authApp *AuthApp) ValidateAccessToken(token string) (*value_objects.AccessDetail, error) {
	return authApp.AuthRepo.ValidateAccessToken(token)
}

func (authApp *AuthApp) FetchAuth(uuid string) (string, error) {
	return authApp.AuthRepo.FetchAuth(uuid)
}

func (authApp *AuthApp) DeleteAuth(uuid string) (int64, error) {
	return authApp.AuthRepo.DeleteAuth(uuid)
}

func (authApp *AuthApp) ValidateRefreshToken(token string) (*value_objects.RefreshDetail, error) {
	return authApp.AuthRepo.ValidateRefreshToken(token)
}
