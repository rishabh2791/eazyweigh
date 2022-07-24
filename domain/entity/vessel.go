package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Vessel struct {
	value_objects.BaseModel
	ID                string   `json:"id" gorm:"size:191;not null;primaryKey;unique;"`
	FactoryID         string   `json:"factory_id" gorm:"size:191;not null;"`
	Factory           *Factory `json:"factory"`
	Name              string   `json:"description" gorm:"size:1000;"`
	CreatedByUsername string   `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User    `json:"created_by"`
	UpdatedByUsername string   `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User    `json:"updated_by"`
}

func (Vessel) Tablename() string {
	return "vessels"
}

func (vessel *Vessel) BeforeCreate(db *gorm.DB) error {
	vessel.ID = uuid.New().String()
	return nil
}

func (vessel *Vessel) Validate() error {
	errors := map[string]interface{}{}
	if vessel.Name == "" || len(vessel.Name) == 0 {
		errors["name"] = "Step Name Required."
	}
	if vessel.FactoryID == "" || len(vessel.FactoryID) == 0 {
		errors["factory"] = "Factory Required."
	}
	if vessel.CreatedByUsername == "" || len(vessel.CreatedByUsername) == 0 {
		errors["created_by"] = "Created By Required."
	}
	if vessel.UpdatedByUsername == "" || len(vessel.UpdatedByUsername) == 0 {
		errors["updated_by"] = "Updated By Required."
	}
	return utilities.ConvertMapToError(errors)
}
