package application

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"
)

type JobAssignmentApp struct {
	jobAssignmentRepository repository.JobAssignmentRepository
}

var _ JobAssignmentAppInterface = &JobAssignmentApp{}

func NewJobAssignmentApp(jobAssignmentRepository repository.JobAssignmentRepository) *JobAssignmentApp {
	return &JobAssignmentApp{
		jobAssignmentRepository: jobAssignmentRepository,
	}
}

type JobAssignmentAppInterface interface {
	Create(jobAssignment *entity.JobAssignment) (*entity.JobAssignment, error)
	Get(id string) (*entity.JobAssignment, error)
	List(conditions string) ([]entity.JobAssignment, error)
	Update(id string, update *entity.JobAssignment) (*entity.JobAssignment, error)
}

func (jobAssignmentApp *JobAssignmentApp) Create(jobAssignment *entity.JobAssignment) (*entity.JobAssignment, error) {
	return jobAssignmentApp.jobAssignmentRepository.Create(jobAssignment)
}

func (jobAssignmentApp *JobAssignmentApp) Get(id string) (*entity.JobAssignment, error) {
	return jobAssignmentApp.jobAssignmentRepository.Get(id)
}

func (jobAssignmentApp *JobAssignmentApp) List(conditions string) ([]entity.JobAssignment, error) {
	return jobAssignmentApp.jobAssignmentRepository.List(conditions)
}

func (jobAssignmentApp *JobAssignmentApp) Update(id string, update *entity.JobAssignment) (*entity.JobAssignment, error) {
	return jobAssignmentApp.jobAssignmentRepository.Update(id, update)
}
