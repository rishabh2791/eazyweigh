package persistance

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
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
	userRoleAccessRepo.DB.Where("user_role_id = ? AND table_name = ?", userRoleAccess.UserRoleID, userRoleAccess.TableName).Take(&existingUserRoleAccess)

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
		updationErr := userRoleAccessRepo.DB.Table(entity.UserRoleAccess{}.Tablename()).Where("user_role_id = ? AND table_name = ?", userRoleAccess.UserRoleID, userRoleAccess.TableName).Updates(userRoleAccess).Error
		if updationErr != nil {
			return nil, updationErr
		}
	}

	return userRoleAccess, nil
}

func (userRoleAccessRepo *UserRoleAccessRepo) List(userRoleID string) ([]entity.UserRoleAccess, error) {
	userRoleAccesses := []entity.UserRoleAccess{}

	getErr := userRoleAccessRepo.DB.Where("user_role_id = ?", userRoleID).Find(&userRoleAccesses).Error
	if getErr != nil {
		return nil, getErr
	}

	return userRoleAccesses, nil
}

func (userRoleAccessRepo *UserRoleAccessRepo) Update(userRoleID string, userRoleAccess *entity.UserRoleAccess) (*entity.UserRoleAccess, error) {
	existingRoleAccess := entity.UserRoleAccess{}

	getErr := userRoleAccessRepo.DB.Where("user_role_id = ?", userRoleID).Take(&existingRoleAccess).Error
	if getErr != nil {
		return nil, getErr
	}

	updationErr := userRoleAccessRepo.DB.Table(entity.UserRoleAccess{}.Tablename()).Where("user_role_id = ?", userRoleID).Updates(userRoleAccess).Error
	if updationErr != nil {
		return nil, updationErr
	}

	updated := entity.UserRoleAccess{}
	userRoleAccessRepo.DB.Where("user_role_id = ?", userRoleID).Take(&updated)

	return &updated, nil
}
