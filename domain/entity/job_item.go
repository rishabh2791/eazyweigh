package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

//sa123
type JobItem struct {
	value_objects.BaseModel
	ID                string         `json:"id" gorm:"size:191;not null;unique;primaryKey;"`
	JobID             string         `json:"job_id" gorm:"size:191;not null;uniqueIndex:job_material;"`
	MaterialID        string         `json:"material_code" gorm:"size:191;not null;uniqueIndex:job_material;"`
	Material          *Material      `json:"material"`
	UnitOfMeasureID   string         `json:"unit_of_measurement_id" gorm:"size:191;not null;"`
	UnitOfMeasure     *UnitOfMeasure `json:"unit_of_measurement"`
	RequiredWeight    float32        `json:"required_weight"`
	UpperBound        float32        `json:"upper_bound"`
	LowerBound        float32        `json:"lower_bound"`
	OverIssue         bool           `json:"over_issue" gorm:"default:false;"`
	UnderIssue        bool           `json:"under_issue" gorm:"default:false;"`
	ActualWeight      float32        `json:"actual_weight"`
	Complete          bool           `json:"complete" gorm:"default:false;"`
	TimeTaken         int            `json:"time_taken" gorm:"default:1;"`
	CreatedByUsername string         `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User          `json:"created_by"`
	UpdatedByUsername string         `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User          `json:"updated_by"`
}

func (JobItem) Tablename() string {
	return "job_items"
}

func (jobItem *JobItem) BeforeCreate(db *gorm.DB) error {
	jobItem.ID = uuid.New().String()
	return nil
}

func (jobItem *JobItem) Validate() error {
	errors := map[string]interface{}{}
	if jobItem.JobID == "" || len(jobItem.JobID) == 0 {
		errors["job"] = "Job Required."
	}
	if jobItem.MaterialID == "" || len(jobItem.MaterialID) == 0 {
		errors["material"] = "Material Required."
	}
	if jobItem.UnitOfMeasureID == "" || len(jobItem.UnitOfMeasureID) == 0 {
		errors["uom"] = "Unit of Measurement Required."
	}
	if jobItem.CreatedByUsername == "" || len(jobItem.CreatedByUsername) == 0 {
		errors["created_by"] = "Created By Required."
	}
	if jobItem.UpdatedByUsername == "" || len(jobItem.UpdatedByUsername) == 0 {
		errors["updated_by"] = "Updated By Required."
	}
	return utilities.ConvertMapToError(errors)
}
