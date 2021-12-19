package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"
)

type BOMItem struct {
	value_objects.BaseModel
	ID                string         `json:"id" gorm:"size:191;not null;unique;primaryKey;"`
	BOMID             string         `json:"bom_id" gorm:"size:191;not null;uniqueIndex:bom_material;"`
	MaterialID        string         `json:"material_code" gorm:"size:191;not null;uniqueIndex:bom_material;"`
	Material          *Material      `json:"material"`
	Quantity          float32        `json:"quantity"`
	Tolerance         float32        `json:"tolerance"`
	OverIssue         bool           `json:"over_issue" gorm:"default:false;"`
	UnderIssue        bool           `json:"under_issue" gorm:"default:false;"`
	UnitOfMeasureID   string         `json:"unit_of_measurement_id" gorm:"size:191;not null;"`
	UnitOfMeasure     *UnitOfMeasure `json:"unit_of_measurement"`
	CreatedByUsername string         `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User          `json:"created_by"`
	UpdatedByUsername string         `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User          `json:"updated_by"`
}

func (BOMItem) Tablename() string {
	return "bom_items"
}

func (bomItem *BOMItem) Validate() error {
	errors := map[string]interface{}{}
	return utilities.ConvertMapToError(errors)
}
