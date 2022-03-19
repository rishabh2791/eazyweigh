package application

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"
)

type UnitOfMeasureApp struct {
	uomRepo repository.UnitOfMeasureRepository
}

var _ UnitOfMeasureAppInterface = &UnitOfMeasureApp{}

func NewUnitOfMeasureApp(uomRepo repository.UnitOfMeasureRepository) *UnitOfMeasureApp {
	return &UnitOfMeasureApp{
		uomRepo: uomRepo,
	}

}

type UnitOfMeasureAppInterface interface {
	Create(uom *entity.UnitOfMeasure) (*entity.UnitOfMeasure, error)
	Get(id string) (*entity.UnitOfMeasure, error)
	List(conditions string) ([]entity.UnitOfMeasure, error)
	Update(id string, uom *entity.UnitOfMeasure) (*entity.UnitOfMeasure, error)
}

func (uomApp *UnitOfMeasureApp) Create(uom *entity.UnitOfMeasure) (*entity.UnitOfMeasure, error) {
	return uomApp.uomRepo.Create(uom)
}

func (uomApp *UnitOfMeasureApp) Get(id string) (*entity.UnitOfMeasure, error) {
	return uomApp.uomRepo.Get(id)
}

func (uomApp *UnitOfMeasureApp) List(conditions string) ([]entity.UnitOfMeasure, error) {
	return uomApp.uomRepo.List(conditions)
}

func (uomApp *UnitOfMeasureApp) Update(id string, uom *entity.UnitOfMeasure) (*entity.UnitOfMeasure, error) {
	return uomApp.uomRepo.Update(id, uom)
}
