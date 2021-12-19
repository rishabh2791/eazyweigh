package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"
)

type Address struct {
	value_objects.BaseModel
	ID                string `json:"id" gorm:"size:191;not null;unique;primaryKey;"`
	Line1             string `json:"line1" gorm:"size:500;not null;"`
	Line2             string `json:"line2" gorm:"size:500;"`
	City              string `json:"city" gorm:"size:100;not null;"`
	State             string `json:"state" gorm:"size:100;not null;"`
	Zip               string `json:"zip" gorm:"size:10;not null;"`
	Country           string `json:"country" gorm:"size:200;not null;"`
	HeadOffice        bool   `json:"head_office" gorm:"default:false;"`
	CreatedByUsername string `json:"created_by_username" gorm:"size:20;not null;"`
	UpdatedByUsername string `json:"updated_by_username" gorm:"size:20;not null;"`
}

func (Address) Tablename() string {
	return "addresses"
}

func (address *Address) Validate() error {
	errors := map[string]interface{}{}
	if address.Line1 == "" || len(address.Line1) == 0 {
		errors["address_line1"] = "Address Line 1 Missing"
	}
	if address.City == "" || len(address.City) == 0 {
		errors["address_city"] = "Address City Missing"
	}
	if address.State == "" || len(address.State) == 0 {
		errors["address_state"] = "Address State Missing"
	}
	if address.Zip == "" || len(address.Zip) == 0 {
		errors["address_zip"] = "Address Zip Code Missing"
	}
	if address.Country == "" || len(address.Country) == 0 {
		errors["address_country"] = "Address Country Missing"
	}
	return utilities.ConvertMapToError(errors)
}
