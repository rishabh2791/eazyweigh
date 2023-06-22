package persistance

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UnitOfMeasureRepo struct {
	DB     *gorm.DB
	Logger hclog.Logger
}

var _ repository.UnitOfMeasureRepository = &UnitOfMeasureRepo{}

func NewUnitOfMeasureRepo(db *gorm.DB, logger hclog.Logger) *UnitOfMeasureRepo {
	return &UnitOfMeasureRepo{
		DB:     db,
		Logger: logger,
	}
}

func (uomRepo *UnitOfMeasureRepo) Create(uom *entity.UnitOfMeasure) (*entity.UnitOfMeasure, error) {
	validationErr := uom.Validate()
	if validationErr != nil {
		return nil, validationErr
	}
	creationErr := uomRepo.DB.Create(&uom).Error
	if creationErr != nil {
		return nil, creationErr
	}
	return uom, nil
}

func (uomRepo *UnitOfMeasureRepo) Get(id string) (*entity.UnitOfMeasure, error) {
	uom := entity.UnitOfMeasure{}

	getErr := uomRepo.DB.Preload(clause.Associations).Where("id = ?", id).Take(&uom).Error
	if getErr != nil {
		return nil, getErr
	}

	return &uom, nil
}

func (uomRepo *UnitOfMeasureRepo) List(conditions string) ([]entity.UnitOfMeasure, error) {
	uoms := []entity.UnitOfMeasure{}

	getErr := uomRepo.DB.Preload("Factory.Address").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload("Factory.CreatedBy").
		Preload("Factory.CreatedBy.UserRole").
		Preload("Factory.UpdatedBy").
		Preload("Factory.UpdatedBy.UserRole").
		Preload(clause.Associations).Where(conditions).Find(&uoms).Error
	if getErr != nil {
		return nil, getErr
	}
	return uoms, nil
}

func (uomRepo *UnitOfMeasureRepo) Update(id string, uom *entity.UnitOfMeasure) (*entity.UnitOfMeasure, error) {
	existingUOM := entity.UnitOfMeasure{}

	getErr := uomRepo.DB.Preload(clause.Associations).Where("id = ?", id).Take(&existingUOM).Error
	if getErr != nil {
		return nil, getErr
	}

	updationErr := uomRepo.DB.Table(entity.UnitOfMeasure{}.Tablename()).Where("id = ?", id).Updates(uom).Error
	if updationErr != nil {
		return nil, updationErr
	}

	updated := entity.UnitOfMeasure{}
	uomRepo.DB.Preload(clause.Associations).Where("id = ?", id).Take(&updated)

	return &updated, nil
}
