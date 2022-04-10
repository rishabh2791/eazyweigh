package persistance

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRoleAccessRepo struct {
	DB     *gorm.DB
	Logger hclog.Logger
}

var _ repository.UserRoleAccessRepository = &UserRoleAccessRepo{}

func NewUserRoleAccessRepo(db *gorm.DB, logger hclog.Logger) *UserRoleAccessRepo {
	return &UserRoleAccessRepo{
		DB:     db,
		Logger: logger,
	}
}

func (userRoleAccessRepo *UserRoleAccessRepo) Create(userRoleAccess *entity.UserRoleAccess) (*entity.UserRoleAccess, error) {
	existingUserRoleAccess := []entity.UserRoleAccess{}
	userRoleAccessRepo.DB.Preload(clause.Associations).Where("user_role_role = ? AND table_name = ?", userRoleAccess.UserRoleRole, userRoleAccess.TableName).Take(&existingUserRoleAccess)

	if len(existingUserRoleAccess) == 0 {
		validationErr := userRoleAccess.Validate()
		if validationErr != nil {
			return nil, validationErr
		}

		creationErr := userRoleAccessRepo.DB.Create(&userRoleAccess).Error
		if creationErr != nil {
			return nil, creationErr
		}
	} else {
		updationErr := userRoleAccessRepo.DB.Table(entity.UserRoleAccess{}.Tablename()).Where("user_role_role = ? AND table_name = ?", userRoleAccess.UserRoleRole, userRoleAccess.TableName).Updates(userRoleAccess).Error
		if updationErr != nil {
			return nil, updationErr
		}
	}

	return userRoleAccess, nil
}

func (userRoleAccessRepo *UserRoleAccessRepo) List(userRole string) ([]entity.UserRoleAccess, error) {
	userRoleAccesses := []entity.UserRoleAccess{}

	getErr := userRoleAccessRepo.DB.
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).Where("user_role_role = ?", userRole).Find(&userRoleAccesses).Error
	if getErr != nil {
		return nil, getErr
	}

	return userRoleAccesses, nil
}

func (userRoleAccessRepo *UserRoleAccessRepo) Update(userRole string, userRoleAccess *entity.UserRoleAccess) (*entity.UserRoleAccess, error) {
	existingRoleAccess := entity.UserRoleAccess{}

	getErr := userRoleAccessRepo.DB.Preload(clause.Associations).Where("user_role_role = ?", userRole).Take(&existingRoleAccess).Error
	if getErr != nil {
		return nil, getErr
	}

	updationErr := userRoleAccessRepo.DB.Table(entity.UserRoleAccess{}.Tablename()).Where("user_role_role = ?", userRole).Updates(userRoleAccess).Error
	if updationErr != nil {
		return nil, updationErr
	}

	updated := entity.UserRoleAccess{}
	userRoleAccessRepo.DB.
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).Take(&updated)

	return &updated, nil
}
