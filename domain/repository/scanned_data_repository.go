package repository

import "eazyweigh/domain/entity"

type ScannedDataRepository interface {
	Create(scannedData *entity.ScannedData) (*entity.ScannedData, error)
	Get(id string) (*entity.ScannedData, error)
	List(conditions string) ([]entity.ScannedData, error)
}
