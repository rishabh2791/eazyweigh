package entity

import "eazyweigh/domain/value_objects"

type UserCompany struct {
	value_objects.BaseModel
	CompanyID    string   `json:"company_id" gorm:"size:191;not null;"`
	Company      *Company `json:"company"`
	UserUsername string   `json:"user_username" gorm:"size:20;not null;"`
	User         *User    `json:"user"`
}

func (UserCompany) Tablename() string {
	return "user_company"
}
