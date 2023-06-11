package repository

import "eazyweigh/domain/entity"

type DeviceTypeRepository interface {
	Create(*entity.DeviceType) (*entity.DeviceType, error)
	Get(string) (*entity.DeviceType, error)
	List(string) ([]entity.DeviceType, error)
	Update(string, *entity.DeviceType) (*entity.DeviceType, error)
}
