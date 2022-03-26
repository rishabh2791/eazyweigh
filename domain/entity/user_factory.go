package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"
)

type UserFactory struct {
	value_objects.BaseModel
	FactoryID    string   `json:"factory_id" gorm:"size:191;not null;uniqueIndex:user_factory;"`
	Factory      *Factory `json:"factory"`
	UserUsername string   `json:"user_username" gorm:"size:20;not null;uniqueIndex:user_factory"`
	User         *User    `json:"user"`
}

func (UserFactory) Tablename() string {
	return "user_factory"
}

func (userFactory *UserFactory) Validate() error {
	errors := map[string]interface{}{}
	if userFactory.FactoryID == "" || len(userFactory.FactoryID) == 0 {
		errors["factory"] = "Factory Details Missing"
	}
	if userFactory.UserUsername == "" || len(userFactory.UserUsername) == 0 {
		errors["user"] = "User Details Missing"
	}
	return utilities.ConvertMapToError(errors)
}
