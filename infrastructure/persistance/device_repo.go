package persistance

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
)

type DeviceRepo struct {
	DB     *gorm.DB
	Logger hclog.Logger
}

var _ repository.DeviceRepository = &DeviceRepo{}

func NewDeviceRepo(db *gorm.DB, logger hclog.Logger) *DeviceRepo {
	return &DeviceRepo{
		DB:     db,
		Logger: logger,
	}
}

func (deviceRepo *DeviceRepo) Create(device *entity.Device) (*entity.Device, error) {
	validationErr := device.Validate()
	if validationErr != nil {
		return nil, validationErr
	}

	creationErr := deviceRepo.DB.Create(&device).Error
	return device, creationErr
}

func (deviceRepo *DeviceRepo) Get(id string) (*entity.Device, error) {
	device := entity.Device{}

	getErr := deviceRepo.DB.Where("id = ?", id).First(&device).Error
	return &device, getErr
}

func (deviceRepo *DeviceRepo) List(conditions string) ([]entity.Device, error) {
	devices := []entity.Device{}

	getErr := deviceRepo.DB.Where(conditions).Find(&devices).Error
	return devices, getErr
}

func (deviceRepo *DeviceRepo) Update(id string, device *entity.Device) (*entity.Device, error) {
	existingDevice := entity.Device{}

	getErr := deviceRepo.DB.Where("id = ?", id).First(&existingDevice).Error
	if getErr != nil {
		return nil, getErr
	}

	updationErr := deviceRepo.DB.Table(device.Tablename()).Where("id = ?", id).Updates(&device).Error
	if updationErr != nil {
		return nil, updationErr
	}

	updated := entity.Device{}
	deviceRepo.DB.Where("id = ?", id).First(&updated)

	return &updated, nil
}
