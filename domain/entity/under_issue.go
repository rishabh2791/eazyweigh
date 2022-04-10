package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Can only be created by Team Leaders
type UnderIssue struct {
	value_objects.BaseModel
	ID                string         `json:"id" gorm:"size:191;not null;unique;primaryKey;"`
	JobItemID         string         `json:"job_item" gorm:"size:191;not null"`
	UnitOfMeasureID   string         `json:"unit_of_measurement_id" gorm:"size:191;not null;"`
	UnitOfMeasure     *UnitOfMeasure `json:"unit_of_measurement"`
	Required          float32        `json:"required"`
	Actual            float32        `json:"actual"`
	CreatedByUsername string         `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User          `json:"created_by"`
	UpdatedByUsername string         `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User          `json:"updated_by"`
}

func (UnderIssue) Tablename() string {
	return "under_issues"
}

func (underIssue *UnderIssue) BeforeCreate(db *gorm.DB) error {
	underIssue.ID = uuid.New().String()
	return nil
}

func (underIssue *UnderIssue) Validate() error {
	errors := map[string]interface{}{}
	if underIssue.JobItemID == "" || len(underIssue.JobItemID) == 0 {
		errors["job"] = "Job Details Required."
	}
	if underIssue.UnitOfMeasureID == "" || len(underIssue.UnitOfMeasureID) == 0 {
		errors["unit_of_measure"] = "Unit of Measure Required."
	}
	if underIssue.CreatedByUsername == "" || len(underIssue.CreatedByUsername) == 0 {
		errors["created_by"] = "Created By Required."
	}
	if underIssue.UpdatedByUsername == "" || len(underIssue.UpdatedByUsername) == 0 {
		errors["updated_by"] = "Updated By Required."
	}
	return utilities.ConvertMapToError(errors)
}
