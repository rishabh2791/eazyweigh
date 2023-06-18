package persistance

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
)

type MaterialRepo struct {
	DB     *gorm.DB
	Logger hclog.Logger
}

var _ repository.MaterialRepository = &MaterialRepo{}

func NewMaterialRepo(db *gorm.DB, logger hclog.Logger) *MaterialRepo {
	return &MaterialRepo{
		DB:     db,
		Logger: logger,
	}
}

func (materialRepo *MaterialRepo) Create(material *entity.Material) (*entity.Material, error) {
	validationErr := material.Validate()
	if validationErr != nil {
		return nil, validationErr
	}

	creationErr := materialRepo.DB.Create(&material).Error
	if creationErr != nil {
		return nil, creationErr
	}

	return material, nil
}

func (materialRepo *MaterialRepo) Get(id string) (*entity.Material, error) {
	material := entity.Material{}

	getErr := materialRepo.DB.Where("id = ?", id).Take(&material).Error
	if getErr != nil {
		return nil, getErr
	}

	return &material, nil
}

func (materialRepo *MaterialRepo) List(conditions string) ([]entity.Material, error) {
	materials := []entity.Material{}

	getErr := materialRepo.DB.Where(conditions).Find(&materials).Error
	if getErr != nil {
		return nil, getErr
	}

	return materials, nil
}

func (materialRepo *MaterialRepo) Update(id string, update *entity.Material) (*entity.Material, error) {
	existingMaterial := entity.Material{}

	getErr := materialRepo.DB.Where("id = ?", id).Take(&existingMaterial).Error
	if getErr != nil {
		return nil, getErr
	}

	updationErr := materialRepo.DB.Table(entity.Material{}.Tablename()).Where("id = ?", id).Updates(update).Error
	if updationErr != nil {
		return nil, updationErr
	}

	updated := entity.Material{}
	materialRepo.DB.Where("id = ?", id).Take(&updated)

	return &updated, nil
}
