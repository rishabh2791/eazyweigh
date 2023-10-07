package persistance

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"
	"errors"
	"time"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BatchRunRepo struct {
	DB     *gorm.DB
	Logger hclog.Logger
}

var _ repository.BatchRunRepository = &BatchRunRepo{}

func NewBatchRunRepo(db *gorm.DB, logger hclog.Logger) *BatchRunRepo {
	return &BatchRunRepo{
		DB:     db,
		Logger: logger,
	}
}

func (batchRepo *BatchRunRepo) Create(batch *entity.BatchRun) (*entity.BatchRun, error) {
	validationErr := batch.Validate()
	if validationErr != nil {
		return nil, validationErr
	}

	tx := batchRepo.DB.Begin()

	errs := ""
	runningBatches := []entity.BatchRun{}
	//Check if there is another batch running on the same vessel, if running stop it and start a new one
	getErr := tx.Where("start_time IS NOT NULL AND end_time IS NULL AND vessel_id = ?", batch.VesselID).Find(&runningBatches).Error
	if getErr != nil {
		tx.Rollback()
		return nil, getErr
	}

	if len(runningBatches) != 0 {
		for _, runningBatch := range runningBatches {
			err := tx.Table(runningBatch.Tablename()).Where("id = ?", runningBatch.ID).Update("end_time", time.Now().UTC()).Error
			if err != nil {
				errs += err.Error()
			}
		}
	}

	if len(errs) != 0 {
		tx.Rollback()
		return nil, errors.New(errs)
	}

	creationErr := tx.Create(&batch).Error
	if creationErr != nil {
		tx.Rollback()
		return nil, creationErr
	}

	tx.Commit()

	return batch, nil
}

func (batchRepo *BatchRunRepo) CreateSuper(batch *entity.BatchRun) (*entity.BatchRun, error) {
	validationErr := batch.Validate()
	if validationErr != nil {
		return nil, validationErr
	}

	tx := batchRepo.DB.Begin()

	creationErr := tx.Create(&batch).Error
	if creationErr != nil {
		tx.Rollback()
		return nil, creationErr
	}

	tx.Commit()

	return batch, nil
}

func (batchRepo *BatchRunRepo) Get(id string) (*entity.BatchRun, error) {
	batch := entity.BatchRun{}

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
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).Where("id = ?", id).First(&batch).Error
	return &batch, getErr
}

func (batchRepo *BatchRunRepo) List(conditions string) ([]entity.BatchRun, error) {
	batches := []entity.BatchRun{}

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
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).Where(conditions).Find(&batches).Error
	return batches, getErr
}

func (batchRepo *BatchRunRepo) Update(id string, batch *entity.BatchRun) (*entity.BatchRun, error) {
	existingBatch := entity.BatchRun{}

	getErr := batchRepo.DB.Where("id = ?", id).First(&existingBatch).Error
	if getErr != nil {
		return nil, getErr
	}

	updationErr := batchRepo.DB.Table(batch.Tablename()).Updates(&batch).Error
	if updationErr != nil {
		return nil, updationErr
	}

	updated := entity.BatchRun{}
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
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).Where("id = ?", id).First(&updated)
	return &updated, nil
}
