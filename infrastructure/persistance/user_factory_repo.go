package persistance

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserFactoryRepo struct {
	DB     *gorm.DB
	Logger hclog.Logger
}

var _ repository.UserFactoryRepository = &UserFactoryRepo{}

func NewUserFactoryRepo(db *gorm.DB, logger hclog.Logger) *UserFactoryRepo {
	return &UserFactoryRepo{
		DB:     db,
		Logger: logger,
	}
}

func (userFactoryRepo *UserFactoryRepo) Create(userFactory *entity.UserFactory) (*entity.UserFactory, error) {
	validationErr := userFactory.Validate()
	if validationErr != nil {
		return nil, validationErr
	}

	creationErr := userFactoryRepo.DB.Create(&userFactory).Error
	if creationErr != nil {
		return nil, creationErr
	}
	return userFactory, nil
}

func (userFactoryRepo *UserFactoryRepo) Get(conditions string) ([]entity.UserFactory, error) {
	factoryUsers := []entity.UserFactory{}
	getErr := userFactoryRepo.DB.
		Preload("User.UserRole").
		Preload("Factory.Address").
		Preload("Factory.CreatedBy").
		Preload("Factory.CreatedBy.UserRole").
		Preload("Factory.UpdatedBy").
		Preload("Factory.UpdatedBy.UserRole").
		Preload(clause.Associations).Where(conditions).Find(&factoryUsers).Error
	if getErr != nil {
		return nil, getErr
	}

	return factoryUsers, nil
}
