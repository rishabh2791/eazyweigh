package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
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

func (jobAssignment *JobAssignment) BeforeCreate(db *gorm.DB) error {
	jobAssignment.ID = uuid.New().String()
	return nil
}

func (job *JobAssignment) Validate() error {
	errors := map[string]interface{}{}
	if job.JobID == "" || len(job.JobID) == 0 {
		errors["job_id"] = "Job Details Missing."
	}
	if job.ShiftScheduleID == "" || len(job.ShiftScheduleID) == 0 {
		errors["shift_id"] = "Shift Details Missing."
	}
	if job.CreatedByUsername == "" || len(job.CreatedByUsername) == 0 {
		errors["created_by"] = "Created By Details Missing."
	}
	if job.UpdatedByUsername == "" || len(job.UpdatedByUsername) == 0 {
		errors["updated_by"] = "Updated By Details Missing."
	}
	return utilities.ConvertMapToError(errors)
}
