package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"
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

func (factory *Factory) Validate() error {
	errors := map[string]interface{}{}
	return utilities.ConvertMapToError(errors)
}
