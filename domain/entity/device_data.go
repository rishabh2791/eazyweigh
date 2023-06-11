package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DeviceData struct {
	value_objects.BaseModel
	ID       string  `json:"id" gorm:"size:191;not null;primaryKey;unique;"`
	DeviceID string  `json:"device_id" gorm:"size:191;not null;"`
	Value    float32 `json:"value" gorm:"defalut:0;"`
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
		errors["device"] = "Device Required."
	}

	return utilities.ConvertMapToError(errors)
}
