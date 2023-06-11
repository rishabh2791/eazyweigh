package persistance

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DeviceDataRepo struct {
	DB     *gorm.DB
	Logger hclog.Logger
}

var _ repository.DeviceDataRepository = &DeviceDataRepo{}

func NewDeviceDataRepo(db *gorm.DB, logger hclog.Logger) *DeviceDataRepo {
	return &DeviceDataRepo{
		DB:     db,
		Logger: logger,
	}
}

func (deviceDataRepo *DeviceDataRepo) Create(device *entity.DeviceData) (*entity.DeviceData, error) {
	validationErr := device.Validate()

	if validationErr != nil {
		return nil, validationErr
	}

	creationErr := deviceDataRepo.DB.Create(&device).Error
	return device, creationErr
}

func (deviceDataRepo *DeviceDataRepo) Get(id string) (*entity.DeviceData, error) {
	deviceData := entity.DeviceData{}

	getErr := deviceDataRepo.DB.
		Preload(clause.Associations).Where("id = ?", id).Take(&deviceData).Error

	return &deviceData, getErr
}

func (deviceDataRepo *DeviceDataRepo) List(conditions string) ([]entity.DeviceData, error) {
	deviceData := []entity.DeviceData{}

	getErr := deviceDataRepo.DB.
		Preload(clause.Associations).Where(conditions).Find(&deviceData).Error
	return deviceData, getErr
}
