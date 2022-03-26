package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Factory struct {
	value_objects.BaseModel
	ID                string   `json:"id" gorm:"size:191;not null;primaryKey;unique;"`
	Name              string   `json:"name" gorm:"size:100;not null;uniqueIndex:company_factory;"`
	CompanyID         string   `json:"company_id" gorm:"size:191;uniqueIndex:company_factory;"`
	AddressID         string   `json:"address_id" gorm:"size:191;not null;"`
	Address           *Address `json:"address"`
	CreatedByUsername string   `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User    `json:"created_by"`
	UpdatedByUsername string   `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User    `json:"updated_by"`
}

func (Factory) Tablename() string {
	return "factories"
}

func (factory *Factory) BeforeCreate(db *gorm.DB) error {
	factory.ID = uuid.New().String()
	return nil
}

func (factory *Factory) Validate() error {
	errors := map[string]interface{}{}
	if factory.Name == "" || len(factory.Name) == 0 {
		errors["name"] = "Name Required."
	}
	if factory.CompanyID == "" || len(factory.CompanyID) == 0 {
		errors["company"] = "Company ID Required."
	}
	if factory.AddressID == "" || len(factory.AddressID) == 0 {
		errors["factory"] = "Factory ID Required."
	}
	if factory.CreatedByUsername == "" || len(factory.CreatedByUsername) == 0 {
		errors["created_by"] = "Created By Required."
	}
	if factory.UpdatedByUsername == "" || len(factory.UpdatedByUsername) == 0 {
		errors["updated_by"] = "Updated By Required."
	}
	return utilities.ConvertMapToError(errors)
}
