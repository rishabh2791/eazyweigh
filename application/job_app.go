package application

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"
)

type JobApp struct {
	jobRepository repository.JobRepository
}

var _ JobAppInterface = &JobApp{}

func NewJobApp(jobRepositoty repository.JobRepository) *JobApp {
	return &JobApp{
		jobRepository: jobRepositoty,
	}
}

type JobAppInterface interface {
	Create(job *entity.Job) (*entity.Job, error)
	Get(jobCode string) (*entity.Job, error)
	List(conditions string) ([]entity.Job, error)
	Update(id string, update *entity.Job) (*entity.Job, error)
	PullFromRemote(factoryID string, username string) error
}

func (jobApp *JobApp) Create(job *entity.Job) (*entity.Job, error) {
	return jobApp.jobRepository.Create(job)
}

func (jobApp *JobApp) Get(jobCode string) (*entity.Job, error) {
	return jobApp.jobRepository.Get(jobCode)
}

func (jobApp *JobApp) List(conditions string) ([]entity.Job, error) {
	return jobApp.jobRepository.List(conditions)
}

func (jobApp *JobApp) Update(id string, update *entity.Job) (*entity.Job, error) {
	return jobApp.jobRepository.Update(id, update)
}

func (jobApp *JobApp) PullFromRemote(factoryID string, username string) error {
	return jobApp.jobRepository.PullFromRemote(factoryID, username)
}
