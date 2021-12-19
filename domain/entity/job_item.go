package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"
)

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
	CreatedByUsername string         `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User          `json:"created_by"`
	UpdatedByUsername string         `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User          `json:"updated_by"`
}

func (JobItem) Tablename() string {
	return "job_items"
}

func (jobItem *JobItem) Validate() error {
	errors := map[string]interface{}{}
	return utilities.ConvertMapToError(errors)
}
