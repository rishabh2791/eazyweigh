package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Default options are:
// 1. Superuser  				Superuser can create companies only and has no access to any data related to any company.
// 2. Administrator  			Has access to all tables but only to objects which are in the company to which Administrator belongs.
// 					  			Can not create Companies, but can create factories, addresses ect.
// 3. Production Manager  		Next in line to Administrator has access to most tables, but can not create factories, addresses.
// 4. Shift Incharge
// 5. Operator
// 6. Viewer
// 7. Verifier 					Can verify if all the materials weighed out and put on pallet is correct
// In addition administrator can create custom Roles and assign access levels to these roles.

// As soon as company is created, need to assign Administrator, Production Manager, Shift Incharge, Operator and Viewer

type UserRole struct {
	value_objects.BaseModel
	ID          string `json:"id" gorm:"size:191;not null;primaryKey;unique;"`
	Role        string `json:"role" gorm:"size:20;not null;"`
	Description string `json:"description" gorm:"size:2000;not null;"`
	CompanyID   string `json:"company_id" gorm:"size:191;"`
	Active      bool   `json:"active" gorm:"default:false;"`
}

func (UserRole) Tablename() string {
	return "user_roles"
}

func (userRole *UserRole) BeforeCreate(db *gorm.DB) error {
	userRole.ID = uuid.New().String()
	return nil
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
