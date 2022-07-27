package persistance

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProcessRepo struct {
	DB     *gorm.DB
	Logger hclog.Logger
}

var _ repository.ProcessRepository = &ProcessRepo{}

func NewProcessRepo(db *gorm.DB, logger hclog.Logger) *ProcessRepo {
	return &ProcessRepo{
		DB:     db,
		Logger: logger,
	}
}

func (processRepo *ProcessRepo) Create(process *entity.Process) (*entity.Process, error) {
	validationErr := process.Validate()
	if validationErr == nil {
		return nil, validationErr
	}

	creationErr := processRepo.DB.Create(&process).Error
	return process, creationErr
}

func (processRepo *ProcessRepo) Get(materialID string) (*entity.Process, error) {
	process := entity.Process{}

	getErr := processRepo.DB.
		Preload("Material.CreatedBy").
		Preload("Material.UpdatedBy").
		Preload("Material.CreatedBy.UserRole").
		Preload("Material.UpdatedBy.UserRole").
		Preload("Material.UnitOfMeasure.Factory").
		Preload("Material.UnitOfMeasure.Factory.Address").
		Preload("Material.UnitOfMeasure.Factory.CreatedBy").
		Preload("Material.UnitOfMeasure.Factory.UpdatedBy").
		Preload("Material.UnitOfMeasure.Factory.CreatedBy.UserRole").
		Preload("Material.UnitOfMeasure.Factory.UpdatedBy.UserRole").
		Preload("Material.UnitOfMeasure.CreatedBy").
		Preload("Material.UnitOfMeasure.UpdatedBy").
		Preload("Material.UnitOfMeasure.CreatedBy.UserRole").
		Preload("Material.UnitOfMeasure.UpdatedBy.UserRole").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload("Step.StepType.Factory.Address").
		Preload("Step.StepType.Factory.CreatedBy").
		Preload("Step.StepType.Factory.CreatedBy.UserRole").
		Preload("Step.StepType.Factory.UpdatedBy").
		Preload("Step.StepType.Factory.UpdatedBy.UserRole").
		Preload("Step.StepType.CreatedBy").
		Preload("Step.StepType.CreatedBy.UserRole").
		Preload("Step.StepType.UpdatedBy").
		Preload("Step.StepType.UpdatedBy.UserRole").
		Preload("Step.CreatedBy").
		Preload("Step.UpdatedBy").
		Preload("Step.CreatedBy.UserRole").
		Preload("Step.UpdatedBy.UserRole").
		Preload(clause.Associations).Where("material_id = ?", materialID).First(&process).Error
	return &process, getErr
}

func (processRepo *ProcessRepo) List(conditions string) ([]entity.Process, error) {
	processes := []entity.Process{}

	getErr := processRepo.DB.
		Preload("Material.CreatedBy").
		Preload("Material.UpdatedBy").
		Preload("Material.CreatedBy.UserRole").
		Preload("Material.UpdatedBy.UserRole").
		Preload("Material.UnitOfMeasure.Factory").
		Preload("Material.UnitOfMeasure.Factory.Address").
		Preload("Material.UnitOfMeasure.Factory.CreatedBy").
		Preload("Material.UnitOfMeasure.Factory.UpdatedBy").
		Preload("Material.UnitOfMeasure.Factory.CreatedBy.UserRole").
		Preload("Material.UnitOfMeasure.Factory.UpdatedBy.UserRole").
		Preload("Material.UnitOfMeasure.CreatedBy").
		Preload("Material.UnitOfMeasure.UpdatedBy").
		Preload("Material.UnitOfMeasure.CreatedBy.UserRole").
		Preload("Material.UnitOfMeasure.UpdatedBy.UserRole").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload("Step.StepType.Factory.Address").
		Preload("Step.StepType.Factory.CreatedBy").
		Preload("Step.StepType.Factory.CreatedBy.UserRole").
		Preload("Step.StepType.Factory.UpdatedBy").
		Preload("Step.StepType.Factory.UpdatedBy.UserRole").
		Preload("Step.StepType.CreatedBy").
		Preload("Step.StepType.CreatedBy.UserRole").
		Preload("Step.StepType.UpdatedBy").
		Preload("Step.StepType.UpdatedBy.UserRole").
		Preload("Step.CreatedBy").
		Preload("Step.UpdatedBy").
		Preload("Step.CreatedBy.UserRole").
		Preload("Step.UpdatedBy.UserRole").
		Preload(clause.Associations).Where(conditions).Find(&processes).Error
	return processes, getErr
}

func (processRepo *ProcessRepo) Update(id string, process *entity.Process) (*entity.Process, error) {
	existingProcess := entity.Process{}

	getErr := processRepo.DB.
		Preload(clause.Associations).Where("id = ?", id).First(&existingProcess).Error
	if getErr != nil {
		return nil, getErr
	}

	updationErr := processRepo.DB.Table(process.Tablename()).Updates(&process).Error
	if updationErr != nil {
		return nil, updationErr
	}

	updated := entity.Process{}
	processRepo.DB.
		Preload("Material.CreatedBy").
		Preload("Material.UpdatedBy").
		Preload("Material.CreatedBy.UserRole").
		Preload("Material.UpdatedBy.UserRole").
		Preload("Material.UnitOfMeasure.Factory").
		Preload("Material.UnitOfMeasure.Factory.Address").
		Preload("Material.UnitOfMeasure.Factory.CreatedBy").
		Preload("Material.UnitOfMeasure.Factory.UpdatedBy").
		Preload("Material.UnitOfMeasure.Factory.CreatedBy.UserRole").
		Preload("Material.UnitOfMeasure.Factory.UpdatedBy.UserRole").
		Preload("Material.UnitOfMeasure.CreatedBy").
		Preload("Material.UnitOfMeasure.UpdatedBy").
		Preload("Material.UnitOfMeasure.CreatedBy.UserRole").
		Preload("Material.UnitOfMeasure.UpdatedBy.UserRole").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload("Step.StepType.Factory.Address").
		Preload("Step.StepType.Factory.CreatedBy").
		Preload("Step.StepType.Factory.CreatedBy.UserRole").
		Preload("Step.StepType.Factory.UpdatedBy").
		Preload("Step.StepType.Factory.UpdatedBy.UserRole").
		Preload("Step.StepType.CreatedBy").
		Preload("Step.StepType.CreatedBy.UserRole").
		Preload("Step.StepType.UpdatedBy").
		Preload("Step.StepType.UpdatedBy.UserRole").
		Preload("Step.CreatedBy").
		Preload("Step.UpdatedBy").
		Preload("Step.CreatedBy.UserRole").
		Preload("Step.UpdatedBy.UserRole").
		Preload(clause.Associations).Where("id = ?", id).First(&updated)
	return &updated, nil
}
