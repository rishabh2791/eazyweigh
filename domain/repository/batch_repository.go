package repository

import "eazyweigh/domain/entity"

type BatchRepository interface {
	Create(batch *entity.Batch) (*entity.Batch, error)
	Get(id string) (*entity.Batch, error)
	List(conditions string) ([]entity.Batch, error)
	Update(id string, batch *entity.Batch) (*entity.Batch, error)
}
