package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"

	"github.com/google/uuid"
	"gorm.io/gorm"
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

func (uom *UnitOfMeasure) BeforeCreate(db *gorm.DB) error {
	uom.ID = uuid.New().String()
	return nil
}

func (uom *UnitOfMeasure) Validate() error {
	errors := map[string]interface{}{}
	if uom.FactoryID == "" || len(uom.FactoryID) == 0 {
		errors["factory"] = "Factory Required."
	}
	if uom.Code == "" || len(uom.Code) == 0 {
		errors["code"] = "Short Text for Unit of Measure Required."
	}
	if uom.Description == "" || len(uom.Description) == 0 {
		errors["description"] = "Description for Unit of Measure Required."
	}
	if uom.FactoryID == "" || len(uom.FactoryID) == 0 {
		errors["factory"] = "Factory Required."
	}
	if uom.CreatedByUsername == "" || len(uom.CreatedByUsername) == 0 {
		errors["created_by"] = "Created By Required."
	}
	if uom.UpdatedByUsername == "" || len(uom.UpdatedByUsername) == 0 {
		errors["updated_by"] = "Updated By Required."
	}
	return utilities.ConvertMapToError(errors)
}
