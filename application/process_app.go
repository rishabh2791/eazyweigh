package application

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"
)

type ProcessApp struct {
	processRepository repository.ProcessRepository
}

var _ ProcessAppInterface = &ProcessApp{}

func NewProcessApp(processRepository repository.ProcessRepository) *ProcessApp {
	return &ProcessApp{
		processRepository: processRepository,
	}
}

func (processApp *ProcessApp) Create(process *entity.Process) (*entity.Process, error) {
	return processApp.processRepository.Create(process)
}

func (processApp *ProcessApp) Get(materialID string) (*entity.Process, error) {
	return processApp.processRepository.Get(materialID)
}

func (processApp *ProcessApp) List(conditions string) ([]entity.Process, error) {
	return processApp.processRepository.List(conditions)
}

func (processApp *ProcessApp) Update(id string, process *entity.Process) (*entity.Process, error) {
	return processApp.processRepository.Update(id, process)
}

type ProcessAppInterface interface {
	Create(process *entity.Process) (*entity.Process, error)
	Get(materialID string) (*entity.Process, error)
	List(conditions string) ([]entity.Process, error)
	Update(id string, process *entity.Process) (*entity.Process, error)
}
