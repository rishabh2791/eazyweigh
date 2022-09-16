package repository

import "eazyweigh/domain/entity"

type JobRepository interface {
	Create(job *entity.Job) (*entity.Job, error)
	Get(jobCode string) (*entity.Job, error)
	List(conditions string) ([]entity.Job, error)
	Update(jobCode string, update *entity.Job) (*entity.Job, error)
	PullFromRemote(factoryID string, username string) error
}
