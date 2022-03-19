package persistance

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRoleRepo struct {
	DB     *gorm.DB
	Logger hclog.Logger
}

var _ repository.UserRoleRepository = &UserRoleRepo{}

func NewUserRoleRepo(db *gorm.DB, logger hclog.Logger) *UserRoleRepo {
	return &UserRoleRepo{
		DB:     db,
		Logger: logger,
	}
}

func (userRoleRepo *UserRoleRepo) Create(userRole *entity.UserRole) (*entity.UserRole, error) {
	validationErrors := userRole.Validate()
	if validationErrors != nil {
		return nil, validationErrors
	}
	creationError := userRoleRepo.DB.Create(userRole).Error
	if creationError != nil {
		return nil, creationError
	}
	return userRole, nil
}

func (userRoleRepo *UserRoleRepo) Get(userRole string) (*entity.UserRole, error) {
	role := entity.UserRole{}
	getErr := userRoleRepo.DB.Preload(clause.Associations).Take(&role).Error
	if getErr != nil {
		return nil, getErr
	}
	return &role, nil
}

func (userRoleRepo *UserRoleRepo) List(conditions string) ([]entity.UserRole, error) {
	roles := []entity.UserRole{}
	getErr := userRoleRepo.DB.Preload(clause.Associations).Where(conditions).Find(&roles).Error
	if getErr != nil {
		return nil, getErr
	}
	return roles, nil
}

func (userRoleRepo *UserRoleRepo) Update(userRole string, update *entity.UserRole) (*entity.UserRole, error) {
	existingRole := entity.UserRole{}

	err := userRoleRepo.DB.Where("role = ?", userRole).Take(&existingRole).Error
	if err != nil {
		return nil, err
	}

	updationErr := userRoleRepo.DB.Table(entity.UserRole{}.Tablename()).Where("role = ?", userRole).Updates(update).Error
	if updationErr != nil {
		return nil, updationErr
	}

	updated := entity.UserRole{}
	userRoleRepo.DB.Where("role = ?", userRole).Take(&updated)

	return &updated, nil
}
