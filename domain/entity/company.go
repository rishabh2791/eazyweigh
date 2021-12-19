package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"
)

type Company struct {
	value_objects.BaseModel
	ID                int       `json:"id" gorm:"primaryKey;autoIncrement;"`
	Name              string    `json:"name" gorm:"size:2000;not null;"`
	Users             []User    `json:"company_users"`
	Factories         []Factory `json:"company_factories"`
	Active            bool      `json:"active" gorm:"default:true;"`
	CreatedByUsername string    `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User     `json:"created_by"`
	UpdatedByUsername string    `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User     `json:"updated_by"`
}

func (Company) Tablename() string {
	return "companies"
}

func (company *Company) Validate() error {
	errors := map[string]interface{}{}
	return utilities.ConvertMapToError(errors)
}
