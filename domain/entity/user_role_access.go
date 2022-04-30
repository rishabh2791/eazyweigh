package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"
)

type UserRoleAccess struct {
	value_objects.BaseModel
	UserRoleID        string    `json:"user_role_id" gorm:"size:191;not null;"`
	UserRole          *UserRole `json:"user_role_role"`
	TableName         string    `json:"table_name" gorm:"size:100;not null;"`
	AccessLevel       string    `json:"access_level" gorm:"size:4;default:'0000';"`
	CreatedByUsername string    `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User     `json:"created_by"`
	UpdatedByUsername string    `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User     `json:"updated_by"`
}

func (UserRoleAccess) Tablename() string {
	return "user_role_accesses"
}

func (roleAccess *UserRoleAccess) Validate() error {
	errors := map[string]interface{}{}
	if roleAccess.UserRoleID == "" || len(roleAccess.UserRoleID) == 0 {
		errors["role"] = "User Role Missing"
	}
	if roleAccess.TableName == "" || len(roleAccess.TableName) == 0 {
		errors["table"] = "Table Missing"
	}
	return utilities.ConvertMapToError(errors)
}
