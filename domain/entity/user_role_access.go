package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"
)

type UserRoleAccess struct {
	value_objects.BaseModel
	UserRoleRole      string `json:"user_role_role" gorm:"size:20;not null;"`
	Table             string `json:"table" gorm:"size:100;not null;"`
	AccessLevel       string `json:"access_level" gorm:"size:4;default:'0000';"`
	CreatedByUsername string `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User  `json:"created_by"`
	UpdatedByUsername string `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User  `json:"updated_by"`
}

func (UserRoleAccess) Tablename() string {
	return "user_role_accesses"
}

func (roleAccess *UserRoleAccess) Validate() error {
	errors := map[string]interface{}{}
	return utilities.ConvertMapToError(errors)
}
