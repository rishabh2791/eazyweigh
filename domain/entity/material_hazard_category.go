package entity

import (
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MaterialHazardCategory struct {
	value_objects.BaseModel
	ID                string `json:"id" gorm:"size:191;not null;primaryKey;unique;"`
	HazardCategoryID  string `json:"hazard_category_id" gorm:"size:191;not null;"`
	MaterialID        string `json:"material_id" gorm:"size:191;not null;"`
	CreatedByUsername string `json:"created_by_username" gorm:"size:20;not null;"`
	CreatedBy         *User  `json:"created_by"`
	UpdatedByUsername string `json:"updated_by_username" gorm:"size:20;not null;"`
	UpdatedBy         *User  `json:"updated_by"`
}

func (MaterialHazardCategory) Tablename() string {
	return "material_hazard_categories"
}

func (hazardCategory *MaterialHazardCategory) BeforeCreate(db *gorm.DB) error {
	hazardCategory.ID = uuid.New().String()
	return nil
}

func (hazardCategory *MaterialHazardCategory) Validate() error {
	errors := map[string]interface{}{}
	if hazardCategory.HazardCategoryID == "" || len(hazardCategory.HazardCategoryID) == 0 {
		errors["category"] = "Hazard Category Required."
	}
	if hazardCategory.MaterialID == "" || len(hazardCategory.MaterialID) == 0 {
		errors["material"] = "Material Required."
	}
	if hazardCategory.CreatedByUsername == "" || len(hazardCategory.CreatedByUsername) == 0 {
		errors["created_by"] = "Created By Required."
	}
	if hazardCategory.UpdatedByUsername == "" || len(hazardCategory.UpdatedByUsername) == 0 {
		errors["updated_by"] = "Updated By Required."
	}
	return utilities.ConvertMapToError(errors)
}
