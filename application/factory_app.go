package application

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"
)

type FactoryApp struct {
	factoryRepository repository.FactoryRepository
}

var _ FactoryAppInterface = &FactoryApp{}

func NewFactoryRepository(factoryRepository repository.FactoryRepository) *FactoryApp {
	return &FactoryApp{
		factoryRepository: factoryRepository,
	}
}

type FactoryAppInterface interface {
	Create(factory *entity.Factory) (*entity.Factory, error)
	Get(id string) (*entity.Factory, error)
	List(conditions string) ([]entity.Factory, error)
	Update(id string, factory *entity.Factory) (*entity.Factory, error)
}

func (factoryApp *FactoryApp) Create(factory *entity.Factory) (*entity.Factory, error) {
	return factoryApp.factoryRepository.Create(factory)
}

func (factoryApp *FactoryApp) Get(id string) (*entity.Factory, error) {
	return factoryApp.factoryRepository.Get(id)
}

func (factoryApp *FactoryApp) List(conditions string) ([]entity.Factory, error) {
	return factoryApp.factoryRepository.List(conditions)
}

func (factoryApp *FactoryApp) Update(id string, factory *entity.Factory) (*entity.Factory, error) {
	return factoryApp.factoryRepository.Update(id, factory)
}
