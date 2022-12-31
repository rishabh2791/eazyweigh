package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StepType struct {
	value_objects.BaseModel
	ID                string   `json:"id" gorm:"size:191;not null;primaryKey;unique;"`
	FactoryID         string   `json:"factory_id" gorm:"size:191;not null;uniqueIndex:factory_name;"`
	Factory           *Factory `json:"factory"`
	Name              string   `json:"description" gorm:"size:300;uniqueIndex:factory_name;"`
	Title             string   `json:"title" gorm:"size:50;not null;"`
	Body              string   `json:"body" gorm:"size:50;not null;"`
	Footer            string   `json:"footer" gorm:"size:50;not null;"`
	CreatedByUsername string   `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User    `json:"created_by"`
	UpdatedByUsername string   `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User    `json:"updated_by"`
}

func (StepType) Tablename() string {
	return "step_types"
}

func (stepType *StepType) BeforeCreate(db *gorm.DB) error {
	stepType.ID = uuid.New().String()
	return nil
}

func (stepType *StepType) Validate() error {
	errors := map[string]interface{}{}
	if stepType.Name == "" || len(stepType.Name) == 0 {
		errors["name"] = "Step Name Required."
	}
	if stepType.FactoryID == "" || len(stepType.FactoryID) == 0 {
		errors["factory"] = "Factory Required."
	}
	if stepType.Title == "" || len(stepType.Title) == 0 {
		errors["title"] = "Title Required."
	}
	if stepType.Body == "" || len(stepType.Body) == 0 {
		errors["body"] = "Body Required."
	}
	if stepType.Footer == "" || len(stepType.Footer) == 0 {
		errors["footer"] = "Footer Required."
	}
	if stepType.CreatedByUsername == "" || len(stepType.CreatedByUsername) == 0 {
		errors["created_by"] = "Created By Required."
	}
	if stepType.UpdatedByUsername == "" || len(stepType.UpdatedByUsername) == 0 {
		errors["updated_by"] = "Updated By Required."
	}
	return utilities.ConvertMapToError(errors)
}
