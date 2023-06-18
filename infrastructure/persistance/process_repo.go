package persistance

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
)

type ProcessRepo struct {
	DB     *gorm.DB
	Logger hclog.Logger
}

var _ repository.ProcessRepository = &ProcessRepo{}

func NewProcessRepo(db *gorm.DB, logger hclog.Logger) *ProcessRepo {
	return &ProcessRepo{
		DB:     db,
		Logger: logger,
	}
}

func (processRepo *ProcessRepo) Create(process *entity.Process) (*entity.Process, error) {
	validationErr := process.Validate()
	if validationErr != nil {
		return nil, validationErr
	}

	existingProcess, _ := processRepo.List("material_id = '" + process.MaterialID + "'")
	if len(existingProcess) != 0 {
		revisionNumber := len(existingProcess) + 1
		process.Version = revisionNumber
	}
	creationErr := processRepo.DB.Create(&process).Error
	return process, creationErr
}

func (processRepo *ProcessRepo) Get(materialID string) (*entity.Process, error) {
	process := entity.Process{}

	getErr := processRepo.DB.Where("material_id = ?", materialID).First(&process).Error
	return &process, getErr
}

func (processRepo *ProcessRepo) List(conditions string) ([]entity.Process, error) {
	processes := []entity.Process{}

	getErr := processRepo.DB.Where(conditions).Find(&processes).Error
	return processes, getErr
}

func (processRepo *ProcessRepo) Update(id string, process *entity.Process) (*entity.Process, error) {
	existingProcess := entity.Process{}

	getErr := processRepo.DB.Where("id = ?", id).First(&existingProcess).Error
	if getErr != nil {
		return nil, getErr
	}

	updationErr := processRepo.DB.Table(process.Tablename()).Updates(&process).Error
	if updationErr != nil {
		return nil, updationErr
	}

	updated := entity.Process{}
	processRepo.DB.Where("id = ?", id).First(&updated)
	return &updated, nil
}
