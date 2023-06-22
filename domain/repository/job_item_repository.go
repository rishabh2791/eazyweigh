package repository

import "eazyweigh/domain/entity"

type JobItemRepository interface {
	Create(jobItem *entity.JobItem) (*entity.JobItem, error)
	Get(conditions string) ([]entity.JobItem, error)
	Update(id string, jobItem *entity.JobItem) (*entity.JobItem, error)
}
