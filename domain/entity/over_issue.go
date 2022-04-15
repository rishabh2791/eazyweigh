package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Can only be created by Team Leaders
type OverIssue struct {
	value_objects.BaseModel
	ID                string         `json:"id" gorm:"size:191;not null;unique;primaryKey;"`
	JobItemID         string         `json:"job_item" gorm:"size:191;not null"`
	UnitOfMeasureID   string         `json:"unit_of_measurement_id" gorm:"size:191;not null;"`
	UnitOfMeasure     *UnitOfMeasure `json:"unit_of_measurement"`
	Required          float32        `json:"required"`
	Actual            float32        `json:"actual"`
	Weight            float32        `json:"weight" gorm:"default:0;"`
	Verified          bool           `json:"verified" gorm:"default:false;"`
	Weighed           bool           `json:"weighed" gorm:"default:false"`
	CreatedByUsername string         `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User          `json:"created_by"`
	UpdatedByUsername string         `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User          `json:"updated_by"`
}

func (OverIssue) Tablename() string {
	return "over_issues"
}

func (overIssue *OverIssue) BeforeCreate(db *gorm.DB) error {
	overIssue.ID = uuid.New().String()
	return nil
}

func (overIssue *OverIssue) Validate() error {
	errors := map[string]interface{}{}
	if overIssue.JobItemID == "" || len(overIssue.JobItemID) == 0 {
		errors["job"] = "Job Details Required."
	}
	if overIssue.UnitOfMeasureID == "" || len(overIssue.UnitOfMeasureID) == 0 {
		errors["unit_of_measure"] = "Unit of Measure Required."
	}
	if overIssue.CreatedByUsername == "" || len(overIssue.CreatedByUsername) == 0 {
		errors["created_by"] = "Created By Required."
	}
	if overIssue.UpdatedByUsername == "" || len(overIssue.UpdatedByUsername) == 0 {
		errors["updated_by"] = "Updated By Required."
	}
	return utilities.ConvertMapToError(errors)
}
