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
	bomItemRepo *BOMItemRepo
}

var _ repository.BOMRepository = &BOMRepo{}

func NewBOMRepo(db *gorm.DB, logger hclog.Logger, bomItemRepo *BOMItemRepo) *BOMRepo {
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

	getErr := bomRepo.DB.Preload(clause.Associations).Where(conditions).Find(&boms).Error
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

	updationErr := bomRepo.DB.Table(entity.BOM{}.Tablename()).Updates(update).Error
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
