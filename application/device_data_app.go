package application

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"
)

type DeviceDataApp struct {
	devicedataRepository repository.DeviceDataRepository
}

var _ DeviceDataAppInterface = &DeviceDataApp{}

func NewDeviceDataApp(devcieDataRepository repository.DeviceDataRepository) *DeviceDataApp {
	return &DeviceDataApp{
		devicedataRepository: devcieDataRepository,
	}
}

func (deviceDataApp *DeviceDataApp) Create(device *entity.DeviceData) (*entity.DeviceData, error) {
	return deviceDataApp.devicedataRepository.Create(device)
}

func (deviceDataApp *DeviceDataApp) GetForDevice(deviceID string, conditions string) ([]entity.DeviceData, error) {
	return deviceDataApp.devicedataRepository.GetForDevice(deviceID, conditions)
}

type DeviceDataAppInterface interface {
	Create(device *entity.DeviceData) (*entity.DeviceData, error)
	GetForDevice(deviceID string, conditions string) ([]entity.DeviceData, error)
}
