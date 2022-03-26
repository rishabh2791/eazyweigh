package persistance

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

	getErr := materialRepo.DB.Preload(clause.Associations).Where("id = ?", id).Take(&material).Error
	if getErr != nil {
		return nil, getErr
	}

	return &material, nil
}

func (materialRepo *MaterialRepo) List(conditions string) ([]entity.Material, error) {
	materials := []entity.Material{}

	getErr := materialRepo.DB.
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload("UnitOfMeasure.Factory").
		Preload("UnitOfMeasure.Factory.Address").
		Preload("UnitOfMeasure.Factory.CreatedBy").
		Preload("UnitOfMeasure.Factory.UpdatedBy").
		Preload("UnitOfMeasure.Factory.CreatedBy.UserRole").
		Preload("UnitOfMeasure.Factory.UpdatedBy.UserRole").
		Preload("UnitOfMeasure.CreatedBy").
		Preload("UnitOfMeasure.UpdatedBy").
		Preload("UnitOfMeasure.CreatedBy.UserRole").
		Preload("UnitOfMeasure.UpdatedBy.UserRole").
		Preload(clause.Associations).Where(conditions).Find(&materials).Error
	if getErr != nil {
		return nil, getErr
	}

	return materials, nil
}

func (materialRepo *MaterialRepo) Update(id string, update *entity.Material) (*entity.Material, error) {
	existingMaterial := entity.Material{}

	getErr := materialRepo.DB.Preload(clause.Associations).Where("id = ?", id).Take(&existingMaterial).Error
	if getErr != nil {
		return nil, getErr
	}

	updationErr := materialRepo.DB.Table(entity.Material{}.Tablename()).Where("id = ?", id).Updates(update).Error
	if updationErr != nil {
		return nil, updationErr
	}

	updated := entity.Material{}
	materialRepo.DB.Preload(clause.Associations).Take(&updated)

	return &updated, nil
}
