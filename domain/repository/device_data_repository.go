package repository

import "eazyweigh/domain/entity"

type DeviceDataRepository interface {
	Create(*entity.DeviceData) (*entity.DeviceData, error)
	Get(string) (*entity.DeviceData, error)
	List(string) ([]entity.DeviceData, error)
}
