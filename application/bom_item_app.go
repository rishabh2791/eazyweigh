package application

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"
)

type BOMItemApp struct {
	bomItemRepository repository.BOMItemsRepository
}

var _ BOMItemAppInterface = &BOMItemApp{}

func NewBOMItemApp(bomItemRepository repository.BOMItemsRepository) *BOMItemApp {
	return &BOMItemApp{
		bomItemRepository: bomItemRepository,
	}
}

type BOMItemAppInterface interface {
	Create(bomItems *entity.BOMItem) (*entity.BOMItem, error)
	Get(id string) (*entity.BOMItem, error)
	List(conditions string) ([]entity.BOMItem, error)
	Update(id string, update *entity.BOMItem) (*entity.BOMItem, error)
}

func (bomItemApp *BOMItemApp) Create(bomItems *entity.BOMItem) (*entity.BOMItem, error) {
	return bomItemApp.bomItemRepository.Create(bomItems)
}

func (bomItemApp *BOMItemApp) Get(id string) (*entity.BOMItem, error) {
	return bomItemApp.bomItemRepository.Get(id)
}

func (bomItemApp *BOMItemApp) List(conditions string) ([]entity.BOMItem, error) {
	return bomItemApp.bomItemRepository.List(conditions)
}

func (bomItemApp *BOMItemApp) Update(id string, update *entity.BOMItem) (*entity.BOMItem, error) {
	return bomItemApp.bomItemRepository.Update(id, update)
}
