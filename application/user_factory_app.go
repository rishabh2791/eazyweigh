package application

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"
)

type UserFactoryApp struct {
	userFactoryRepository repository.UserFactoryRepository
}

var _ UserFactoryAppInterface = &UserFactoryApp{}

func NewUserFactoryApp(userFactoryRepository repository.UserFactoryRepository) *UserFactoryApp {
	return &UserFactoryApp{
		userFactoryRepository: userFactoryRepository,
	}
}

type UserFactoryAppInterface interface {
	Create(userFactory *entity.UserFactory) (*entity.UserFactory, error)
	Get(factoryID string) ([]entity.UserFactory, error)
}

func (userFactoryApp *UserFactoryApp) Create(userFactory *entity.UserFactory) (*entity.UserFactory, error) {
	return userFactoryApp.userFactoryRepository.Create(userFactory)
}

func (userFactoryApp *UserFactoryApp) Get(factoryID string) ([]entity.UserFactory, error) {
	return userFactoryApp.userFactoryRepository.Get(factoryID)
}
