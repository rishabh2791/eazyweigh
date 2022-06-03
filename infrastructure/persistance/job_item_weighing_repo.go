package persistance

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type JobItemWeighingRepo struct {
	DB     *gorm.DB
	Logger hclog.Logger
}

var _ repository.JobItemWeighingRepository = &JobItemWeighingRepo{}

func NewJobItemWeighingRepo(db *gorm.DB, logger hclog.Logger) *JobItemWeighingRepo {
	return &JobItemWeighingRepo{
		DB:     db,
		Logger: logger,
	}
}

func (jobItemWeighingRepo *JobItemWeighingRepo) Create(jobItemWeight *entity.JobItemWeighing) (*entity.JobItemWeighing, error) {
	validationErr := jobItemWeight.Validate()
	if validationErr != nil {
		return nil, validationErr
	}

	creationErr := jobItemWeighingRepo.DB.Create(&jobItemWeight).Error
	if creationErr != nil {
		return nil, creationErr
	}

	jobItem := entity.JobItem{}
	getErr := jobItemWeighingRepo.DB.Where("id = ?", jobItemWeight.JOBItemID).Take(&jobItem).Error
	if getErr != nil {
		return nil, getErr
	}

	update := map[string]interface{}{
		"actual_weight": jobItem.ActualWeight + jobItemWeight.Weight,
	}

	if (jobItem.ActualWeight+jobItemWeight.Weight) <= jobItem.UpperBound && (jobItem.ActualWeight+jobItemWeight.Weight) >= jobItem.LowerBound {
		update["complete"] = true
	}

	updationErr := jobItemWeighingRepo.DB.Model(&entity.JobItem{}).Where("id = ?", jobItemWeight.JOBItemID).Updates(update).Error
	if updationErr != nil {
		return nil, updationErr
	}

	job := entity.Job{}

	jobItemWeighingRepo.DB.
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
		Preload(clause.Associations).
		Where("id = ?", jobItem.JobID).Take(&job)
	complete := job.IsComplete()

	jobItemWeighingRepo.DB.Table(entity.Job{}.Tablename()).Where("id = ?", jobItem.JobID).Update("complete = ?", complete)

	return jobItemWeight, nil
}

func (jobItemWeighingRepo *JobItemWeighingRepo) List(jobItemID string) ([]entity.JobItemWeighing, error) {
	jobItemWeighings := []entity.JobItemWeighing{}

	getErr := jobItemWeighingRepo.DB.
		Preload("JobItem.Material.UnitOfMeasure").
		Preload("JobItem.Material.UnitOfMeasure.Factory").
		Preload("JobItem.Material.UnitOfMeasure.Factory.CreatedBy").
		Preload("JobItem.Material.UnitOfMeasure.Factory.CreatedBy.UserRole").
		Preload("JobItem.Material.UnitOfMeasure.Factory.UpdatedBy").
		Preload("JobItem.Material.UnitOfMeasure.Factory.UpdatedBy.UserRole").
		Preload("JobItem.Material.UnitOfMeasure.CreatedBy").
		Preload("JobItem.Material.UnitOfMeasure.CreatedBy.UserRole").
		Preload("JobItem.Material.UnitOfMeasure.UpdatedBy").
		Preload("JobItem.Material.UnitOfMeasure.UpdatedBy.UserRole").
		Preload("JobItem.Material.UnitOfMeasure.Factory.Address").
		Preload("JobItem.UnitOfMeasure.Factory").
		Preload("JobItem.UnitOfMeasure.Factory.CreatedBy").
		Preload("JobItem.UnitOfMeasure.Factory.CreatedBy.UserRole").
		Preload("JobItem.UnitOfMeasure.Factory.UpdatedBy").
		Preload("JobItem.UnitOfMeasure.Factory.UpdatedBy.UserRole").
		Preload("JobItem.UnitOfMeasure.CreatedBy").
		Preload("JobItem.UnitOfMeasure.CreatedBy.UserRole").
		Preload("JobItem.UnitOfMeasure.UpdatedBy").
		Preload("JobItem.UnitOfMeasure.UpdatedBy.UserRole").
		Preload("JobItem.UnitOfMeasure.Factory.Address").
		Preload("JobItem.CreatedBy.UserRole").
		Preload("JobItem.UpdatedBy.UserRole").
		Preload("JobItem.Material.CreatedBy").
		Preload("JobItem.Material.CreatedBy.UserRole").
		Preload("JobItem.Material.UpdatedBy").
		Preload("JobItem.Material.UpdatedBy.UserRole").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).Where("job_item_id = ?", jobItemID).Find(&jobItemWeighings).Error
	if getErr != nil {
		return nil, getErr
	}

	return jobItemWeighings, nil
}
