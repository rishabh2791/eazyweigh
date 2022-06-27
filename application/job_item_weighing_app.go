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
	List(jobItemID string) ([]entity.JobItemWeighing, error)
	Update(id string, jobItemWeighing *entity.JobItemWeighing) (*entity.JobItemWeighing, error)
}

func (jobItemWeighingApp *JobItemWeighingApp) Create(jobItemWeight *entity.JobItemWeighing) (*entity.JobItemWeighing, error) {
	return jobItemWeighingApp.jobItemWeighingRepository.Create(jobItemWeight)
}

func (jobItemWeighingApp *JobItemWeighingApp) List(jobItemID string) ([]entity.JobItemWeighing, error) {
	return jobItemWeighingApp.jobItemWeighingRepository.List(jobItemID)
}

func (jobItemWeighingApp *JobItemWeighingApp) Update(id string, jobItemWeighing *entity.JobItemWeighing) (*entity.JobItemWeighing, error) {
	return jobItemWeighingApp.jobItemWeighingRepository.Update(id, jobItemWeighing)
}
