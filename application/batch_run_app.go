package application

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"
	"log"
)

type BatchRunApp struct {
	batchRunRepository repository.BatchRunRepository
}

var _ BatchRunAppInterface = &BatchRunApp{}

func NewBatchRunApp(batchRunRepository repository.BatchRunRepository) *BatchRunApp {
	return &BatchRunApp{
		batchRunRepository: batchRunRepository,
	}
}

func (batchApp *BatchRunApp) Create(batch *entity.BatchRun) (*entity.BatchRun, error) {
	log.Println("here")
	return batchApp.batchRunRepository.Create(batch)
}

func (batchApp *BatchRunApp) CreateSuper(batch *entity.BatchRun) (*entity.BatchRun, error) {
	log.Println("here")
	return batchApp.batchRunRepository.CreateSuper(batch)
}

func (batchApp *BatchRunApp) Get(id string) (*entity.BatchRun, error) {
	return batchApp.batchRunRepository.Get(id)
}

func (batchApp *BatchRunApp) List(conditions string) ([]entity.BatchRun, error) {
	return batchApp.batchRunRepository.List(conditions)
}

func (batchApp *BatchRunApp) Update(id string, batch *entity.BatchRun) (*entity.BatchRun, error) {
	return batchApp.batchRunRepository.Update(id, batch)
}

type BatchRunAppInterface interface {
	Create(batch *entity.BatchRun) (*entity.BatchRun, error)
	CreateSuper(batch *entity.BatchRun) (*entity.BatchRun, error)
	Get(id string) (*entity.BatchRun, error)
	List(conditions string) ([]entity.BatchRun, error)
	Update(id string, batch *entity.BatchRun) (*entity.BatchRun, error)
}
