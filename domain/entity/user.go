package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"
	"eazyweigh/infrastructure/utilities/security"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	value_objects.BaseModel
	Username          string `json:"username" gorm:"size:20;not null;primaryKey;unique;"`
	Password          string `json:"password,omitempty" gorm:"size:200;not null;"`
	FirstName         string `json:"first_name" gorm:"size:100;not null;"`
	LastName          string `json:"last_name" gorm:"size:100;"`
	Email             string `json:"email" gorm:"size:100;not null;unique;"`
	UserRoleRole      string `json:"user_role_role" gorm:"size:20;not null;"`
	SecretKey         string `json:"-" gorm:"size:191;not null;unique;"`
	Superuser         bool   `json:"superuser" gorm:"default:false;"`
	CompanyID         int    `json:"company_id"`
	FactoryID         string `json:"factory_id" gorm:"size:191;not null;"`
	Active            bool   `json:"active" gorm:"default:true;"`
	CreatedByUsername string `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User  `json:"created_by"`
	UpdatedByUsername string `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User  `json:"updated_by"`
}

func (User) Tablename() string {
	return "users"
}

func (user *User) BeforeCreate(db *gorm.DB) error {
	user.SecretKey = strings.ReplaceAll(uuid.NewString(), "-", "")
	if user.Superuser && user.CreatedByUsername == "" {
		user.CreatedByUsername = user.Username
	}
	hashedPassword, passError := security.Hash(user.Password)
	if passError != nil {
		return passError
	}
	user.Password = string(hashedPassword)
	return nil
}

func (user *User) Validate(action string) error {
	errors := map[string]interface{}{}
	switch action {
	case "login":
		if user.Username == "" || len(user.Username) == 0 {
			errors["username"] = "Username Missing"
		}
		if user.Password == "" || len(user.Password) == 0 {
			errors["password"] = "Password Missing"
		}
	case "superuser":
		if user.Username == "" || len(user.Username) == 0 {
			errors["username"] = "Username Missing"
		}
		if user.FirstName == "" || len(user.FirstName) == 0 {
			errors["first_name"] = "First Name Missing"
		}
		if user.Email == "" || len(user.Email) == 0 {
			errors["email"] = "EMail Missing"
		}
		if user.Password == "" || len(user.Password) == 0 {
			errors["password"] = "Password Missing"
		}
	case "register":
		if user.Username == "" || len(user.Username) == 0 {
			errors["username"] = "Username Missing"
		}
		if user.FirstName == "" || len(user.FirstName) == 0 {
			errors["first_name"] = "First Name Missing"
		}
		if user.Email == "" || len(user.Email) == 0 {
			errors["email"] = "EMail Missing"
		}
		if user.CreatedByUsername == "" || len(user.CreatedByUsername) == 0 {
			errors["created_by"] = "Created By Missing"
		}
		if user.UpdatedByUsername == "" || len(user.UpdatedByUsername) == 0 {
			errors["updated_by"] = "Updated By Missing"
		}
		if user.Password == "" || len(user.Password) == 0 {
			errors["password"] = "Password Missing"
		}
	case "reset_password":
		if user.Username == "" || len(user.Username) == 0 {
			errors["username"] = "Username Missing"
		}
		if user.Password == "" || len(user.Password) == 0 {
			errors["password"] = "Password Missing"
		}
	default:
		if user.Username == "" || len(user.Username) == 0 {
			errors["username"] = "Username Missing"
		}
		if user.FirstName == "" || len(user.FirstName) == 0 {
			errors["first_name"] = "First Name Missing"
		}
		if user.Email == "" || len(user.Email) == 0 {
			errors["email"] = "EMail Missing"
		}
		if user.CreatedByUsername == "" || len(user.CreatedByUsername) == 0 {
			errors["created_by"] = "Created By Missing"
		}
		if user.UpdatedByUsername == "" || len(user.UpdatedByUsername) == 0 {
			errors["updated_by"] = "Updated By Missing"
		}
		if user.Password == "" || len(user.Password) == 0 {
			errors["password"] = "Password Missing"
		}
	}
	return utilities.ConvertMapToError(errors)
}
