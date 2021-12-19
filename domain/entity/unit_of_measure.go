package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"
)

type UnitOfMeasure struct {
	value_objects.BaseModel
	ID                string   `json:"id" gorm:"size:191;not null;primaryKey;unique;"`
	FactoryID         string   `json:"factory_id" gorm:"size:191;not null;"`
	Factory           *Factory `json:"factory"`
	Code              string   `json:"code" gorm:"size:10;not null;"`
	Description       string   `json:"description" gorm:"size:100;not null;"`
	CreatedByUsername string   `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User    `json:"created_by"`
	UpdatedByUsername string   `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User    `json:"updated_by"`
}

func (UnitOfMeasure) Tablename() string {
	return "unit_of_measures"
}

func (uom *UnitOfMeasure) Validate() error {
	errors := map[string]interface{}{}
	return utilities.ConvertMapToError(errors)
}
