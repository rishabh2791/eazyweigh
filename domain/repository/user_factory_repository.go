package repository

import "eazyweigh/domain/entity"

type UserFactoryRepository interface {
	Create(userFactory *entity.UserFactory) (*entity.UserFactory, error)
	Get(conditions string) ([]entity.UserFactory, error)
}
