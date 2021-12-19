package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"
)

type UnitOfMeasureConversion struct {
	value_objects.BaseModel
	UnitOfMeasure1ID  string         `json:"unit1_id" gorm:"size:191;not null;"`
	UnitOfMeasure1    *UnitOfMeasure `json:"unit1"`
	Value1            float32        `json:"value1"`
	UnitOfMeasure2ID  string         `json:"unit2_id" gorm:"size:191;not null;"`
	UnitOfMeasure2    *UnitOfMeasure `json:"unit2"`
	Value2            float32        `json:"value2"`
	CreatedByUsername string         `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User          `json:"created_by"`
	UpdatedByUsername string         `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User          `json:"updated_by"`
}

func (UnitOfMeasureConversion) Tablename() string {
	return "unit_of_measure_conversions"
}

func (conversion *UnitOfMeasureConversion) Validate() error {
	errors := map[string]interface{}{}
	return utilities.ConvertMapToError(errors)
}
