package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Material struct {
	value_objects.BaseModel
	ID                string         `json:"id" gorm:"size:191;not null;unique;primaryKey;"`
	Type              string         `json:"type" gorm:"size:50;not null;"` // Option of Raw Material or Bulk
	FactoryID         string         `json:"factory_id" gorm:"size:191;not null;uniqueIndex:factory_material;"`
	Code              string         `json:"code" gorm:"size:20;not null;uniqueIndex:factory_material;"`
	Description       string         `json:"description" gorm:"size:200;not null;"`
	UnitOfMeasureID   string         `json:"unit_of_measurement_id" gorm:"size:191;not null;"`
	UnitOfMeasure     *UnitOfMeasure `json:"unit_of_measurement"`
	IsWeighed         bool           `json:"is_weighed" gorm:"default:true;"`
	CreatedByUsername string         `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User          `json:"created_by"`
	UpdatedByUsername string         `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User          `json:"updated_by"`
}

func (Material) Tablename() string {
	return "materials"
}

func (material *Material) BeforeCreate(db *gorm.DB) error {
	material.ID = uuid.New().String()
	return nil
}

func (material *Material) Validate() error {
	errors := map[string]interface{}{}
	if material.FactoryID == "" || len(material.FactoryID) == 0 {
		errors["factory"] = "Factory Details Required."
	}
	if material.Code == "" || len(material.Code) == 0 {
		errors["code"] = "Material Code Required."
	}
	if material.Description == "" || len(material.Description) == 0 {
		errors["description"] = "Material Description Required."
	}
	if material.FactoryID == "" || len(material.FactoryID) == 0 {
		errors["factory"] = "Factory Details Required."
	}
	if material.UnitOfMeasureID == "" || len(material.UnitOfMeasureID) == 0 {
		errors["uom"] = "Unit of Measure Required."
	}
	if material.CreatedByUsername == "" || len(material.CreatedByUsername) == 0 {
		errors["created_by"] = "Created By Required."
	}
	if material.UpdatedByUsername == "" || len(material.UpdatedByUsername) == 0 {
		errors["updated_by"] = "Updated By Required."
	}
	return utilities.ConvertMapToError(errors)
}
