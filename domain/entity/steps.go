package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Step struct {
	value_objects.BaseModel
	ID                string    `json:"id" gorm:"size:191;not null;unique;primaryKey;"`
	ProcessID         string    `json:"process_id"`
	StepTypeID        string    `json:"step_type_id" gorm:"size:191;not null;"`
	StepType          *StepType `json:"step_type"`
	Description       string    `json:"description" gorm:"size:2000;not null;"`
	MaterialID        string    `json:"material_id" gorm:"size:191;"`
	Value             float32   `json:"value" gorm:"default:0;"`
	Sequence          int       `json:"sequence"`
	CreatedByUsername string    `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User     `json:"created_by"`
	UpdatedByUsername string    `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User     `json:"updated_by"`
}

func (Step) Tablename() string {
	return "steps"
}

func (step *Step) BeforeCreate(db *gorm.DB) error {
	step.ID = uuid.New().String()
	return nil
}

func (step *Step) Validate() error {
	errors := map[string]interface{}{}
	if step.StepTypeID == "" || len(step.StepTypeID) == 0 {
		errors["step_type"] = "Step Type Required."
	}
	if step.Description == "" || len(step.Description) == 0 {
		errors["step_type"] = "Step Full Description Required."
	}
	if step.CreatedByUsername == "" || len(step.CreatedByUsername) == 0 {
		errors["created_by"] = "Created By Required."
	}
	if step.UpdatedByUsername == "" || len(step.UpdatedByUsername) == 0 {
		errors["updated_by"] = "Updated By Required."
	}
	return utilities.ConvertMapToError(errors)
}
