package persistance

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type JobRepo struct {
	DB     *gorm.DB
	Logger hclog.Logger
}

var _ repository.JobRepository = &JobRepo{}

func NewJobRepo(db *gorm.DB, logger hclog.Logger) *JobRepo {
	return &JobRepo{
		DB:     db,
		Logger: logger,
	}
}

func (jobRepo *JobRepo) Create(job *entity.Job) (*entity.Job, error) {
	validationErr := job.Validate()
	if validationErr != nil {
		return nil, validationErr
	}

	creationErr := jobRepo.DB.Create(*job).Error
	if creationErr != nil {
		return nil, creationErr
	}

	return job, nil
}

func (jobRepo *JobRepo) Get(jobCode string) (*entity.Job, error) {
	job := entity.Job{}

	getErr := jobRepo.DB.Preload(clause.Associations).Where("job_code = ?", jobCode).Take(&job).Error
	if getErr != nil {
		return nil, getErr
	}

	return &job, nil
}

func (jobRepo *JobRepo) List(conditions string) ([]entity.Job, error) {
	jobs := []entity.Job{}

	getErr := jobRepo.DB.Preload(clause.Associations).Where(conditions).Find(&jobs).Error
	if getErr != nil {
		return nil, getErr
	}

	return jobs, nil
}

func (jobRepo *JobRepo) Update(id string, update *entity.Job) (*entity.Job, error) {
	existingJob := entity.Job{}
	getErr := jobRepo.DB.Preload(clause.Associations).Where("id = ?", id).Take(&existingJob).Error
	if getErr != nil {
		return nil, getErr
	}

	updationErr := jobRepo.DB.Table(entity.Job{}.Tablename()).Where("id = ?", id).Updates(update).Error
	if updationErr != nil {
		return nil, updationErr
	}

	updated := entity.Job{}
	jobRepo.DB.Preload(clause.Associations).Take(&updated)

	return &updated, nil
}
