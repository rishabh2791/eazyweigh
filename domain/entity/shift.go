package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"
	"time"
)

type Shift struct {
	value_objects.BaseModel
	ID                string    `json:"id" gorm:"size:191;not null;unique;primaryKey;"`
	FactoryID         string    `json:"factory_id" gorm:"size:191;not null;uniqueIndex:factory_shift;"`
	Code              string    `json:"code" gorm:"size:10;not null;uniqueIndex:factory_shift;"`
	Description       string    `json:"description" gorm:"size:200;not null;"`
	StartTime         time.Time `json:"start_time"`
	EndTime           time.Time `json:"end_time"`
	CreatedByUsername string    `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User     `json:"created_by"`
	UpdatedByUsername string    `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User     `json:"updated_by"`
}

func (Shift) Tablename() string {
	return "shifts"
}

func (shift *Shift) Validate() error {
	errors := map[string]interface{}{}
	return utilities.ConvertMapToError(errors)
}
