package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"
	"time"
)

type JobAssignment struct {
	value_objects.BaseModel
	ID                string         `json:"id" gorm:"size:191;not null;unique;primaryKey;"`
	JobID             string         `json:"job_id" gorm:"size:191;not null;"`
	Job               *Job           `json:"job"`
	ShiftScheduleID   string         `json:"shift_schedule_id" gorm:"size:191;not null;"`
	ShiftSchedule     *ShiftSchedule `json:"shift_schedule"`
	StartTime         time.Time      `json:"start_time"`
	EndTime           time.Time      `json:"end_time"`
	CreatedByUsername string         `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User          `json:"created_by"`
	UpdatedByUsername string         `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User          `json:"updated_by"`
}

func (JobAssignment) Tablename() string {
	return "job_assignments"
}

func (job *JobAssignment) Validate() error {
	errors := map[string]interface{}{}
	return utilities.ConvertMapToError(errors)
}
