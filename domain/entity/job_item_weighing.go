package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type JobItemWeighing struct {
	value_objects.BaseModel
	ID                string    `json:"id" gorm:"size:191;not null;primaryKey;unique;"`
	JOBItemID         string    `json:"job_item_id" gorm:"size:191;not null;"`
	JobItem           *JobItem  `json:"job_item"`
	Weight            float32   `json:"weight" `
	Batch             string    `json:"batch" gorm:"size:50;"`
	Verified          bool      `json:"verified" gorm:"default:false;"`
	StartTime         time.Time `json:"start_time"`
	EndTime           time.Time `json:"end_time"`
	CreatedByUsername string    `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User     `json:"created_by"`
	UpdatedByUsername string    `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User     `json:"updated_by"`
}

func (JobItemWeighing) Tablename() string {
	return "job_item_weighings"
}

func (jobItemWeighing *JobItemWeighing) BeforeCreate(db *gorm.DB) error {
	jobItemWeighing.ID = uuid.New().String()
	return nil
}

func (jobItemWeighing *JobItemWeighing) Validate() error {
	errors := map[string]interface{}{}
	if jobItemWeighing.JOBItemID == "" || len(jobItemWeighing.JOBItemID) == 0 {
		errors["job_item"] = "Job Item Required."
	}
	if jobItemWeighing.Batch == "" || len(jobItemWeighing.Batch) == 0 {
		errors["batch"] = "Batch Required."
	}
	if jobItemWeighing.CreatedByUsername == "" || len(jobItemWeighing.CreatedByUsername) == 0 {
		errors["created_by"] = "Created By Required."
	}
	if jobItemWeighing.UpdatedByUsername == "" || len(jobItemWeighing.UpdatedByUsername) == 0 {
		errors["updated_by"] = "Updated By Required."
	}
	return utilities.ConvertMapToError(errors)
}

type WeighingBatch struct {
	JobItemWeighing
	JobItemID         string  `json:"job_item_id"`
	JobItemMaterialID string  `json:"job_item_material_id"`
	RequiredWeight    float32 `json:"required_weight"`
	ActualWeight      float32 `json:"actual_weight"`
	JobCode           string  `json:"job_code"`
	JobMaterialID     string  `json:"job_material_id"`
}
