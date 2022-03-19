package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ShiftSchedule struct {
	value_objects.BaseModel
	ID                string    `json:"id" gorm:"size:191;not null;unique;primaryKey;"`
	Date              time.Time `json:"date" gorm:"uniqueIndex:date_shift_user;"`
	ShiftID           string    `json:"shift_id" gorm:"size:191;not null;uniqueIndex:date_shift_user;"`
	Shift             *Shift    `json:"shift"`
	UserUsername      string    `json:"user_username" gorm:"size:20;not null;uniqueIndex:date_shift_user;"`
	User              *User     `json:"user"`
	CreatedByUsername string    `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User     `json:"created_by"`
	UpdatedByUsername string    `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User     `json:"updated_by"`
}

func (ShiftSchedule) Tablename() string {
	return "shift_schedules"
}

func (shiftSchedule *ShiftSchedule) BeforeCreate(db *gorm.DB) error {
	shiftSchedule.ID = uuid.New().String()
	return nil
}

func (shiftSchedule *ShiftSchedule) Validate() error {
	errors := map[string]interface{}{}
	if shiftSchedule.Date.String() == "" || len(shiftSchedule.Date.String()) == 0 {
		errors["date"] = "Schedule Date Required."
	}
	if shiftSchedule.ShiftID == "" || len(shiftSchedule.ShiftID) == 0 {
		errors["shift"] = "Shift Required."
	}
	if shiftSchedule.UserUsername == "" || len(shiftSchedule.UserUsername) == 0 {
		errors["user"] = "Shift User Required."
	}
	if shiftSchedule.CreatedByUsername == "" || len(shiftSchedule.CreatedByUsername) == 0 {
		errors["created_by"] = "Created By Required."
	}
	if shiftSchedule.UpdatedByUsername == "" || len(shiftSchedule.UpdatedByUsername) == 0 {
		errors["updated_by"] = "Updated By Required."
	}
	return utilities.ConvertMapToError(errors)
}
