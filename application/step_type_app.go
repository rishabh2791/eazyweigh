package application

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"
)

type StepTypeApp struct {
	stepTypeRepository repository.StepTypeRepository
}

var _ StepTypeAppInterface = &StepTypeApp{}

func NewStepTypeApp(stepTypeRepository repository.StepTypeRepository) *StepTypeApp {
	return &StepTypeApp{
		stepTypeRepository: stepTypeRepository,
	}
}

func (stepTypeApp *StepTypeApp) Create(stepType *entity.StepType) (*entity.StepType, error) {
	return stepTypeApp.stepTypeRepository.Create(stepType)
}

func (stepTypeApp *StepTypeApp) Get(id string) (*entity.StepType, error) {
	return stepTypeApp.stepTypeRepository.Get(id)
}

func (stepTypeApp *StepTypeApp) List(conditions string) ([]entity.StepType, error) {
	return stepTypeApp.stepTypeRepository.List(conditions)
}

func (stepTypeApp *StepTypeApp) Update(id string, stepType *entity.StepType) (*entity.StepType, error) {
	return stepTypeApp.stepTypeRepository.Update(id, stepType)
}

type StepTypeAppInterface interface {
	Create(stepType *entity.StepType) (*entity.StepType, error)
	Get(id string) (*entity.StepType, error)
	List(conditions string) ([]entity.StepType, error)
	Update(id string, stepType *entity.StepType) (*entity.StepType, error)
}
