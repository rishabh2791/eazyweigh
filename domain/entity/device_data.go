package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DeviceData struct {
	value_objects.BaseModel
	ID                string  `json:"id" gorm:"size:191;not null;primaryKey;unique;"`
	DeviceID          string  `json:"device_id" gorm:"size:191;not null;"`
	Value             float32 `json:"value" gorm:"defalut:0;"`
	CreatedByUsername string  `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User   `json:"created_by"`
	UpdatedByUsername string  `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User   `json:"updated_by"`
}

func (DeviceData) Tablename() string {
	return "device_data"
}

func (deviceData *DeviceData) BeforeCreate(db *gorm.DB) error {
	deviceData.ID = uuid.New().String()
	return nil
}

func (deviceData *DeviceData) Validate() error {
	errors := map[string]interface{}{}
	if deviceData.DeviceID == "" || len(deviceData.DeviceID) == 0 {
		errors["factory"] = "Factory Required."
	}
	if deviceData.CreatedByUsername == "" || len(deviceData.CreatedByUsername) == 0 {
		errors["created_by"] = "Created By Required."
	}
	if deviceData.UpdatedByUsername == "" || len(deviceData.UpdatedByUsername) == 0 {
		errors["updated_by"] = "Updated By Required."
	}
	return utilities.ConvertMapToError(errors)
}
