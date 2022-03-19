package application

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"
)

type UserApp struct {
	userRepo repository.UserRepository
}

var _ UserAppInterface = &UserApp{}

func NewUserApp(userRepo repository.UserRepository) *UserApp {
	return &UserApp{
		userRepo: userRepo,
	}
}

type UserAppInterface interface {
	Create(user *entity.User, action string) (*entity.User, error)
	Get(username string) (*entity.User, error)
	List(conditions string) ([]entity.User, error)
	Update(username string, update map[string]interface{}, user string) (*entity.User, error)
}

func (userApp *UserApp) Create(user *entity.User, action string) (*entity.User, error) {
	return userApp.userRepo.Create(user, action)
}

func (userApp *UserApp) Get(username string) (*entity.User, error) {
	return userApp.userRepo.Get(username)
}

func (userApp *UserApp) List(conditions string) ([]entity.User, error) {
	return userApp.userRepo.List(conditions)
}

func (userApp *UserApp) Update(username string, update map[string]interface{}, user string) (*entity.User, error) {
	return userApp.userRepo.Update(username, update, user)
}
