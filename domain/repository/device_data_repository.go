package repository

import "eazyweigh/domain/entity"

type DeviceDataRepository interface {
	Create(device *entity.DeviceData) (*entity.DeviceData, error)
	GetForDevice(deviceID string, conditions string) ([]entity.DeviceData, error)
}
