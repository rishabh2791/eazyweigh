package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Device struct {
	value_objects.BaseModel
	ID                string       `json:"id" gorm:"size:191;not null;primaryKey;unique;"`
	VesselID          string       `json:"vessel_id" gorm:"size:191;not null;"`
	Vessel            *Vessel      `json:"vessel"`
	Name              string       `json:"name" gorm:"size:200;not null;unique;"`
	CreatedByUsername string       `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User        `json:"created_by"`
	UpdatedByUsername string       `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User        `json:"updated_by"`
	DeviceData        []DeviceData `json:"device_data"`
}

func (Device) Tablename() string {
	return "devices"
}

func (device *Device) BeforeCreate(db *gorm.DB) error {
	device.ID = uuid.New().String()
	return nil
}

func (device *Device) Validate() error {
	errors := map[string]interface{}{}
	if device.Name == "" || len(device.Name) == 0 {
		errors["name"] = "Step Name Required."
	}
	if device.VesselID == "" || len(device.VesselID) == 0 {
		errors["vessel"] = "Vessel Required."
	}
	if device.CreatedByUsername == "" || len(device.CreatedByUsername) == 0 {
		errors["created_by"] = "Created By Required."
	}
	if device.UpdatedByUsername == "" || len(device.UpdatedByUsername) == 0 {
		errors["updated_by"] = "Updated By Required."
	}
	return utilities.ConvertMapToError(errors)
}
