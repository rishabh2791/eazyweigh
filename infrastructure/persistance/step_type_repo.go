package persistance

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
)

type StepTypeRepo struct {
	DB     *gorm.DB
	Logger hclog.Logger
}

var _ repository.StepTypeRepository = &StepTypeRepo{}

func NewStepTypeRepo(db *gorm.DB, logger hclog.Logger) *StepTypeRepo {
	return &StepTypeRepo{
		DB:     db,
		Logger: logger,
	}
}

func (stepTypeRepo *StepTypeRepo) Create(stepType *entity.StepType) (*entity.StepType, error) {
	validationErr := stepType.Validate()
	if validationErr != nil {
		return nil, validationErr
	}

	creationErr := stepTypeRepo.DB.Create(&stepType).Error
	return stepType, creationErr
}

func (stepTypeRepo *StepTypeRepo) Get(id string) (*entity.StepType, error) {
	stepType := entity.StepType{}

	getErr := stepTypeRepo.DB.Where("id = ?", stepType).First(&stepType).Error
	return &stepType, getErr
}

func (stepTypeRepo *StepTypeRepo) List(conditions string) ([]entity.StepType, error) {
	stepTypes := []entity.StepType{}

	getErr := stepTypeRepo.DB.Where(conditions).Find(&stepTypes).Error
	return stepTypes, getErr
}

func (stepTypeRepo *StepTypeRepo) Update(id string, stepType *entity.StepType) (*entity.StepType, error) {
	existingStepType := entity.StepType{}

	getErr := stepTypeRepo.DB.Where("id = ?", id).First(&existingStepType).Error

	if getErr != nil {
		return nil, getErr
	}

	updationErr := stepTypeRepo.DB.Table(stepType.Tablename()).Where("id = ?", id).Updates(&stepType).Error
	if updationErr != nil {
		return nil, updationErr
	}

	updated := entity.StepType{}
	stepTypeRepo.DB.Table(stepType.Tablename()).Where("id = ?", id).First(&updated)

	return &updated, nil
}
