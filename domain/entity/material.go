package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"
)

type Material struct {
	value_objects.BaseModel
	ID                string         `json:"id" gorm:"size:191;not null;unique;primaryKey;"`
	FactoryID         string         `json:"factory_id" gorm:"size:191;not null;uniqueIndex:factory_material;"`
	Code              string         `json:"code" gorm:"size:20;not null;uniqueIndex:factory_material;"`
	Description       string         `json:"description" gorm:"size:200;not null;"`
	UnitOfMeasureID   string         `json:"unit_of_measurement_id" gorm:"size:191;not null;"`
	UnitOfMeasure     *UnitOfMeasure `json:"unit_of_measurement"`
	CreatedByUsername string         `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User          `json:"created_by"`
	UpdatedByUsername string         `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User          `json:"updated_by"`
}

func (Material) Tablename() string {
	return "materials"
}

func (material *Material) Validate() error {
	errors := map[string]interface{}{}
	return utilities.ConvertMapToError(errors)
}
