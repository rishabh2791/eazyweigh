package persistance

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type StepRepo struct {
	DB     *gorm.DB
	Logger hclog.Logger
}

var _ repository.StepRepository = &StepRepo{}

func NewStepRepo(db *gorm.DB, logger hclog.Logger) *StepRepo {
	return &StepRepo{
		DB:     db,
		Logger: logger,
	}
}

func (stepRepo *StepRepo) Create(step *entity.Step) (*entity.Step, error) {
	validationErr := step.Validate()
	if validationErr != nil {
		return nil, validationErr
	}

	creationErr := stepRepo.DB.Create(&step).Error
	return step, creationErr
}

func (stepRepo *StepRepo) Get(id string) (*entity.Step, error) {
	step := entity.Step{}

	getErr := stepRepo.DB.
		Preload("StepType.Factory.Address").
		Preload("StepType.Factory.CreatedBy").
		Preload("StepType.Factory.CreatedBy.UserRole").
		Preload("StepType.Factory.UpdatedBy").
		Preload("StepType.Factory.UpdatedBy.UserRole").
		Preload("StepType.CreatedBy").
		Preload("StepType.CreatedBy.UserRole").
		Preload("StepType.UpdatedBy").
		Preload("StepType.UpdatedBy.UserRole").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).Where("id = ?", id).First(&step).Error
	return &step, getErr
}

func (stepRepo *StepRepo) List(conditions string) ([]entity.Step, error) {
	steps := []entity.Step{}

	getErr := stepRepo.DB.
		Preload("StepType.Factory.Address").
		Preload("StepType.Factory.CreatedBy").
		Preload("StepType.Factory.CreatedBy.UserRole").
		Preload("StepType.Factory.UpdatedBy").
		Preload("StepType.Factory.UpdatedBy.UserRole").
		Preload("StepType.CreatedBy").
		Preload("StepType.CreatedBy.UserRole").
		Preload("StepType.UpdatedBy").
		Preload("StepType.UpdatedBy.UserRole").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).Where(conditions).Find(steps).Error

	return steps, getErr
}

func (stepRepo *StepRepo) Update(id string, step *entity.Step) (*entity.Step, error) {
	existingStep := entity.Step{}

	getErr := stepRepo.DB.
		Preload(clause.Associations).Where("id = ?", id).First(&existingStep).Error
	if getErr != nil {
		return nil, getErr
	}

	updationErr := stepRepo.DB.Table(step.Tablename()).Where("id = ?", id).Updates(&step).Error
	if updationErr != nil {
		return nil, updationErr
	}

	updated := entity.Step{}
	stepRepo.DB.
		Preload("StepType.Factory.Address").
		Preload("StepType.Factory.CreatedBy").
		Preload("StepType.Factory.CreatedBy.UserRole").
		Preload("StepType.Factory.UpdatedBy").
		Preload("StepType.Factory.UpdatedBy.UserRole").
		Preload("StepType.CreatedBy").
		Preload("StepType.CreatedBy.UserRole").
		Preload("StepType.UpdatedBy").
		Preload("StepType.UpdatedBy.UserRole").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).Where("id = ?", id).First(&updated)
	return &updated, nil
}
