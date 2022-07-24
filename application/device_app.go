package application

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"
)

type DeviceApp struct {
	deviceRepository repository.DeviceRepository
}

var _ DeviceAppInterface = &DeviceApp{}

func NewDeviceApp(deviceRepository repository.DeviceRepository) *DeviceApp {
	return &DeviceApp{
		deviceRepository: deviceRepository,
	}
}

func (deviceApp *DeviceApp) Create(device *entity.Device) (*entity.Device, error) {
	return deviceApp.deviceRepository.Create(device)
}

func (deviceApp *DeviceApp) Get(id string) (*entity.Device, error) {
	return deviceApp.deviceRepository.Get(id)
}

func (deviceApp *DeviceApp) List(conditions string) ([]entity.Device, error) {
	return deviceApp.deviceRepository.List(conditions)
}

func (deviceApp *DeviceApp) Update(id string, device *entity.Device) (*entity.Device, error) {
	return deviceApp.deviceRepository.Update(id, device)
}

type DeviceAppInterface interface {
	Create(device *entity.Device) (*entity.Device, error)
	Get(id string) (*entity.Device, error)
	List(conditions string) ([]entity.Device, error)
	Update(id string, device *entity.Device) (*entity.Device, error)
}
