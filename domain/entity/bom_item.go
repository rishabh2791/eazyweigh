package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BOMItem struct {
	value_objects.BaseModel
	ID                string         `json:"id" gorm:"size:191;not null;unique;primaryKey;"`
	BOMID             string         `json:"bom_id" gorm:"size:191;not null;uniqueIndex:bom_material;"`
	MaterialID        string         `json:"material_code" gorm:"size:191;not null;uniqueIndex:bom_material;"`
	Material          *Material      `json:"material"`
	Quantity          float32        `json:"quantity"`
	Tolerance         float32        `json:"tolerance" gorm:"default:0;"`
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

func (bomItem *BOMItem) BeforeCreate(db *gorm.DB) error {
	bomItem.ID = uuid.New().String()
	return nil
}

func (bomItem *BOMItem) Validate() error {
	errors := map[string]interface{}{}
	if bomItem.BOMID == "" || len(bomItem.BOMID) == 0 {
		errors["bom"] = "BOM ID Required."
	}
	if bomItem.MaterialID == "" || len(bomItem.MaterialID) == 0 {
		errors["material"] = "Material Required."
	}
	if bomItem.UnitOfMeasureID == "" || len(bomItem.UnitOfMeasureID) == 0 {
		errors["uom"] = "Unit of Measure Required."
	}
	if bomItem.CreatedByUsername == "" || len(bomItem.CreatedByUsername) == 0 {
		errors["created_by"] = "Created By Required."
	}
	if bomItem.UpdatedByUsername == "" || len(bomItem.UpdatedByUsername) == 0 {
		errors["updated_by"] = "Updated By Required."
	}
	return utilities.ConvertMapToError(errors)
}
