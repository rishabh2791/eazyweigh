package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type HazardCategory struct {
	value_objects.BaseModel
	ID                string `json:"id" gorm:"size:191;not null;primaryKey;unique;"`
	CategoryName      string `json:"category_name" gorm:"size:100;"`
	DisplayText       string `json:"display_text" gorm:"size:1000;not null;"`
	DisplayImageURL   string `json:"display_image_url" gorm:"size:1000;"`
	CreatedByUsername string `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User  `json:"created_by"`
	UpdatedByUsername string `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User  `json:"updated_by"`
}

func (HazardCategory) Tablename() string {
	return "hazard_categories"
}

func (hazardCategory *HazardCategory) BeforeCreate(db *gorm.DB) error {
	hazardCategory.ID = uuid.New().String()
	return nil
}

func (hazardCategory *HazardCategory) Validate() error {
	errors := map[string]interface{}{}
	if hazardCategory.CategoryName == "" || len(hazardCategory.CategoryName) == 0 {
		errors["category"] = "Category Name Required."
	}
	if hazardCategory.DisplayText == "" || len(hazardCategory.DisplayText) == 0 {
		errors["display_text"] = "Display Text Required."
	}
	if hazardCategory.CreatedByUsername == "" || len(hazardCategory.CreatedByUsername) == 0 {
		errors["created_by"] = "Created By Required."
	}
	if hazardCategory.UpdatedByUsername == "" || len(hazardCategory.UpdatedByUsername) == 0 {
		errors["updated_by"] = "Updated By Required."
	}
	return utilities.ConvertMapToError(errors)
}
