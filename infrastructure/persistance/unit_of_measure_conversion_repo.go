package persistance

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UnitOfMeasureConversionRepo struct {
	DB     *gorm.DB
	Logger hclog.Logger
}

var _ repository.UnitOfMeasureConversionRepository = &UnitOfMeasureConversionRepo{}

func NewUnitOfMeasureConversionRepo(db *gorm.DB, logger hclog.Logger) *UnitOfMeasureConversionRepo {
	return &UnitOfMeasureConversionRepo{
		DB:     db,
		Logger: logger,
	}
}

func (conversionRepo *UnitOfMeasureConversionRepo) Create(conversion *entity.UnitOfMeasureConversion) (*entity.UnitOfMeasureConversion, error) {
	validationErr := conversion.Validate()
	if validationErr != nil {
		return nil, validationErr
	}

	existingConversion := entity.UnitOfMeasureConversion{}
	getErr := conversionRepo.DB.Preload(clause.Associations).Where("unit_of_measure1_id = ? AND unit_of_measure2_id = ?", conversion.UnitOfMeasure1ID, conversion.UnitOfMeasure2ID).Take(&existingConversion).Error
	if getErr != nil {
		creationErr := conversionRepo.DB.Create(&conversion).Error
		if creationErr != nil {
			return nil, creationErr
		}
	} else {
		updationErr := conversionRepo.DB.Table(entity.UnitOfMeasureConversion{}.Tablename()).Where("unit_of_measure1_id = ? AND unit_of_measure2_id = ?", conversion.UnitOfMeasure1ID, conversion.UnitOfMeasure2ID).Updates(&conversion).Error
		if updationErr != nil {
			return nil, updationErr
		}
	}

	return conversion, nil
}

func (conversionRepo *UnitOfMeasureConversionRepo) Get(id string) (*entity.UnitOfMeasureConversion, error) {
	conversion := entity.UnitOfMeasureConversion{}

	getErr := conversionRepo.DB.Preload(clause.Associations).Where("id = ?", id).Take(&conversion).Error
	if getErr != nil {
		return nil, getErr
	}

	return &conversion, nil
}

func (conversionRepo *UnitOfMeasureConversionRepo) List(conditions string) ([]entity.UnitOfMeasureConversion, error) {
	conversions := []entity.UnitOfMeasureConversion{}

	getErr := conversionRepo.DB.
		Preload("UnitOfMeasure1.Factory").
		Preload("UnitOfMeasure1.Factory.Address").
		Preload("UnitOfMeasure1.Factory.CreatedBy").
		Preload("UnitOfMeasure1.Factory.CreatedBy.UserRole").
		Preload("UnitOfMeasure1.Factory.UpdatedBy").
		Preload("UnitOfMeasure1.Factory.UpdatedBy.UserRole").
		Preload("UnitOfMeasure1.CreatedBy").
		Preload("UnitOfMeasure1.CreatedBy.UserRole").
		Preload("UnitOfMeasure1.UpdatedBy").
		Preload("UnitOfMeasure1.UpdatedBy.UserRole").
		Preload("UnitOfMeasure2.Factory").
		Preload("UnitOfMeasure2.Factory.Address").
		Preload("UnitOfMeasure2.Factory.CreatedBy").
		Preload("UnitOfMeasure2.Factory.CreatedBy.UserRole").
		Preload("UnitOfMeasure2.Factory.UpdatedBy").
		Preload("UnitOfMeasure2.Factory.UpdatedBy.UserRole").
		Preload("UnitOfMeasure2.CreatedBy").
		Preload("UnitOfMeasure2.CreatedBy.UserRole").
		Preload("UnitOfMeasure2.UpdatedBy").
		Preload("UnitOfMeasure2.UpdatedBy.UserRole").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).Where(conditions).Find(&conversions).Error
	if getErr != nil {
		return nil, getErr
	}

	return conversions, nil
}
