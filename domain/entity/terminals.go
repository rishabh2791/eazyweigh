package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Terminal struct {
	value_objects.BaseModel
	ID                string         `json:"id" gorm:"size:191;not null;primaryKey;unique;"`
	Description       string         `json:"description" gorm:"size:200;not null;"`
	FactoryID         string         `json:"factory_id" gorm:"size:191;not null;"`
	APIKey            string         `json:"api_key" gorm:"size:200;not null;unique;"`
	MACAddress        string         `json:"mac_address" gorm:"size:20;not null;unique;"`
	Capacity          float32        `json:"capacity"`
	UnitOfMeasureID   string         `json:"unit_of_measurement_id" gorm:"size:191;not null;"`
	UnitOfMeasure     *UnitOfMeasure `json:"unit_of_measurement"`
	LeastCount        float32        `json:"least_count"`
	Occupied          bool           `json:"occupied" gorm:"default:false;"` // changes to occupied once a user logs in on one terminal
	CreatedByUsername string         `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User          `json:"created_by"`
	UpdatedByUsername string         `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User          `json:"updated_by"`
}

func (Terminal) Tablename() string {
	return "terminals"
}

func (terminal *Terminal) BeforeCreate(db *gorm.DB) error {
	terminal.ID = uuid.New().String()
	terminal.APIKey = uuid.New().String()
	return nil
}

func (terminal *Terminal) Validate() error {
	errors := map[string]interface{}{}
	if terminal.Description == "" || len(terminal.Description) == 0 {
		errors["description"] = "Terminal Description Required."
	}
	if terminal.MACAddress == "" || len(terminal.MACAddress) == 0 {
		errors["mac_address"] = "MAC Address Required."
	}
	if terminal.UnitOfMeasureID == "" || len(terminal.UnitOfMeasureID) == 0 {
		errors["uom_required"] = "Unit of Measure Required."
	}
	if terminal.CreatedByUsername == "" || len(terminal.CreatedByUsername) == 0 {
		errors["created_by"] = "Created By Required."
	}
	if terminal.UpdatedByUsername == "" || len(terminal.UpdatedByUsername) == 0 {
		errors["updated_by"] = "Updated By Required."
	}
	return utilities.ConvertMapToError(errors)
}
