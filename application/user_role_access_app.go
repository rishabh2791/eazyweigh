package application

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"
)

type UserRoleAccessApp struct {
	userRoleAccessRepository repository.UserRoleAccessRepository
}

var _ UserRoleAccessAppInterface = &UserRoleAccessApp{}

func NewUserRoleAccessApp(userRoleAccessRepository repository.UserRoleAccessRepository) *UserRoleAccessApp {
	return &UserRoleAccessApp{
		userRoleAccessRepository: userRoleAccessRepository,
	}
}

type UserRoleAccessAppInterface interface {
	Create(userRoleAccess *entity.UserRoleAccess) (*entity.UserRoleAccess, error)
	List(userRole string) ([]entity.UserRoleAccess, error)
	Update(userRole string, userRoleAccess *entity.UserRoleAccess) (*entity.UserRoleAccess, error)
}

func (userRoleAccessApp *UserRoleAccessApp) Create(userRoleAccess *entity.UserRoleAccess) (*entity.UserRoleAccess, error) {
	return userRoleAccessApp.userRoleAccessRepository.Create(userRoleAccess)
}

func (userRoleAccessApp *UserRoleAccessApp) List(userRole string) ([]entity.UserRoleAccess, error) {
	return userRoleAccessApp.userRoleAccessRepository.List(userRole)
}

func (userRoleAccessApp *UserRoleAccessApp) Update(userRole string, userRoleAccess *entity.UserRoleAccess) (*entity.UserRoleAccess, error) {
	return userRoleAccessApp.userRoleAccessRepository.Update(userRole, userRoleAccess)
}
