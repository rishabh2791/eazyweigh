package application

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"
)

type BOMApp struct {
	bomRepository repository.BOMRepository
}

var _ BOMAppInterface = &BOMApp{}

func NewBOMApp(bomRepository repository.BOMRepository) *BOMApp {
	return &BOMApp{
		bomRepository: bomRepository,
	}
}

type BOMAppInterface interface {
	Create(bom *entity.BOM) (*entity.BOM, error)
	Get(id string) (*entity.BOM, error)
	List(conditions string) ([]entity.BOM, error)
	Update(id string, update *entity.BOM) (*entity.BOM, error)
}

func (bomApp *BOMApp) Create(bom *entity.BOM) (*entity.BOM, error) {
	return bomApp.bomRepository.Create(bom)
}

func (bomApp *BOMApp) Get(id string) (*entity.BOM, error) {
	return bomApp.bomRepository.Get(id)
}

func (bomApp *BOMApp) List(conditions string) ([]entity.BOM, error) {
	return bomApp.bomRepository.List(conditions)
}

func (bomApp *BOMApp) Update(id string, update *entity.BOM) (*entity.BOM, error) {
	return bomApp.bomRepository.Update(id, update)
}
