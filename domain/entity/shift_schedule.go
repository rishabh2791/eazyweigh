package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"
	"time"
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

func (shiftSchedule *ShiftSchedule) Validate() error {
	errors := map[string]interface{}{}
	return utilities.ConvertMapToError(errors)
}
