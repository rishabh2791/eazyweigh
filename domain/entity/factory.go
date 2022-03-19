package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"
	"strconv"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Factory struct {
	value_objects.BaseModel
	ID                string     `json:"id" gorm:"size:191;not null;primaryKey;unique;"`
	CompanyID         int        `json:"company_id"`
	AddressID         string     `json:"address_id" gorm:"size:191;not null;"`
	Address           *Address   `json:"address"`
	Users             []User     `json:"factory_users"`
	Terminals         []Terminal `json:"factory_terminals"`
	Materials         []Material `json:"factory_materials"`
	CreatedByUsername string     `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User      `json:"created_by"`
	UpdatedByUsername string     `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User      `json:"updated_by"`
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
	if strconv.Itoa(factory.CompanyID) == "" || len(strconv.Itoa(factory.CompanyID)) == 0 {
		errors["company_id"] = "Company ID Required."
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
