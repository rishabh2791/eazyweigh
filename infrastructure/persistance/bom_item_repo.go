package persistance

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BOMItemRepo struct {
	DB     *gorm.DB
	Logger hclog.Logger
}

var _ repository.BOMItemsRepository = &BOMItemRepo{}

func NewBOMItemRepo(db *gorm.DB, logger hclog.Logger) *BOMItemRepo {
	return &BOMItemRepo{
		DB:     db,
		Logger: logger,
	}
}

func (bomItemRepo *BOMItemRepo) Create(bomItems *entity.BOMItem) (*entity.BOMItem, error) {
	validationErr := bomItems.Validate()
	if validationErr != nil {
		return nil, validationErr
	}

	creationErr := bomItemRepo.DB.Create(&bomItems).Error
	if creationErr != nil {
		return nil, creationErr
	}

	return bomItems, nil
}

func (bomItemRepo *BOMItemRepo) Get(id string) (*entity.BOMItem, error) {
	bomItem := entity.BOMItem{}

	getErr := bomItemRepo.DB.Preload(clause.Associations).Where("id = ?", id).Take(&bomItem).Error
	if getErr != nil {
		return nil, getErr
	}

	return &bomItem, nil
}

func (bomItemRepo *BOMItemRepo) List(conditions string) ([]entity.BOMItem, error) {
	bomItems := []entity.BOMItem{}

	getErr := bomItemRepo.DB.Preload("Material.UnitOfMeasure").
		Preload("Material.UnitOfMeasure.Factory").
		Preload("Material.UnitOfMeasure.Factory.Address").
		Preload("Material.UnitOfMeasure.Factory.CreatedBy").
		Preload("Material.UnitOfMeasure.Factory.CreatedBy.UserRole").
		Preload("Material.UnitOfMeasure.UpdatedBy").
		Preload("Material.UnitOfMeasure.UpdatedBy.UserRole").
		Preload("Material.UnitOfMeasure.CreatedBy").
		Preload("Material.UnitOfMeasure.CreatedBy.UserRole").
		Preload("Material.UnitOfMeasure.UpdatedBy").
		Preload("Material.UnitOfMeasure.UpdatedBy.UserRole").
		Preload("Material.CreatedBy").
		Preload("Material.CreatedBy.UserRole").
		Preload("Material.UpdatedBy").
		Preload("Material.UpdatedBy.UserRole").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).Where(conditions).Find(&bomItems).Error
	if getErr != nil {
		return nil, getErr
	}

	return bomItems, nil
}

func (bomItemRepo *BOMItemRepo) Update(id string, update *entity.BOMItem) (*entity.BOMItem, error) {
	bomItem := entity.BOMItem{}

	getErr := bomItemRepo.DB.Preload(clause.Associations).Take(&bomItem).Error
	if getErr != nil {
		return nil, getErr
	}

	updationErr := bomItemRepo.DB.Table(entity.BOMItem{}.Tablename()).Where("id = ?", id).Updates(update).Error
	if updationErr != nil {
		return nil, updationErr
	}

	updated := entity.BOMItem{}
	bomItemRepo.DB.Preload(clause.Associations).Take(&updated)

	return &updated, nil
}
