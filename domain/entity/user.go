package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"
)

type User struct {
	value_objects.BaseModel
	Username          string `json:"username" gorm:"size:20;not null;primaryKey;unique;"`
	FirstName         string `json:"first_name" gorm:"size:100;not null;"`
	LastName          string `json:"last_name" gorm:"size:100;"`
	UserRoleRole      string `json:"user_role_role" gorm:"size:20;not null;"`
	CompanyID         int    `json:"company_id"`
	FactoryID         string `json:"factory_id" gorm:"size:191;not null;"`
	Active            bool   `json:"active" gorm:"default:true;"`
	CreatedByUsername string `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User  `json:"created_by"`
	UpdatedByUsername string `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User  `json:"updated_by"`
}

func (User) Tablename() string {
	return "users"
}

func (user *User) Validate() error {
	errors := map[string]interface{}{}
	return utilities.ConvertMapToError(errors)
}
