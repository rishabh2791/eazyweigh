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

func (deviceDataApp *DeviceDataApp) Create(deviceData *entity.DeviceData) (*entity.DeviceData, error) {
	return deviceDataApp.devicedataRepository.Create(deviceData)
}

func (deviceDataApp *DeviceDataApp) Get(id string) (*entity.DeviceData, error) {
	return deviceDataApp.devicedataRepository.Get(id)
}

func (deviceDataApp *DeviceDataApp) List(conditions string) ([]entity.DeviceData, error) {
	return deviceDataApp.devicedataRepository.List(conditions)
}

type DeviceDataAppInterface interface {
	Create(*entity.DeviceData) (*entity.DeviceData, error)
	Get(string) (*entity.DeviceData, error)
	List(string) ([]entity.DeviceData, error)
}
