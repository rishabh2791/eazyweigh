package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"
)

type UserCompany struct {
	value_objects.BaseModel
	CompanyID    string   `json:"company_id" gorm:"size:191;not null;uniqueIndex:user_company;"`
	Company      *Company `json:"company"`
	UserUsername string   `json:"user_username" gorm:"size:20;not null;uniqueIndex:user_company;"`
	User         *User    `json:"user"`
}

func (UserCompany) Tablename() string {
	return "user_company"
}

func (userCompany *UserCompany) Validate() error {
	errors := map[string]interface{}{}
	if userCompany.CompanyID == "" || len(userCompany.CompanyID) == 0 {
		errors["company"] = "Company Details Missing"
	}
	if userCompany.UserUsername == "" || len(userCompany.UserUsername) == 0 {
		errors["user"] = "User Details Missing"
	}
	return utilities.ConvertMapToError(errors)
}
