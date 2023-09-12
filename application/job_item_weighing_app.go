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
	Details(conditions string) ([]entity.WeighingBatch, error)
	MaterialDetails(materialID string) ([]entity.MaterialWeighing, error)
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

func (jobItemWeighingApp *JobItemWeighingApp) Details(conditions string) ([]entity.WeighingBatch, error) {
	return jobItemWeighingApp.jobItemWeighingRepository.Details(conditions)
}

func (jobItemWeighingApp *JobItemWeighingApp) MaterialDetails(materialID string) ([]entity.MaterialWeighing, error) {
	return jobItemWeighingApp.jobItemWeighingRepository.MaterialDetails(materialID)
}
