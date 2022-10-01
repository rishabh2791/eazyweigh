package repository

import (
	"eazyweigh/domain/entity"
)

type JobItemWeighingRepository interface {
	Create(jobItemWeight *entity.JobItemWeighing) (*entity.JobItemWeighing, error)
	List(jobItemID string) ([]entity.JobItemWeighing, error)
	Update(id string, jobItemWeighing *entity.JobItemWeighing) (*entity.JobItemWeighing, error)
	Details(conditions string) ([]entity.WeighingBatch, error)
}
