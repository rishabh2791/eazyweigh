package repository

import "eazyweigh/domain/entity"

type JobItemWeighingRepository interface {
	Create(jobItemWeight *entity.JobItemWeighing) (*entity.JobItemWeighing, error)
}
