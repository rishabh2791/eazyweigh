package repository

import "eazyweigh/domain/entity"

type DeviceRepository interface {
	Create(device *entity.Device) (*entity.Device, error)
	Get(id string) (*entity.Device, error)
	List(conditions string) ([]entity.Device, error)
	Update(id string, device *entity.Device) (*entity.Device, error)
}
