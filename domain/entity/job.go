package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"
)

type Job struct {
	value_objects.BaseModel
	ID                string         `json:"id" gorm:"size:191;not null;unique;primaryKey;"`
	JobCode           string         `json:"job_code" gorm:"size:10;not null;"`
	MaterialID        string         `json:"material_code" gorm:"size:191;not null;"`
	Material          *Material      `json:"material"`
	Quantity          float32        `json:"quantity"`
	UnitOfMeasureID   string         `json:"unit_of_measurement_id" gorm:"size:191;not null;"`
	UnitOfMeasure     *UnitOfMeasure `json:"unit_of_measurement"`
	JobItems          []JobItem      `json:"job_items"`
	CreatedByUsername string         `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User          `json:"created_by"`
	UpdatedByUsername string         `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User          `json:"updated_by"`
}

func (Job) Tablename() string {
	return "jobs"
}

func (job *Job) Validate() error {
	errors := map[string]interface{}{}
	return utilities.ConvertMapToError(errors)
}

func (job *Job) IsIncomplete() bool {
	for _, jobItem := range job.JobItems {
		if !jobItem.Complete {
			return true
		}
	}
	return false
}
