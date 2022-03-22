package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Company struct {
	value_objects.BaseModel
	ID                string `json:"id" gorm:"primaryKey;size:191;not null;unique;"`
	Name              string `json:"name" gorm:"size:2000;not null;"`
	Active            bool   `json:"active" gorm:"default:true;"`
	CreatedByUsername string `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User  `json:"created_by"`
	UpdatedByUsername string `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User  `json:"updated_by"`
}

func (Company) Tablename() string {
	return "companies"
}

func (company *Company) BeforeCreate(db *gorm.DB) error {
	company.ID = uuid.New().String()
	return nil
}

func (company *Company) Validate() error {
	errors := map[string]interface{}{}
	if company.Name == "" || len(company.Name) == 0 {
		errors["name"] = "Company Name Required."
	}
	if company.CreatedByUsername == "" || len(company.CreatedByUsername) == 0 {
		errors["created_by"] = "Created By Required."
	}
	if company.UpdatedByUsername == "" || len(company.UpdatedByUsername) == 0 {
		errors["updated_by"] = "Updated By Required."
	}
	return utilities.ConvertMapToError(errors)
}
