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

	creationErr := jobRepo.DB.Create(&job).Error
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

	getErr := jobRepo.DB.
		Preload("Factory.Address").
		Preload("Factory.CreatedBy").
		Preload("Factory.CreatedBy.UserRole").
		Preload("Factory.UpdatedBy").
		Preload("Factory.UpdatedBy.UserRole").
		Preload("Material.UnitOfMeasure").
		Preload("Material.UnitOfMeasure.Factory").
		Preload("Material.UnitOfMeasure.Factory.Address").
		Preload("Material.UnitOfMeasure.Factory.CreatedBy").
		Preload("Material.UnitOfMeasure.Factory.CreatedBy.UserRole").
		Preload("Material.UnitOfMeasure.Factory.UpdatedBy").
		Preload("Material.UnitOfMeasure.Factory.UpdatedBy.UserRole").
		Preload("Material.UnitOfMeasure.CreatedBy").
		Preload("Material.UnitOfMeasure.CreatedBy.UserRole").
		Preload("Material.UnitOfMeasure.UpdatedBy").
		Preload("Material.UnitOfMeasure.UpdatedBy.UserRole").
		Preload("Material.CreatedBy").
		Preload("Material.CreatedBy.UserRole").
		Preload("Material.UpdatedBy").
		Preload("Material.UpdatedBy.UserRole").
		Preload("UnitOfMeasure.Factory").
		Preload("UnitOfMeasure.Factory.Address").
		Preload("UnitOfMeasure.Factory.CreatedBy").
		Preload("UnitOfMeasure.Factory.CreatedBy.UserRole").
		Preload("UnitOfMeasure.Factory.UpdatedBy").
		Preload("UnitOfMeasure.Factory.UpdatedBy.UserRole").
		Preload("UnitOfMeasure.CreatedBy").
		Preload("UnitOfMeasure.CreatedBy.UserRole").
		Preload("UnitOfMeasure.UpdatedBy").
		Preload("UnitOfMeasure.UpdatedBy.UserRole").
		Preload("JobItems.Material").
		Preload("JobItems.Material.UnitOfMeasure").
		Preload("JobItems.Material.UnitOfMeasure.Factory").
		Preload("JobItems.Material.UnitOfMeasure.Factory.Address").
		Preload("JobItems.Material.UnitOfMeasure.Factory.CreatedBy").
		Preload("JobItems.Material.UnitOfMeasure.Factory.CreatedBy.UserRole").
		Preload("JobItems.Material.UnitOfMeasure.Factory.UpdatedBy").
		Preload("JobItems.Material.UnitOfMeasure.Factory.UpdatedBy.UserRole").
		Preload("JobItems.Material.UnitOfMeasure.CreatedBy").
		Preload("JobItems.Material.UnitOfMeasure.CreatedBy.UserRole").
		Preload("JobItems.Material.UnitOfMeasure.UpdatedBy").
		Preload("JobItems.Material.UnitOfMeasure.UpdatedBy.UserRole").
		Preload("JobItems.Material.CreatedBy").
		Preload("JobItems.Material.CreatedBy.UserRole").
		Preload("JobItems.Material.UpdatedBy").
		Preload("JobItems.Material.UpdatedBy.UserRole").
		Preload("JobItems.UnitOfMeasure").
		Preload("JobItems.UnitOfMeasure.Factory").
		Preload("JobItems.UnitOfMeasure.Factory.Address").
		Preload("JobItems.UnitOfMeasure.Factory.CreatedBy").
		Preload("JobItems.UnitOfMeasure.Factory.CreatedBy.UserRole").
		Preload("JobItems.UnitOfMeasure.Factory.UpdatedBy").
		Preload("JobItems.UnitOfMeasure.Factory.UpdatedBy.UserRole").
		Preload("JobItems.UnitOfMeasure.CreatedBy").
		Preload("JobItems.UnitOfMeasure.CreatedBy.UserRole").
		Preload("JobItems.UnitOfMeasure.UpdatedBy").
		Preload("JobItems.UnitOfMeasure.UpdatedBy.UserRole").
		Preload("JobItems.CreatedBy").
		Preload("JobItems.CreatedBy.UserRole").
		Preload("JobItems.UpdatedBy").
		Preload("JobItems.UpdatedBy.UserRole").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload("JobItems.JobItemWeighing").
		Preload("JobItems.JobItemWeighing.CreatedBy").
		Preload("JobItems.JobItemWeighing.CreatedBy.UserRole").
		Preload("JobItems.JobItemWeighing.UpdatedBy").
		Preload("JobItems.JobItemWeighing.UpdatedBy.UserRole").
		Preload(clause.Associations).Where(conditions).Find(&jobs).Error
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
