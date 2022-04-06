package application

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"
)

type JobItemWeighingApp struct {
	jobItemWeighingRepository repository.JobItemWeighingRepository
}

var _ JobItemWeighingAppInterface = &JobItemWeighingApp{}

func NewJobItemWeighingApp(jobItemWeighingRepository repository.JobItemWeighingRepository) *JobItemWeighingApp {
	return &JobItemWeighingApp{
		jobItemWeighingRepository: jobItemWeighingRepository,
	}
}

type JobItemWeighingAppInterface interface {
	Create(jobItemWeight *entity.JobItemWeighing) (*entity.JobItemWeighing, error)
}

func (jobItemWeighingApp *JobItemWeighingApp) Create(jobItemWeight *entity.JobItemWeighing) (*entity.JobItemWeighing, error) {
	return jobItemWeighingApp.jobItemWeighingRepository.Create(jobItemWeight)
}
