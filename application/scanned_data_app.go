package application

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"
)

type ScannedDataApp struct {
	scannedDataRepository repository.ScannedDataRepository
}

var _ ScannedDataAppInterface = &ScannedDataApp{}

func NewScannedDataApp(scannedDataRepository repository.ScannedDataRepository) *ScannedDataApp {
	return &ScannedDataApp{
		scannedDataRepository: scannedDataRepository,
	}
}

type ScannedDataAppInterface interface {
	Create(scannedData *entity.ScannedData) (*entity.ScannedData, error)
	Get(id string) (*entity.ScannedData, error)
	List(conditions string) ([]entity.ScannedData, error)
}

func (scannedDataApp *ScannedDataApp) Create(scannedData *entity.ScannedData) (*entity.ScannedData, error) {
	return scannedDataApp.scannedDataRepository.Create(scannedData)
}

func (scannedDataApp *ScannedDataApp) Get(id string) (*entity.ScannedData, error) {
	return scannedDataApp.scannedDataRepository.Get(id)
}

func (scannedDataApp *ScannedDataApp) List(conditions string) ([]entity.ScannedData, error) {
	return scannedDataApp.scannedDataRepository.List(conditions)
}
