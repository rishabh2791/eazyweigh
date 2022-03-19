package application

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"
)

type UnitOfMeasureConversionApp struct {
	conversionRepo repository.UnitOfMeasureConversionRepository
}

var _ UnitOfMeasureConversionAppInterface = &UnitOfMeasureConversionApp{}

func NewUnitOfMeasureConversionApp(conversionRepo repository.UnitOfMeasureConversionRepository) *UnitOfMeasureConversionApp {
	return &UnitOfMeasureConversionApp{
		conversionRepo: conversionRepo,
	}
}

type UnitOfMeasureConversionAppInterface interface {
	Create(conversion *entity.UnitOfMeasureConversion) (*entity.UnitOfMeasureConversion, error)
	Get(id string) (*entity.UnitOfMeasureConversion, error)
	List(conditions string) ([]entity.UnitOfMeasureConversion, error)
}

func (conversionApp *UnitOfMeasureConversionApp) Create(conversion *entity.UnitOfMeasureConversion) (*entity.UnitOfMeasureConversion, error) {
	return conversionApp.conversionRepo.Create(conversion)
}

func (conversionApp *UnitOfMeasureConversionApp) Get(id string) (*entity.UnitOfMeasureConversion, error) {
	return conversionApp.conversionRepo.Get(id)
}

func (conversionApp *UnitOfMeasureConversionApp) List(conditions string) ([]entity.UnitOfMeasureConversion, error) {
	return conversionApp.conversionRepo.List(conditions)
}
