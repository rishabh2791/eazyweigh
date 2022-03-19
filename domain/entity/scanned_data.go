package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ScannedData struct {
	value_objects.BaseModel
	ID           string `json:"id" gorm:"size:191;not null;primaryKey;unique;"`
	ActualCode   string `json:"actual_code" gorm:"size:1000;not null;"`
	ExpectedCode string `json:"expected_code" gorm:"size:1000;not null;"`
	UserUsername string `json:"user_username" gorm:"size:20;not null;"`
	User         *User  `json:"user"`
	JobID        string `json:"job_id" gorm:"size:191;not null;"`
	Job          *Job   `json:"job"`
}

func (ScannedData) Tablename() string {
	return "scanned_data"
}

func (scannedData *ScannedData) BeforeCreate(db *gorm.DB) error {
	scannedData.ID = uuid.New().String()
	return nil
}

func (scannedData *ScannedData) Validate() error {
	errors := map[string]interface{}{}
	if scannedData.ActualCode == "" || len(scannedData.ActualCode) == 0 {
		errors["actual_code"] = "Actual Scanned Code Required."
	}
	if scannedData.ExpectedCode == "" || len(scannedData.ExpectedCode) == 0 {
		errors["expected_code"] = "Expected Scanned Code Required."
	}
	if scannedData.UserUsername == "" || len(scannedData.UserUsername) == 0 {
		errors["user"] = "User Required."
	}
	if scannedData.JobID == "" || len(scannedData.JobID) == 0 {
		errors["job"] = "Job Details Required."
	}
	return utilities.ConvertMapToError(errors)
}
