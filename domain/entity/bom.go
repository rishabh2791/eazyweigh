package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"
)

type BOM struct {
	value_objects.BaseModel
	ID                string         `json:"id" gorm:"size:191;not null;unique;primaryKey;"`
	MaterialID        string         `json:"material_code" gorm:"size:191;not null;"`
	Material          *Material      `json:"material"`
	BOMItems          []BOMItem      `json:"bom_items"`
	UnitSize          float32        `json:"unit_size" gorm:"default:1000;"`
	UnitOfMeasureID   string         `json:"unit_of_measurement_id" gorm:"size:191;not null;"`
	UnitOfMeasure     *UnitOfMeasure `json:"unit_of_measurement"`
	CreatedByUsername string         `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User          `json:"created_by"`
	UpdatedByUsername string         `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User          `json:"updated_by"`
}

func (BOM) Tablename() string {
	return "boms"
}

func (bom *BOM) Validate() error {
	errors := map[string]interface{}{}
	return utilities.ConvertMapToError(errors)
}
