package persistance

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BatchRepo struct {
	DB     *gorm.DB
	Logger hclog.Logger
}

var _ repository.BatchRepository = &BatchRepo{}

func NewBatchRepo(db *gorm.DB, logger hclog.Logger) *BatchRepo {
	return &BatchRepo{
		DB:     db,
		Logger: logger,
	}
}

func (batchRepo *BatchRepo) Create(batch *entity.Batch) (*entity.Batch, error) {
	validationErr := batch.Validate()
	if validationErr != nil {
		return nil, validationErr
	}

	creationErr := batchRepo.DB.Create(&batch).Error
	return batch, creationErr
}

func (batchRepo *BatchRepo) Get(id string) (*entity.Batch, error) {
	batch := entity.Batch{}

	getErr := batchRepo.DB.
		Preload("Job.Factory.Address").
		Preload("Job.Factory.CreatedBy").
		Preload("Job.Factory.CreatedBy.UserRole").
		Preload("Job.Factory.UpdatedBy").
		Preload("Job.Factory.UpdatedBy.UserRole").
		Preload("Job.Material.UnitOfMeasure").
		Preload("Job.Material.UnitOfMeasure.Factory").
		Preload("Job.Material.UnitOfMeasure.Factory.Address").
		Preload("Job.Material.UnitOfMeasure.Factory.CreatedBy").
		Preload("Job.Material.UnitOfMeasure.Factory.CreatedBy.UserRole").
		Preload("Job.Material.UnitOfMeasure.Factory.UpdatedBy").
		Preload("Job.Material.UnitOfMeasure.Factory.UpdatedBy.UserRole").
		Preload("Job.Material.UnitOfMeasure.CreatedBy").
		Preload("Job.Material.UnitOfMeasure.CreatedBy.UserRole").
		Preload("Job.Material.UnitOfMeasure.UpdatedBy").
		Preload("Job.Material.UnitOfMeasure.UpdatedBy.UserRole").
		Preload("Job.Material.CreatedBy").
		Preload("Job.Material.CreatedBy.UserRole").
		Preload("Job.Material.UpdatedBy").
		Preload("Job.Material.UpdatedBy.UserRole").
		Preload("Job.UnitOfMeasure.Factory").
		Preload("Job.UnitOfMeasure.Factory.Address").
		Preload("Job.UnitOfMeasure.Factory.CreatedBy").
		Preload("Job.UnitOfMeasure.Factory.CreatedBy.UserRole").
		Preload("Job.UnitOfMeasure.Factory.UpdatedBy").
		Preload("Job.UnitOfMeasure.Factory.UpdatedBy.UserRole").
		Preload("Job.UnitOfMeasure.CreatedBy").
		Preload("Job.UnitOfMeasure.CreatedBy.UserRole").
		Preload("Job.UnitOfMeasure.UpdatedBy").
		Preload("Job.UnitOfMeasure.UpdatedBy.UserRole").
		Preload("Job.JobItems.Material").
		Preload("Job.JobItems.Material.UnitOfMeasure").
		Preload("Job.JobItems.Material.UnitOfMeasure.Factory").
		Preload("Job.JobItems.Material.UnitOfMeasure.Factory.Address").
		Preload("Job.JobItems.Material.UnitOfMeasure.Factory.CreatedBy").
		Preload("Job.JobItems.Material.UnitOfMeasure.Factory.CreatedBy.UserRole").
		Preload("Job.JobItems.Material.UnitOfMeasure.Factory.UpdatedBy").
		Preload("Job.JobItems.Material.UnitOfMeasure.Factory.UpdatedBy.UserRole").
		Preload("Job.JobItems.Material.UnitOfMeasure.CreatedBy").
		Preload("Job.JobItems.Material.UnitOfMeasure.CreatedBy.UserRole").
		Preload("Job.JobItems.Material.UnitOfMeasure.UpdatedBy").
		Preload("Job.JobItems.Material.UnitOfMeasure.UpdatedBy.UserRole").
		Preload("Job.JobItems.Material.CreatedBy").
		Preload("Job.JobItems.Material.CreatedBy.UserRole").
		Preload("Job.JobItems.Material.UpdatedBy").
		Preload("Job.JobItems.Material.UpdatedBy.UserRole").
		Preload("Job.JobItems.UnitOfMeasure").
		Preload("Job.JobItems.UnitOfMeasure.Factory").
		Preload("Job.JobItems.UnitOfMeasure.Factory.Address").
		Preload("Job.JobItems.UnitOfMeasure.Factory.CreatedBy").
		Preload("Job.JobItems.UnitOfMeasure.Factory.CreatedBy.UserRole").
		Preload("Job.JobItems.UnitOfMeasure.Factory.UpdatedBy").
		Preload("Job.JobItems.UnitOfMeasure.Factory.UpdatedBy.UserRole").
		Preload("Job.JobItems.UnitOfMeasure.CreatedBy").
		Preload("Job.JobItems.UnitOfMeasure.CreatedBy.UserRole").
		Preload("Job.JobItems.UnitOfMeasure.UpdatedBy").
		Preload("Job.JobItems.UnitOfMeasure.UpdatedBy.UserRole").
		Preload("Job.JobItems.CreatedBy").
		Preload("Job.JobItems.CreatedBy.UserRole").
		Preload("Job.JobItems.UpdatedBy").
		Preload("Job.JobItems.UpdatedBy.UserRole").
		Preload("Job.CreatedBy").
		Preload("Job.UpdatedBy").
		Preload("Job.CreatedBy.UserRole").
		Preload("Job.UpdatedBy.UserRole").
		Preload("Job.JobItems.JobItemWeighing").
		Preload("Job.JobItems.JobItemWeighing.CreatedBy").
		Preload("Job.JobItems.JobItemWeighing.CreatedBy.UserRole").
		Preload("Job.JobItems.JobItemWeighing.UpdatedBy").
		Preload("Job.JobItems.JobItemWeighing.UpdatedBy.UserRole").
		Preload("Vessel.Factory.Address").
		Preload("Vessel.Factory.CreatedBy").
		Preload("Vessel.Factory.CreatedBy.UserRole").
		Preload("Vessel.Factory.UpdatedBy").
		Preload("Vessel.Factory.UpdatedBy.UserRole").
		Preload("Vessel.CreatedBy").
		Preload("Vessel.UpdatedBy").
		Preload("Vessel.CreatedBy.UserRole").
		Preload("Vessel.UpdatedBy.UserRole").
		Preload("Process.Material.CreatedBy").
		Preload("Process.Material.UpdatedBy").
		Preload("Process.Material.CreatedBy.UserRole").
		Preload("Process.Material.UpdatedBy.UserRole").
		Preload("Process.Material.UnitOfMeasure.Factory").
		Preload("Process.Material.UnitOfMeasure.Factory.Address").
		Preload("Process.Material.UnitOfMeasure.Factory.CreatedBy").
		Preload("Process.Material.UnitOfMeasure.Factory.UpdatedBy").
		Preload("Process.Material.UnitOfMeasure.Factory.CreatedBy.UserRole").
		Preload("Process.Material.UnitOfMeasure.Factory.UpdatedBy.UserRole").
		Preload("Process.Material.UnitOfMeasure.CreatedBy").
		Preload("Process.Material.UnitOfMeasure.UpdatedBy").
		Preload("Process.Material.UnitOfMeasure.CreatedBy.UserRole").
		Preload("Process.Material.UnitOfMeasure.UpdatedBy.UserRole").
		Preload("Process.CreatedBy").
		Preload("Process.UpdatedBy").
		Preload("Process.CreatedBy.UserRole").
		Preload("Process.UpdatedBy.UserRole").
		Preload("Process.Step.StepType.Factory.Address").
		Preload("Process.Step.StepType.Factory.CreatedBy").
		Preload("Process.Step.StepType.Factory.CreatedBy.UserRole").
		Preload("Process.Step.StepType.Factory.UpdatedBy").
		Preload("Process.Step.StepType.Factory.UpdatedBy.UserRole").
		Preload("Process.Step.StepType.CreatedBy").
		Preload("Process.Step.StepType.CreatedBy.UserRole").
		Preload("Process.Step.StepType.UpdatedBy").
		Preload("Process.Step.StepType.UpdatedBy.UserRole").
		Preload("Process.Step.CreatedBy").
		Preload("Process.Step.UpdatedBy").
		Preload("Process.Step.CreatedBy.UserRole").
		Preload("Process.Step.UpdatedBy.UserRole").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).Where("id = ?", id).First(&batch).Error
	return &batch, getErr
}

func (batchRepo *BatchRepo) List(conditions string) ([]entity.Batch, error) {
	batches := []entity.Batch{}

	getErr := batchRepo.DB.
		Preload("Job.Factory.Address").
		Preload("Job.Factory.CreatedBy").
		Preload("Job.Factory.CreatedBy.UserRole").
		Preload("Job.Factory.UpdatedBy").
		Preload("Job.Factory.UpdatedBy.UserRole").
		Preload("Job.Material.UnitOfMeasure").
		Preload("Job.Material.UnitOfMeasure.Factory").
		Preload("Job.Material.UnitOfMeasure.Factory.Address").
		Preload("Job.Material.UnitOfMeasure.Factory.CreatedBy").
		Preload("Job.Material.UnitOfMeasure.Factory.CreatedBy.UserRole").
		Preload("Job.Material.UnitOfMeasure.Factory.UpdatedBy").
		Preload("Job.Material.UnitOfMeasure.Factory.UpdatedBy.UserRole").
		Preload("Job.Material.UnitOfMeasure.CreatedBy").
		Preload("Job.Material.UnitOfMeasure.CreatedBy.UserRole").
		Preload("Job.Material.UnitOfMeasure.UpdatedBy").
		Preload("Job.Material.UnitOfMeasure.UpdatedBy.UserRole").
		Preload("Job.Material.CreatedBy").
		Preload("Job.Material.CreatedBy.UserRole").
		Preload("Job.Material.UpdatedBy").
		Preload("Job.Material.UpdatedBy.UserRole").
		Preload("Job.UnitOfMeasure.Factory").
		Preload("Job.UnitOfMeasure.Factory.Address").
		Preload("Job.UnitOfMeasure.Factory.CreatedBy").
		Preload("Job.UnitOfMeasure.Factory.CreatedBy.UserRole").
		Preload("Job.UnitOfMeasure.Factory.UpdatedBy").
		Preload("Job.UnitOfMeasure.Factory.UpdatedBy.UserRole").
		Preload("Job.UnitOfMeasure.CreatedBy").
		Preload("Job.UnitOfMeasure.CreatedBy.UserRole").
		Preload("Job.UnitOfMeasure.UpdatedBy").
		Preload("Job.UnitOfMeasure.UpdatedBy.UserRole").
		Preload("Job.JobItems.Material").
		Preload("Job.JobItems.Material.UnitOfMeasure").
		Preload("Job.JobItems.Material.UnitOfMeasure.Factory").
		Preload("Job.JobItems.Material.UnitOfMeasure.Factory.Address").
		Preload("Job.JobItems.Material.UnitOfMeasure.Factory.CreatedBy").
		Preload("Job.JobItems.Material.UnitOfMeasure.Factory.CreatedBy.UserRole").
		Preload("Job.JobItems.Material.UnitOfMeasure.Factory.UpdatedBy").
		Preload("Job.JobItems.Material.UnitOfMeasure.Factory.UpdatedBy.UserRole").
		Preload("Job.JobItems.Material.UnitOfMeasure.CreatedBy").
		Preload("Job.JobItems.Material.UnitOfMeasure.CreatedBy.UserRole").
		Preload("Job.JobItems.Material.UnitOfMeasure.UpdatedBy").
		Preload("Job.JobItems.Material.UnitOfMeasure.UpdatedBy.UserRole").
		Preload("Job.JobItems.Material.CreatedBy").
		Preload("Job.JobItems.Material.CreatedBy.UserRole").
		Preload("Job.JobItems.Material.UpdatedBy").
		Preload("Job.JobItems.Material.UpdatedBy.UserRole").
		Preload("Job.JobItems.UnitOfMeasure").
		Preload("Job.JobItems.UnitOfMeasure.Factory").
		Preload("Job.JobItems.UnitOfMeasure.Factory.Address").
		Preload("Job.JobItems.UnitOfMeasure.Factory.CreatedBy").
		Preload("Job.JobItems.UnitOfMeasure.Factory.CreatedBy.UserRole").
		Preload("Job.JobItems.UnitOfMeasure.Factory.UpdatedBy").
		Preload("Job.JobItems.UnitOfMeasure.Factory.UpdatedBy.UserRole").
		Preload("Job.JobItems.UnitOfMeasure.CreatedBy").
		Preload("Job.JobItems.UnitOfMeasure.CreatedBy.UserRole").
		Preload("Job.JobItems.UnitOfMeasure.UpdatedBy").
		Preload("Job.JobItems.UnitOfMeasure.UpdatedBy.UserRole").
		Preload("Job.JobItems.CreatedBy").
		Preload("Job.JobItems.CreatedBy.UserRole").
		Preload("Job.JobItems.UpdatedBy").
		Preload("Job.JobItems.UpdatedBy.UserRole").
		Preload("Job.CreatedBy").
		Preload("Job.UpdatedBy").
		Preload("Job.CreatedBy.UserRole").
		Preload("Job.UpdatedBy.UserRole").
		Preload("Job.JobItems.JobItemWeighing").
		Preload("Job.JobItems.JobItemWeighing.CreatedBy").
		Preload("Job.JobItems.JobItemWeighing.CreatedBy.UserRole").
		Preload("Job.JobItems.JobItemWeighing.UpdatedBy").
		Preload("Job.JobItems.JobItemWeighing.UpdatedBy.UserRole").
		Preload("Vessel.Factory.Address").
		Preload("Vessel.Factory.CreatedBy").
		Preload("Vessel.Factory.CreatedBy.UserRole").
		Preload("Vessel.Factory.UpdatedBy").
		Preload("Vessel.Factory.UpdatedBy.UserRole").
		Preload("Vessel.CreatedBy").
		Preload("Vessel.UpdatedBy").
		Preload("Vessel.CreatedBy.UserRole").
		Preload("Vessel.UpdatedBy.UserRole").
		Preload("Process.Material.CreatedBy").
		Preload("Process.Material.UpdatedBy").
		Preload("Process.Material.CreatedBy.UserRole").
		Preload("Process.Material.UpdatedBy.UserRole").
		Preload("Process.Material.UnitOfMeasure.Factory").
		Preload("Process.Material.UnitOfMeasure.Factory.Address").
		Preload("Process.Material.UnitOfMeasure.Factory.CreatedBy").
		Preload("Process.Material.UnitOfMeasure.Factory.UpdatedBy").
		Preload("Process.Material.UnitOfMeasure.Factory.CreatedBy.UserRole").
		Preload("Process.Material.UnitOfMeasure.Factory.UpdatedBy.UserRole").
		Preload("Process.Material.UnitOfMeasure.CreatedBy").
		Preload("Process.Material.UnitOfMeasure.UpdatedBy").
		Preload("Process.Material.UnitOfMeasure.CreatedBy.UserRole").
		Preload("Process.Material.UnitOfMeasure.UpdatedBy.UserRole").
		Preload("Process.CreatedBy").
		Preload("Process.UpdatedBy").
		Preload("Process.CreatedBy.UserRole").
		Preload("Process.UpdatedBy.UserRole").
		Preload("Process.Step.StepType.Factory.Address").
		Preload("Process.Step.StepType.Factory.CreatedBy").
		Preload("Process.Step.StepType.Factory.CreatedBy.UserRole").
		Preload("Process.Step.StepType.Factory.UpdatedBy").
		Preload("Process.Step.StepType.Factory.UpdatedBy.UserRole").
		Preload("Process.Step.StepType.CreatedBy").
		Preload("Process.Step.StepType.CreatedBy.UserRole").
		Preload("Process.Step.StepType.UpdatedBy").
		Preload("Process.Step.StepType.UpdatedBy.UserRole").
		Preload("Process.Step.CreatedBy").
		Preload("Process.Step.UpdatedBy").
		Preload("Process.Step.CreatedBy.UserRole").
		Preload("Process.Step.UpdatedBy.UserRole").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).Where(conditions).Find(&batches).Error
	return batches, getErr
}

func (batchRepo *BatchRepo) Update(id string, batch *entity.Batch) (*entity.Batch, error) {
	existingBatch := entity.Batch{}

	getErr := batchRepo.DB.Where("id = ?", id).First(&existingBatch).Error
	if getErr != nil {
		return nil, getErr
	}

	updationErr := batchRepo.DB.Table(batch.Tablename()).Updates(&batch).Error
	if updationErr != nil {
		return nil, updationErr
	}

	updated := entity.Batch{}
	batchRepo.DB.
		Preload("Job.Factory.Address").
		Preload("Job.Factory.CreatedBy").
		Preload("Job.Factory.CreatedBy.UserRole").
		Preload("Job.Factory.UpdatedBy").
		Preload("Job.Factory.UpdatedBy.UserRole").
		Preload("Job.Material.UnitOfMeasure").
		Preload("Job.Material.UnitOfMeasure.Factory").
		Preload("Job.Material.UnitOfMeasure.Factory.Address").
		Preload("Job.Material.UnitOfMeasure.Factory.CreatedBy").
		Preload("Job.Material.UnitOfMeasure.Factory.CreatedBy.UserRole").
		Preload("Job.Material.UnitOfMeasure.Factory.UpdatedBy").
		Preload("Job.Material.UnitOfMeasure.Factory.UpdatedBy.UserRole").
		Preload("Job.Material.UnitOfMeasure.CreatedBy").
		Preload("Job.Material.UnitOfMeasure.CreatedBy.UserRole").
		Preload("Job.Material.UnitOfMeasure.UpdatedBy").
		Preload("Job.Material.UnitOfMeasure.UpdatedBy.UserRole").
		Preload("Job.Material.CreatedBy").
		Preload("Job.Material.CreatedBy.UserRole").
		Preload("Job.Material.UpdatedBy").
		Preload("Job.Material.UpdatedBy.UserRole").
		Preload("Job.UnitOfMeasure.Factory").
		Preload("Job.UnitOfMeasure.Factory.Address").
		Preload("Job.UnitOfMeasure.Factory.CreatedBy").
		Preload("Job.UnitOfMeasure.Factory.CreatedBy.UserRole").
		Preload("Job.UnitOfMeasure.Factory.UpdatedBy").
		Preload("Job.UnitOfMeasure.Factory.UpdatedBy.UserRole").
		Preload("Job.UnitOfMeasure.CreatedBy").
		Preload("Job.UnitOfMeasure.CreatedBy.UserRole").
		Preload("Job.UnitOfMeasure.UpdatedBy").
		Preload("Job.UnitOfMeasure.UpdatedBy.UserRole").
		Preload("Job.JobItems.Material").
		Preload("Job.JobItems.Material.UnitOfMeasure").
		Preload("Job.JobItems.Material.UnitOfMeasure.Factory").
		Preload("Job.JobItems.Material.UnitOfMeasure.Factory.Address").
		Preload("Job.JobItems.Material.UnitOfMeasure.Factory.CreatedBy").
		Preload("Job.JobItems.Material.UnitOfMeasure.Factory.CreatedBy.UserRole").
		Preload("Job.JobItems.Material.UnitOfMeasure.Factory.UpdatedBy").
		Preload("Job.JobItems.Material.UnitOfMeasure.Factory.UpdatedBy.UserRole").
		Preload("Job.JobItems.Material.UnitOfMeasure.CreatedBy").
		Preload("Job.JobItems.Material.UnitOfMeasure.CreatedBy.UserRole").
		Preload("Job.JobItems.Material.UnitOfMeasure.UpdatedBy").
		Preload("Job.JobItems.Material.UnitOfMeasure.UpdatedBy.UserRole").
		Preload("Job.JobItems.Material.CreatedBy").
		Preload("Job.JobItems.Material.CreatedBy.UserRole").
		Preload("Job.JobItems.Material.UpdatedBy").
		Preload("Job.JobItems.Material.UpdatedBy.UserRole").
		Preload("Job.JobItems.UnitOfMeasure").
		Preload("Job.JobItems.UnitOfMeasure.Factory").
		Preload("Job.JobItems.UnitOfMeasure.Factory.Address").
		Preload("Job.JobItems.UnitOfMeasure.Factory.CreatedBy").
		Preload("Job.JobItems.UnitOfMeasure.Factory.CreatedBy.UserRole").
		Preload("Job.JobItems.UnitOfMeasure.Factory.UpdatedBy").
		Preload("Job.JobItems.UnitOfMeasure.Factory.UpdatedBy.UserRole").
		Preload("Job.JobItems.UnitOfMeasure.CreatedBy").
		Preload("Job.JobItems.UnitOfMeasure.CreatedBy.UserRole").
		Preload("Job.JobItems.UnitOfMeasure.UpdatedBy").
		Preload("Job.JobItems.UnitOfMeasure.UpdatedBy.UserRole").
		Preload("Job.JobItems.CreatedBy").
		Preload("Job.JobItems.CreatedBy.UserRole").
		Preload("Job.JobItems.UpdatedBy").
		Preload("Job.JobItems.UpdatedBy.UserRole").
		Preload("Job.CreatedBy").
		Preload("Job.UpdatedBy").
		Preload("Job.CreatedBy.UserRole").
		Preload("Job.UpdatedBy.UserRole").
		Preload("Job.JobItems.JobItemWeighing").
		Preload("Job.JobItems.JobItemWeighing.CreatedBy").
		Preload("Job.JobItems.JobItemWeighing.CreatedBy.UserRole").
		Preload("Job.JobItems.JobItemWeighing.UpdatedBy").
		Preload("Job.JobItems.JobItemWeighing.UpdatedBy.UserRole").
		Preload("Vessel.Factory.Address").
		Preload("Vessel.Factory.CreatedBy").
		Preload("Vessel.Factory.CreatedBy.UserRole").
		Preload("Vessel.Factory.UpdatedBy").
		Preload("Vessel.Factory.UpdatedBy.UserRole").
		Preload("Vessel.CreatedBy").
		Preload("Vessel.UpdatedBy").
		Preload("Vessel.CreatedBy.UserRole").
		Preload("Vessel.UpdatedBy.UserRole").
		Preload("Process.Material.CreatedBy").
		Preload("Process.Material.UpdatedBy").
		Preload("Process.Material.CreatedBy.UserRole").
		Preload("Process.Material.UpdatedBy.UserRole").
		Preload("Process.Material.UnitOfMeasure.Factory").
		Preload("Process.Material.UnitOfMeasure.Factory.Address").
		Preload("Process.Material.UnitOfMeasure.Factory.CreatedBy").
		Preload("Process.Material.UnitOfMeasure.Factory.UpdatedBy").
		Preload("Process.Material.UnitOfMeasure.Factory.CreatedBy.UserRole").
		Preload("Process.Material.UnitOfMeasure.Factory.UpdatedBy.UserRole").
		Preload("Process.Material.UnitOfMeasure.CreatedBy").
		Preload("Process.Material.UnitOfMeasure.UpdatedBy").
		Preload("Process.Material.UnitOfMeasure.CreatedBy.UserRole").
		Preload("Process.Material.UnitOfMeasure.UpdatedBy.UserRole").
		Preload("Process.CreatedBy").
		Preload("Process.UpdatedBy").
		Preload("Process.CreatedBy.UserRole").
		Preload("Process.UpdatedBy.UserRole").
		Preload("Process.Step.StepType.Factory.Address").
		Preload("Process.Step.StepType.Factory.CreatedBy").
		Preload("Process.Step.StepType.Factory.CreatedBy.UserRole").
		Preload("Process.Step.StepType.Factory.UpdatedBy").
		Preload("Process.Step.StepType.Factory.UpdatedBy.UserRole").
		Preload("Process.Step.StepType.CreatedBy").
		Preload("Process.Step.StepType.CreatedBy.UserRole").
		Preload("Process.Step.StepType.UpdatedBy").
		Preload("Process.Step.StepType.UpdatedBy.UserRole").
		Preload("Process.Step.CreatedBy").
		Preload("Process.Step.UpdatedBy").
		Preload("Process.Step.CreatedBy.UserRole").
		Preload("Process.Step.UpdatedBy.UserRole").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).Where("id = ?", id).First(&updated)
	return &updated, nil
}
