package application

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"
)

type UserRoleApp struct {
	userRoleRepo repository.UserRoleRepository
}

var _ UserRoleAppInterface = &UserRoleApp{}

func NewUserRoleApp(userRoleRepo repository.UserRoleRepository) *UserRoleApp {
	return &UserRoleApp{
		userRoleRepo: userRoleRepo,
	}
}

type UserRoleAppInterface interface {
	Create(userRole *entity.UserRole) (*entity.UserRole, error)
	Get(userRole string) (*entity.UserRole, error)
	List(conditions string) ([]entity.UserRole, error)
	Update(userRole string, update *entity.UserRole) (*entity.UserRole, error)
}

func (userRoleApp *UserRoleApp) Create(userRole *entity.UserRole) (*entity.UserRole, error) {
	return userRoleApp.userRoleRepo.Create(userRole)
}

func (userRoleApp *UserRoleApp) Get(userRole string) (*entity.UserRole, error) {
	return userRoleApp.userRoleRepo.Get(userRole)
}

func (userRoleApp *UserRoleApp) List(conditions string) ([]entity.UserRole, error) {
	return userRoleApp.userRoleRepo.List(conditions)
}

func (userRoleApp *UserRoleApp) Update(userRole string, update *entity.UserRole) (*entity.UserRole, error) {
	return userRoleApp.userRoleRepo.Update(userRole, update)
}
