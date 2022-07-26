package application

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"
)

type BatchApp struct {
	batchRepository repository.BatchRepository
}

var _ BatchAppInterface = &BatchApp{}

func NewBatchApp(batchRepository repository.BatchRepository) *BatchApp {
	return &BatchApp{
		batchRepository: batchRepository,
	}
}

func (batchApp *BatchApp) Create(batch *entity.Batch) (*entity.Batch, error) {
	return batchApp.batchRepository.Create(batch)
}

func (batchApp *BatchApp) Get(id string) (*entity.Batch, error) {
	return batchApp.batchRepository.Get(id)
}

func (batchApp *BatchApp) List(conditions string) ([]entity.Batch, error) {
	return batchApp.batchRepository.List(conditions)
}

func (batchApp *BatchApp) Update(id string, batch *entity.Batch) (*entity.Batch, error) {
	return batchApp.batchRepository.Update(id, batch)
}

type BatchAppInterface interface {
	Create(batch *entity.Batch) (*entity.Batch, error)
	Get(id string) (*entity.Batch, error)
	List(conditions string) ([]entity.Batch, error)
	Update(id string, batch *entity.Batch) (*entity.Batch, error)
}
