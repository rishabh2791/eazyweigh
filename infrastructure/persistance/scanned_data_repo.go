package persistance

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"
	"log"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ScannedDataRepo struct {
	DB     *gorm.DB
	Logger hclog.Logger
}

var _ repository.ScannedDataRepository = &ScannedDataRepo{}

func NewScannedDataRepo(db *gorm.DB, logger hclog.Logger) *ScannedDataRepo {
	return &ScannedDataRepo{
		DB:     db,
		Logger: logger,
	}
}

func (scannedDataRepo *ScannedDataRepo) Create(scannedData *entity.ScannedData) (*entity.ScannedData, error) {
	validationErr := scannedData.Validate()
	if validationErr != nil {
		return nil, validationErr
	}

	creationErr := scannedDataRepo.DB.Create(&scannedData).Error
	if creationErr != nil {
		return nil, creationErr
	}

	return scannedData, nil
}

func (scannedDataRepo *ScannedDataRepo) Get(id string) (*entity.ScannedData, error) {
	scannedData := entity.ScannedData{}

	getErr := scannedDataRepo.DB.Preload(clause.Associations).Where("id = ?", id).Take(&scannedData).Error
	if getErr != nil {
		return nil, getErr
	}

	return &scannedData, nil
}

func (scannedDataRepo *ScannedDataRepo) List(conditions string) ([]entity.ScannedData, error) {
	scannedData := []entity.ScannedData{}

	log.Println(conditions)
	getErr := scannedDataRepo.DB.
		Preload("Terminal.UnitOfMeasure.Factory").
		Preload("Terminal.UnitOfMeasure.Factory.Address").
		Preload("Terminal.UnitOfMeasure.Factory.CreatedBy").
		Preload("Terminal.UnitOfMeasure.Factory.CreatedBy.UserRole").
		Preload("Terminal.UnitOfMeasure.Factory.UpdatedBy").
		Preload("Terminal.UnitOfMeasure.Factory.UpdatedBy.UserRole").
		Preload("Terminal.UnitOfMeasure.CreatedBy").
		Preload("Terminal.UnitOfMeasure.CreatedBy.UserRole").
		Preload("Terminal.UnitOfMeasure.UpdatedBy").
		Preload("Terminal.UnitOfMeasure.UpdatedBy.UserRole").
		Preload("Terminal.CreatedBy.UserRole").
		Preload("Terminal.UpdatedBy.UserRole").
		Preload("User.UserRole").
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
		Preload("Job.CreatedBy.UserRole").
		Preload("Job.UpdatedBy.UserRole").
		Preload("Job.JobItems.JobItemWeighing").
		Preload("Job.JobItems.JobItemWeighing.CreatedBy").
		Preload("Job.JobItems.JobItemWeighing.CreatedBy.UserRole").
		Preload("Job.JobItems.JobItemWeighing.UpdatedBy").
		Preload("Job.JobItems.JobItemWeighing.UpdatedBy.UserRole").
		Preload(clause.Associations).Where(conditions).Find(&scannedData).Error
	if getErr != nil {
		return nil, getErr
	}

	return scannedData, nil
}
