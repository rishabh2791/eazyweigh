package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BatchRun struct {
	value_objects.BaseModel
	ID                string     `json:"id" gorm:"size:191;not null;primaryKey;unique;"`
	JobID             string     `json:"job_id" gorm:"size:191;not null;"`
	Job               *Job       `json:"job"`
	VesselID          string     `json:"vessel_id" gorm:"size:191;not null;"`
	Vessel            *Vessel    `json:"vessel"`
	StartTime         *time.Time `json:"start_time"`
	EndTime           *time.Time `json:"end_time"`
	CreatedByUsername string     `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User      `json:"created_by"`
	UpdatedByUsername string     `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User      `json:"updated_by"`
}

func (BatchRun) Tablename() string {
	return "batch_runs"
}

func (batch *BatchRun) BeforeCreate(db *gorm.DB) error {
	batch.ID = uuid.New().String()
	return nil
}

func (batch *BatchRun) Validate() error {
	errors := map[string]interface{}{}
	if batch.JobID == "" || len(batch.JobID) == 0 {
		errors["job"] = "Job Required."
	}
	if batch.VesselID == "" || len(batch.VesselID) == 0 {
		errors["vessel"] = "Batch Vessel Required."
	}
	if batch.CreatedByUsername == "" || len(batch.CreatedByUsername) == 0 {
		errors["created_by"] = "Created By Required."
	}
	if batch.UpdatedByUsername == "" || len(batch.UpdatedByUsername) == 0 {
		errors["updated_by"] = "Updated By Required."
	}
	return utilities.ConvertMapToError(errors)
}
