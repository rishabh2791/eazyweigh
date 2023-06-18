package persistance

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
)

type BatchRepo struct {
	DB     *gorm.DB
	Logger hclog.Logger
}

var _ repository.BatchRepository = &BatchRepo{}

func NewBatchRepo(db *gorm.DB, logger hclog.Logger) *BatchRepo {
	return &BatchRepo{
		DB:     db,
		Logger: logger,
	}
}

func (batchRepo *BatchRepo) Create(batch *entity.Batch) (*entity.Batch, error) {
	validationErr := batch.Validate()
	if validationErr != nil {
		return nil, validationErr
	}

	creationErr := batchRepo.DB.Create(&batch).Error
	return batch, creationErr
}

func (batchRepo *BatchRepo) Get(id string) (*entity.Batch, error) {
	batch := entity.Batch{}

	getErr := batchRepo.DB.Where("id = ?", id).First(&batch).Error
	return &batch, getErr
}

func (batchRepo *BatchRepo) List(conditions string) ([]entity.Batch, error) {
	batches := []entity.Batch{}

	getErr := batchRepo.DB.Where(conditions).Find(&batches).Error
	return batches, getErr
}

func (batchRepo *BatchRepo) Update(id string, batch *entity.Batch) (*entity.Batch, error) {
	existingBatch := entity.Batch{}

	getErr := batchRepo.DB.Where("id = ?", id).First(&existingBatch).Error
	if getErr != nil {
		return nil, getErr
	}

	updationErr := batchRepo.DB.Table(batch.Tablename()).Updates(&batch).Error
	if updationErr != nil {
		return nil, updationErr
	}

	updated := entity.Batch{}
	batchRepo.DB.Where("id = ?", id).First(&updated)
	return &updated, nil
}
