package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
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

func (shift *Shift) BeforeCreate(db *gorm.DB) error {
	shift.ID = uuid.New().String()
	return nil
}

func (shift *Shift) Validate() error {
	errors := map[string]interface{}{}
	if shift.FactoryID == "" || len(shift.FactoryID) == 0 {
		errors["factory"] = "Factory Required."
	}
	if shift.Code == "" || len(shift.Code) == 0 {
		errors["code"] = "Short Text for Shift Required Required."
	}
	if shift.Description == "" || len(shift.Description) == 0 {
		errors["description"] = "Description for Shift Required."
	}
	if shift.StartTime.String() == "" || len(shift.StartTime.String()) == 0 {
		errors["start_time"] = "Start Time Required."
	}
	if shift.EndTime.String() == "" || len(shift.EndTime.String()) == 0 {
		errors["end_time"] = "End Time Required."
	}
	if shift.CreatedByUsername == "" || len(shift.CreatedByUsername) == 0 {
		errors["created_by"] = "Created By Required."
	}
	if shift.UpdatedByUsername == "" || len(shift.UpdatedByUsername) == 0 {
		errors["updated_by"] = "Updated By Required."
	}
	return utilities.ConvertMapToError(errors)
}
