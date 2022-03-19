package application

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"
)

type JobItemApp struct {
	jobItemRepository repository.JobItemRepository
}

var _ JobItemAppInterface = &JobItemApp{}

func NewJobItemApp(jobItemRepository repository.JobItemRepository) *JobItemApp {
	return &JobItemApp{
		jobItemRepository: jobItemRepository,
	}
}

type JobItemAppInterface interface {
	Create(jobItem *entity.JobItem) (*entity.JobItem, error)
	Get(jobID string, conditions string) ([]entity.JobItem, error)
	Update(id string, jobItem *entity.JobItem) (*entity.JobItem, error)
}

func (jobItemApp *JobItemApp) Create(jobItem *entity.JobItem) (*entity.JobItem, error) {
	return jobItemApp.jobItemRepository.Create(jobItem)
}

func (jobItemApp *JobItemApp) Get(jobID string, conditions string) ([]entity.JobItem, error) {
	return jobItemApp.jobItemRepository.Get(jobID, conditions)
}

func (jobItemApp *JobItemApp) Update(id string, jobItem *entity.JobItem) (*entity.JobItem, error) {
	return jobItemApp.jobItemRepository.Update(id, jobItem)
}
