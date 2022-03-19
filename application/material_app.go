package application

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"
)

type MaterialApp struct {
	materialRepository repository.MaterialRepository
}

var _ MaterialAppInterface = &MaterialApp{}

func NewMaterialApp(materialRepository repository.MaterialRepository) *MaterialApp {
	return &MaterialApp{
		materialRepository: materialRepository,
	}
}

type MaterialAppInterface interface {
	Create(material *entity.Material) (*entity.Material, error)
	Get(id string) (*entity.Material, error)
	List(conditions string) ([]entity.Material, error)
	Update(id string, update *entity.Material) (*entity.Material, error)
}

func (materialApp *MaterialApp) Create(material *entity.Material) (*entity.Material, error) {
	return materialApp.materialRepository.Create(material)
}

func (materialApp *MaterialApp) Get(id string) (*entity.Material, error) {
	return materialApp.materialRepository.Get(id)
}

func (materialApp *MaterialApp) List(conditions string) ([]entity.Material, error) {
	return materialApp.materialRepository.List(conditions)
}

func (materialApp *MaterialApp) Update(id string, update *entity.Material) (*entity.Material, error) {
	return materialApp.materialRepository.Update(id, update)
}
