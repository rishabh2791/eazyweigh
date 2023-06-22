package persistance

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BOMRepo struct {
	DB          *gorm.DB
	Logger      hclog.Logger
	bomItemRepo repository.BOMItemsRepository
}

var _ repository.BOMRepository = &BOMRepo{}

func NewBOMRepo(db *gorm.DB, logger hclog.Logger, bomItemRepo repository.BOMItemsRepository) *BOMRepo {
	return &BOMRepo{
		DB:          db,
		Logger:      logger,
		bomItemRepo: bomItemRepo,
	}
}

func (bomRepo *BOMRepo) Create(bom *entity.BOM) (*entity.BOM, error) {
	validationErr := bom.Validate()
	if validationErr != nil {
		return nil, validationErr
	}

	creationErr := bomRepo.DB.Create(&bom).Error
	if creationErr != nil {
		return nil, creationErr
	}

	return bom, nil
}

func (bomRepo *BOMRepo) Get(id string) (*entity.BOM, error) {
	bom := entity.BOM{}

	getErr := bomRepo.DB.Preload(clause.Associations).Where("id = ?", id).Take(&bom).Error
	if getErr != nil {
		return nil, getErr
	}

	return &bom, nil
}

func (bomRepo *BOMRepo) List(conditions string) ([]entity.BOM, error) {
	boms := []entity.BOM{}

	getErr := bomRepo.DB.Preload("Material.").
		Preload("Material.UnitOfMeasure").
		Preload("Material.UnitOfMeasure.Factory").
		Preload("Material.UnitOfMeasure.Factory.Address").
		Preload("Material.UnitOfMeasure.Factory.CreatedBy").
		Preload("Material.UnitOfMeasure.Factory.UpdatedBy").
		Preload("Material.UnitOfMeasure.Factory.CreatedBy.UserRole").
		Preload("Material.UnitOfMeasure.Factory.UpdatedBy.UserRole").
		Preload("Material.UnitOfMeasure.CreatedBy").
		Preload("Material.UnitOfMeasure.UpdatedBy").
		Preload("Material.UnitOfMeasure.CreatedBy.UserRole").
		Preload("Material.UnitOfMeasure.UpdatedBy.UserRole").
		Preload("Material.CreatedBy").
		Preload("Material.CreatedBy.UserRole").
		Preload("Material.UpdatedBy").
		Preload("Material.UpdatedBy.UserRole").
		Preload("BOMItems.Material.").
		Preload("BOMItems.Material.UnitOfMeasure").
		Preload("BOMItems.Material.UnitOfMeasure.Factory").
		Preload("BOMItems.Material.UnitOfMeasure.Factory.Address").
		Preload("BOMItems.Material.UnitOfMeasure.Factory.CreatedBy").
		Preload("BOMItems.Material.UnitOfMeasure.Factory.UpdatedBy").
		Preload("BOMItems.Material.UnitOfMeasure.Factory.CreatedBy.UserRole").
		Preload("BOMItems.Material.UnitOfMeasure.Factory.UpdatedBy.UserRole").
		Preload("BOMItems.Material.UnitOfMeasure.CreatedBy").
		Preload("BOMItems.Material.UnitOfMeasure.UpdatedBy").
		Preload("BOMItems.Material.UnitOfMeasure.CreatedBy.UserRole").
		Preload("BOMItems.Material.UnitOfMeasure.UpdatedBy.UserRole").
		Preload("BOMItems.Material.CreatedBy").
		Preload("BOMItems.Material.CreatedBy.UserRole").
		Preload("BOMItems.Material.UpdatedBy").
		Preload("BOMItems.Material.UpdatedBy.UserRole").
		Preload("BOMItems.UnitOfMeasure").
		Preload("BOMItems.UnitOfMeasure.Factory").
		Preload("BOMItems.UnitOfMeasure.Factory.Address").
		Preload("BOMItems.UnitOfMeasure.Factory.CreatedBy").
		Preload("BOMItems.UnitOfMeasure.Factory.UpdatedBy").
		Preload("BOMItems.UnitOfMeasure.Factory.CreatedBy.UserRole").
		Preload("BOMItems.UnitOfMeasure.Factory.UpdatedBy.UserRole").
		Preload("BOMItems.UnitOfMeasure.CreatedBy").
		Preload("BOMItems.UnitOfMeasure.UpdatedBy").
		Preload("BOMItems.UnitOfMeasure.CreatedBy.UserRole").
		Preload("BOMItems.UnitOfMeasure.UpdatedBy.UserRole").
		Preload("BOMItems.CreatedBy").
		Preload("BOMItems.UpdatedBy").
		Preload("BOMItems.CreatedBy.UserRole").
		Preload("BOMItems.UpdatedBy.UserRole").
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
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).Where(conditions).Find(&boms).Error
	if getErr != nil {
		return nil, getErr
	}

	return boms, nil
}

func (bomRepo *BOMRepo) Update(id string, update *entity.BOM) (*entity.BOM, error) {
	existingBOM := entity.BOM{}

	getErr := bomRepo.DB.Preload(clause.Associations).Where("id = ?", id).Take(&existingBOM).Error
	if getErr != nil {
		return nil, getErr
	}

	updationErr := bomRepo.DB.Table(entity.BOM{}.Tablename()).Where("id = ?", id).Updates(update).Error
	if updationErr != nil {
		return nil, updationErr
	}

	if len(update.BOMItems) != 0 {
		for _, bomItem := range update.BOMItems {
			_, err := bomRepo.bomItemRepo.Update(bomItem.ID, &bomItem)
			if err != nil {
				return nil, err
			}
		}
	}

	updated := entity.BOM{}
	bomRepo.DB.Preload(clause.Associations).Where("id = ?", id).Take(&updated)

	return &updated, nil
}
