package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"
)

type OverIssue struct {
	value_objects.BaseModel
	ID                 string         `json:"id" gorm:"size:191;not null;unique;primaryKey;"`
	JobItemID          string         `json:"job_item" gorm:"size:191;not null"`
	SupervisorUsername string         `json:"supervisor_username" gorm:"size:20;not null;"`
	Supervisor         *User          `json:"supervisor"`
	UnitOfMeasureID    string         `json:"unit_of_measurement_id" gorm:"size:191;not null;"`
	UnitOfMeasure      *UnitOfMeasure `json:"unit_of_measurement"`
	Required           float32        `json:"required"`
	Actual             float32        `json:"actual"`
	Approved           bool           `json:"approved" gorm:"default:false;"`
	CreatedByUsername  string         `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy          *User          `json:"created_by"`
	UpdatedByUsername  string         `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy          *User          `json:"updated_by"`
}

func (OverIssue) Tablename() string {
	return "over_issues"
}

func (overIssue *OverIssue) Validate() error {
	errors := map[string]interface{}{}
	return utilities.ConvertMapToError(errors)
}
