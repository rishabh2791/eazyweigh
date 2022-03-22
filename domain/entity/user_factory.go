package entity

import "eazyweigh/domain/value_objects"

type UserFactory struct {
	value_objects.BaseModel
	FactoryID    string   `json:"factory_id" gorm:"size:191;not null;"`
	Factory      *Factory `json:"factory"`
	UserUsername string   `json:"user_username" gorm:"size:20;not null;"`
	User         *User    `json:"user"`
}

func (UserFactory) Tablename() string {
	return "user_factory"
}
