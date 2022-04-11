package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BOM struct {
	value_objects.BaseModel
	ID                string         `json:"id" gorm:"size:191;not null;unique;primaryKey;"`
	FactoryID         string         `json:"factory_id" gorm:"size:191;not null;uniqueIndex:factory_material_revision;"`
	MaterialID        string         `json:"material_id" gorm:"size:191;not null;uniqueIndex:factory_material_revision;"`
	Material          *Material      `json:"material"`
	BOMItems          []BOMItem      `json:"bom_items"`
	UnitSize          float32        `json:"unit_size" gorm:"default:1000;"`
	UnitOfMeasureID   string         `json:"unit_of_measurement_id" gorm:"size:191;not null;"`
	UnitOfMeasure     *UnitOfMeasure `json:"unit_of_measurement"`
	Revision          int            `json:"revision" gorm:"default:1;uniqueIndex:factory_material_revision;"`
	CreatedByUsername string         `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User          `json:"created_by"`
	UpdatedByUsername string         `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User          `json:"updated_by"`
}

func (BOM) Tablename() string {
	return "boms"
}

func (bom *BOM) BeforeCreate(db *gorm.DB) error {
	if bom.ID == "" || len(bom.ID) == 0 {
		bom.ID = uuid.New().String()
	}
	return nil
}

func (bom *BOM) Validate() error {
	errors := map[string]interface{}{}
	if bom.MaterialID == "" || len(bom.MaterialID) == 0 {
		errors["material"] = "Material Required."
	}
	if bom.UnitOfMeasureID == "" || len(bom.UnitOfMeasureID) == 0 {
		errors["uom"] = "Unit of Measure Required."
	}
	if bom.CreatedByUsername == "" || len(bom.CreatedByUsername) == 0 {
		errors["created_by"] = "Created By Required."
	}
	if bom.UpdatedByUsername == "" || len(bom.UpdatedByUsername) == 0 {
		errors["updated_by"] = "Updated By Required."
	}
	return utilities.ConvertMapToError(errors)
}
