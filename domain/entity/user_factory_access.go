package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"
)

type UserFactoryAccess struct {
	value_objects.BaseModel
	UserUsername      string   `json:"user_username" gorm:"size:20;not null;"`
	FactoryID         string   `json:"factory_id" gorm:"size:191;not null;"`
	Factory           *Factory `json:"factory"`
	AccessLevel       string   `json:"access_level" gorm:"size:4;default:'0000';"`
	CreatedByUsername string   `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User    `json:"created_by"`
	UpdatedByUsername string   `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User    `json:"updated_by"`
}

func (UserFactoryAccess) Tablename() string {
	return "user_factory_accesses"
}

func (factoryAccess *UserFactoryAccess) Validate() error {
	errors := map[string]interface{}{}
	return utilities.ConvertMapToError(errors)
}
