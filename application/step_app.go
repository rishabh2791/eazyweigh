package application

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"
)

type StepApp struct {
	stepRepository repository.StepRepository
}

var _ StepAppInterface = &StepApp{}

func NewStepApp(stepRepository repository.StepRepository) *StepApp {
	return &StepApp{
		stepRepository: stepRepository,
	}
}

func (stepApp *StepApp) Create(step *entity.Step) (*entity.Step, error) {
	return stepApp.stepRepository.Create(step)
}

func (stepApp *StepApp) Get(id string) (*entity.Step, error) {
	return stepApp.stepRepository.Get(id)
}

func (stepApp *StepApp) List(conditions string) ([]entity.Step, error) {
	return stepApp.stepRepository.List(conditions)
}

func (stepApp *StepApp) Update(id string, step *entity.Step) (*entity.Step, error) {
	return stepApp.stepRepository.Update(id, step)
}

type StepAppInterface interface {
	Create(step *entity.Step) (*entity.Step, error)
	Get(id string) (*entity.Step, error)
	List(conditions string) ([]entity.Step, error)
	Update(id string, step *entity.Step) (*entity.Step, error)
}
