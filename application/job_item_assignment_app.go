package application

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"
)

type JobItemAssignmentApp struct {
	jobItemAssignmentRepository repository.JobItemAssignmentRepository
}

var _ JobItemAssignmentAppInterface = &JobItemAssignmentApp{}

func NewJobItemAssignmentApp(jobItemAssignmentRepository repository.JobItemAssignmentRepository) *JobItemAssignmentApp {
	return &JobItemAssignmentApp{
		jobItemAssignmentRepository: jobItemAssignmentRepository,
	}
}

type JobItemAssignmentAppInterface interface {
	Create(jobItemAssignment *entity.JobItemAssignment) (*entity.JobItemAssignment, error)
	Get(id string) (*entity.JobItemAssignment, error)
	List(conditions string) ([]entity.JobItemAssignment, error)
	Update(id string, update *entity.JobItemAssignment) (*entity.JobItemAssignment, error)
}

func (jobItemAssignmentApp *JobItemAssignmentApp) Create(jobItemAssignment *entity.JobItemAssignment) (*entity.JobItemAssignment, error) {
	return jobItemAssignmentApp.jobItemAssignmentRepository.Create(jobItemAssignment)
}

func (jobItemAssignmentApp *JobItemAssignmentApp) Get(id string) (*entity.JobItemAssignment, error) {
	return jobItemAssignmentApp.jobItemAssignmentRepository.Get(id)
}

func (jobItemAssignmentApp *JobItemAssignmentApp) List(conditions string) ([]entity.JobItemAssignment, error) {
	return jobItemAssignmentApp.jobItemAssignmentRepository.List(conditions)
}

func (jobItemAssignmentApp *JobItemAssignmentApp) Update(id string, update *entity.JobItemAssignment) (*entity.JobItemAssignment, error) {
	return jobItemAssignmentApp.jobItemAssignmentRepository.Update(id, update)
}
