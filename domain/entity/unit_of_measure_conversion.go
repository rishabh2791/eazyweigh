package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UnitOfMeasureConversion struct {
	value_objects.BaseModel
	ID                string         `json:"id" gorm:"size:191;not null;primaryKey;unique;"`
	FactoryID         string         `json:"factory_id"`
	UnitOfMeasure1ID  string         `json:"unit1_id" gorm:"size:191;not null;uniqueIndex:factory_uom1_uom2;"`
	UnitOfMeasure1    *UnitOfMeasure `json:"unit1"`
	Value1            float32        `json:"value1"`
	UnitOfMeasure2ID  string         `json:"unit2_id" gorm:"size:191;not null;uniqueIndex:factory_uom1_uom2;"`
	UnitOfMeasure2    *UnitOfMeasure `json:"unit2"`
	Value2            float32        `json:"value2"`
	CreatedByUsername string         `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User          `json:"created_by"`
	UpdatedByUsername string         `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User          `json:"updated_by"`
}

func (uomConversion *UnitOfMeasureConversion) BeforeCreate(db *gorm.DB) error {
	uomConversion.ID = uuid.New().String()
	return nil
}

func (UnitOfMeasureConversion) Tablename() string {
	return "unit_of_measure_conversions"
}

func (conversion *UnitOfMeasureConversion) Validate() error {
	errors := map[string]interface{}{}
	if conversion.UnitOfMeasure1ID == "" || len(conversion.UnitOfMeasure1ID) == 0 {
		errors["uom1"] = "Unit of Measure 1 Required."
	}
	if conversion.UnitOfMeasure1ID == "" || len(conversion.UnitOfMeasure2ID) == 0 {
		errors["uom2"] = "Unit of Measure 2 Required."
	}
	if conversion.CreatedByUsername == "" || len(conversion.CreatedByUsername) == 0 {
		errors["created_by"] = "Created By Required."
	}
	if conversion.UpdatedByUsername == "" || len(conversion.UpdatedByUsername) == 0 {
		errors["updated_by"] = "Updated By Required."
	}
	return utilities.ConvertMapToError(errors)
}
