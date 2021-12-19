package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"
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
	CreatedByUsername string         `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User          `json:"created_by"`
	UpdatedByUsername string         `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User          `json:"updated_by"`
}

func (Terminal) Tablename() string {
	return "terminals"
}

func (terminal *Terminal) Validate() error {
	errors := map[string]interface{}{}
	return utilities.ConvertMapToError(errors)
}
