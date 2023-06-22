package repository

import "eazyweigh/domain/entity"

type BatchRunRepository interface {
	Create(batch *entity.BatchRun) (*entity.BatchRun, error)
	Get(id string) (*entity.BatchRun, error)
	List(conditions string) ([]entity.BatchRun, error)
	Update(id string, batch *entity.BatchRun) (*entity.BatchRun, error)
}
