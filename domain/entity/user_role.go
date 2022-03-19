package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"
)

// Default options are:
// 1. Superuser  				Superuser can create companies only and has no access to any data related to any company.
// 2. Administrator  			Has access to all tables but only to objects which are in the company to which Administrator belongs.
// 					  			Can not create Companies, but can create factories, addresses ect.
// 3. Production Manager  		Next in line to Administrator has access to most tables, but can not create factories, addresses.
// 4. Shift Incharge
// 5. Operator
// In addition administrator can create custom Roles and assign access levels to these roles.

type UserRole struct {
	value_objects.BaseModel
	Role        string   `json:"role" gorm:"size:20;not null;unique;primaryKey;"`
	Description string   `json:"description" gorm:"size:2000;not null;"`
	CompanyID   string   `json:"company_id"`
	Company     *Company `json:"company"`
	Active      bool     `json:"active" gorm:"default:false;"`
}

func (UserRole) Tablename() string {
	return "user_roles"
}

func (userRole *UserRole) Validate() error {
	errors := map[string]interface{}{}
	if userRole.Role == "" || len(userRole.Role) == 0 {
		errors["role"] = "Role Required."
	}
	if userRole.Description == "" || len(userRole.Description) == 0 {
		errors["description"] = "Role Description Required."
	}
	if userRole.CompanyID == "" || len(userRole.CompanyID) == 0 {
		errors["company"] = "Company Required."
	}
	return utilities.ConvertMapToError(errors)
}
