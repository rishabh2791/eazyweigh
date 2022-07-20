package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"
)

type UserCompanyAccess struct {
	value_objects.BaseModel
	UserUsername      string   `json:"user_username" gorm:"size:20;not null;"`
	CompanyID         int      `json:"company_id" gorm:"not null;size:191;"`
	Company           *Company `json:"company"`
	AccessLevel       string   `json:"access_level" gorm:"size:4;default:'0000';"`
	CreatedByUsername string   `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User    `json:"created_by"`
	UpdatedByUsername string   `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User    `json:"updated_by"`
}

func (UserCompanyAccess) Tablename() string {
	return "user_company_accesses"
}

func (companyAccess *UserCompanyAccess) Validate() error {
	errors := map[string]interface{}{}
	return utilities.ConvertMapToError(errors)
}
