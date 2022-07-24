package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Process struct {
	value_objects.BaseModel
	ID                string    `json:"id" gorm:"size:191;not null;unique;primaryKey;"`
	MaterialID        string    `json:"material_id" gorm:"size:191;not null;"`
	Material          *Material `json:"material"`
	Steps             []Step    `json:"steps"`
	CreatedByUsername string    `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User     `json:"created_by"`
	UpdatedByUsername string    `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User     `json:"updated_by"`
}

func (Process) Tablename() string {
	return "steps"
}

func (process *Process) BeforeCreate(db *gorm.DB) error {
	process.ID = uuid.New().String()
	return nil
}

func (process *Process) Validate() error {
	errors := map[string]interface{}{}
	if process.MaterialID == "" || len(process.MaterialID) == 0 {
		errors["material"] = "Material Required."
	}
	if len(process.Steps) == 0 {
		errors["step_type"] = "Process Steps Required."
	}
	if process.CreatedByUsername == "" || len(process.CreatedByUsername) == 0 {
		errors["created_by"] = "Created By Required."
	}
	if process.UpdatedByUsername == "" || len(process.UpdatedByUsername) == 0 {
		errors["updated_by"] = "Updated By Required."
	}
	return utilities.ConvertMapToError(errors)
}
