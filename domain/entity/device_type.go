package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

/*
	Options:
	1.	Anchor Speed
	2.	Cowl Speed
	3.	Paddle Speed
	4.	Inner Speed
	5.	Emulsifier Speed
	6.	Main Vessel Temperature
	7.	Hot Pot Temperature
	8.	Main Vessel Pressure
	9.	Main Vessel Load Cell
	10. Hot Pot Load Cell
*/

type DeviceType struct {
	value_objects.BaseModel
	ID                string   `json:"id" gorm:"size:191;primaryKey;unique;not null;"`
	FactoryID         string   `json:"factory_id" gorm:"size:191;not null;uniqueIndex:factory_device_type;"`
	Factory           *Factory `json:"factory"`
	Description       string   `json:"description" gorm:"size:200;not null;uniqueIndex:factory_device_type;"`
	CreatedByUsername string   `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User    `json:"created_by"`
	UpdatedByUsername string   `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User    `json:"updated_by"`
}

func (DeviceType) Tablename() string {
	return "device_types"
}

func (model *DeviceType) BeforeCreate(db *gorm.DB) error {
	model.ID = uuid.New().String()
	return nil
}

func (model *DeviceType) Validate() error {
	errors := map[string]interface{}{}

	if model.Description == "" || len(model.Description) == 0 {
		errors["description"] = "Device Type Description Can Not be Empty.\n"
	}

	if model.FactoryID == "" || len(model.FactoryID) == 0 {
		errors["factory"] = "Factory Required."
	}
	if model.CreatedByUsername == "" || len(model.CreatedByUsername) == 0 {
		errors["created_by"] = "Created By Required."
	}
	if model.UpdatedByUsername == "" || len(model.UpdatedByUsername) == 0 {
		errors["updated_by"] = "Updated By Required."
	}

	return utilities.ConvertMapToError(errors)
}
