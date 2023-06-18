package persistance

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
)

type ScannedDataRepo struct {
	DB     *gorm.DB
	Logger hclog.Logger
}

var _ repository.ScannedDataRepository = &ScannedDataRepo{}

func NewScannedDataRepo(db *gorm.DB, logger hclog.Logger) *ScannedDataRepo {
	return &ScannedDataRepo{
		DB:     db,
		Logger: logger,
	}
}

func (scannedDataRepo *ScannedDataRepo) Create(scannedData *entity.ScannedData) (*entity.ScannedData, error) {
	validationErr := scannedData.Validate()
	if validationErr != nil {
		return nil, validationErr
	}

	creationErr := scannedDataRepo.DB.Create(&scannedData).Error
	if creationErr != nil {
		return nil, creationErr
	}

	return scannedData, nil
}

func (scannedDataRepo *ScannedDataRepo) Get(id string) (*entity.ScannedData, error) {
	scannedData := entity.ScannedData{}

	getErr := scannedDataRepo.DB.Where("id = ?", id).Take(&scannedData).Error
	if getErr != nil {
		return nil, getErr
	}

	return &scannedData, nil
}

func (scannedDataRepo *ScannedDataRepo) List(conditions string) ([]entity.ScannedData, error) {
	scannedData := []entity.ScannedData{}

	getErr := scannedDataRepo.DB.Where(conditions).Find(&scannedData).Error
	if getErr != nil {
		return nil, getErr
	}

	return scannedData, nil
}
