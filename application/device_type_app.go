package application

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"
)

type DeviceTypeApp struct {
	deviceTypeRepository repository.DeviceTypeRepository
}

var _ DeviceTypeAppInterface = &DeviceTypeApp{}

func NewDeviceTypeApp(deviceTypeRepository repository.DeviceTypeRepository) *DeviceTypeApp {
	deviceTypeApp := DeviceTypeApp{}
	deviceTypeApp.deviceTypeRepository = deviceTypeRepository
	return &deviceTypeApp
}

func (deviceTypeApp *DeviceTypeApp) Create(deviceType *entity.DeviceType) (*entity.DeviceType, error) {
	return deviceTypeApp.deviceTypeRepository.Create(deviceType)
}

func (deviceTypeApp *DeviceTypeApp) Get(id string) (*entity.DeviceType, error) {
	return deviceTypeApp.deviceTypeRepository.Get(id)
}

func (deviceTypeApp *DeviceTypeApp) List(conditions string) ([]entity.DeviceType, error) {
	return deviceTypeApp.deviceTypeRepository.List(conditions)
}

func (deviceTypeApp *DeviceTypeApp) Update(id string, deviceType *entity.DeviceType) (*entity.DeviceType, error) {
	return deviceTypeApp.deviceTypeRepository.Update(id, deviceType)
}

type DeviceTypeAppInterface interface {
	Create(*entity.DeviceType) (*entity.DeviceType, error)
	Get(string) (*entity.DeviceType, error)
	List(string) ([]entity.DeviceType, error)
	Update(string, *entity.DeviceType) (*entity.DeviceType, error)
}
