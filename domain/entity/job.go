package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Job struct {
	value_objects.BaseModel
	ID                string         `json:"id" gorm:"size:191;not null;unique;primaryKey;"`
	JobCode           string         `json:"job_code" gorm:"size:10;not null;uniqueIndex:factory_job;"`
	FactoryID         string         `json:"factory_id" gorm:"size:191;not null;uniqueIndex:factory_job;"`
	Factory           *Factory       `json:"factory"`
	MaterialID        string         `json:"material_id" gorm:"size:191;not null;"`
	Material          *Material      `json:"material"`
	Quantity          float32        `json:"quantity" gorm:"default:0.0;"`
	UnitOfMeasureID   string         `json:"unit_of_measurement_id" gorm:"size:191;not null;"`
	UnitOfMeasure     *UnitOfMeasure `json:"unit_of_measurement"`
	JobItems          []JobItem      `json:"job_items"`
	Processing        bool           `json:"processing" gorm:"default:false;"`
	Complete          bool           `json:"complete" gorm:"default:false;"`
	CreatedByUsername string         `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User          `json:"created_by"`
	UpdatedByUsername string         `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User          `json:"updated_by"`
}

func (Job) Tablename() string {
	return "jobs"
}

func (job *Job) BeforeCreate(db *gorm.DB) error {
	if job.ID == "" || len(job.ID) == 0 {
		job.ID = uuid.New().String()
	}
	return nil
}

func (job *Job) Validate() error {
	errors := map[string]interface{}{}
	if job.JobCode == "" || len(job.JobCode) == 0 {
		errors["job_code"] = "Job Code Required."
	}
	if job.MaterialID == "" || len(job.MaterialID) == 0 {
		errors["material"] = "Material Required."
	}
	if job.UnitOfMeasureID == "" || len(job.UnitOfMeasureID) == 0 {
		errors["uom"] = "Unit of Measurement for Material: " + job.MaterialID + " Required."
	}
	if job.CreatedByUsername == "" || len(job.CreatedByUsername) == 0 {
		errors["created_by"] = "Created By Required."
	}
	if job.UpdatedByUsername == "" || len(job.UpdatedByUsername) == 0 {
		errors["updated_by"] = "Updated By Required."
	}
	return utilities.ConvertMapToError(errors)
}

func (job *Job) IsComplete() bool {
	var complete bool
	complete = true
	for _, jobItem := range job.JobItems {
		if jobItem.Material.IsWeighed {
			if jobItem.IsComplete() {
				complete = complete && true
			} else {
				complete = complete && false
			}
		}
	}
	return complete
}
