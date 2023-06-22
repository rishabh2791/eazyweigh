package persistance

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DeviceTypeRepo struct {
	DB     *gorm.DB
	Logger hclog.Logger
}

func NewDeviceTypeRepo(db *gorm.DB, logger hclog.Logger) repository.DeviceTypeRepository {
	return &DeviceTypeRepo{
		DB:     db,
		Logger: logger,
	}
}

func (deviceTypeRepo *DeviceTypeRepo) Create(deviceType *entity.DeviceType) (*entity.DeviceType, error) {
	validationErr := deviceType.Validate()
	if validationErr != nil {
		return nil, validationErr
	}

	creationErr := deviceTypeRepo.DB.Create(&deviceType).Error
	return deviceType, creationErr
}

func (deviceTypeRepo *DeviceTypeRepo) Get(id string) (*entity.DeviceType, error) {
	deviceType := entity.DeviceType{}

	getErr := deviceTypeRepo.DB.
		Preload("Factory.Address").
		Preload("Factory.CreatedBy").
		Preload("Factory.CreatedBy.UserRole").
		Preload("Factory.UpdatedBy").
		Preload("Factory.UpdatedBy.UserRole").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).Where("id = ?", id).First(&deviceType).Error

	return &deviceType, getErr
}

func (deviceTypeRepo *DeviceTypeRepo) List(conditions string) ([]entity.DeviceType, error) {
	deviceTypes := []entity.DeviceType{}

	getErr := deviceTypeRepo.DB.
		Preload("Factory.Address").
		Preload("Factory.CreatedBy").
		Preload("Factory.CreatedBy.UserRole").
		Preload("Factory.UpdatedBy").
		Preload("Factory.UpdatedBy.UserRole").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).Where(conditions).Find(&deviceTypes).Error

	return deviceTypes, getErr
}

func (deviceTypeRepo *DeviceTypeRepo) Update(id string, deviceType *entity.DeviceType) (*entity.DeviceType, error) {
	existingDeviceType := entity.DeviceType{}

	getErr := deviceTypeRepo.DB.
		Preload(clause.Associations).Where("id = ?", id).First(&existingDeviceType).Error
	if getErr != nil {
		return nil, getErr
	}

	updationErr := deviceTypeRepo.DB.Table(deviceType.Tablename()).Where("id = ?", id).Updates(&deviceType).Error
	if updationErr != nil {
		return nil, updationErr
	}

	updated := entity.DeviceType{}
	deviceTypeRepo.DB.
		Preload("Factory.Address").
		Preload("Factory.CreatedBy").
		Preload("Factory.CreatedBy.UserRole").
		Preload("Factory.UpdatedBy").
		Preload("Factory.UpdatedBy.UserRole").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).Where("id = ?", id).First(&updated)

	return &updated, nil
}
