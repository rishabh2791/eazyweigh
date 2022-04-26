package persistance

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"
	"log"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OverIssueRepo struct {
	DB     *gorm.DB
	Logger hclog.Logger
}

var _ repository.OverIssueRepository = &OverIssueRepo{}

func NewOverIssueRepo(db *gorm.DB, logger hclog.Logger) *OverIssueRepo {
	return &OverIssueRepo{
		DB:     db,
		Logger: logger,
	}
}

func (overIssueRepo *OverIssueRepo) Create(overIssue *entity.OverIssue) (*entity.OverIssue, error) {
	validationErr := overIssue.Validate()
	if validationErr != nil {
		log.Println(validationErr)
		return nil, validationErr
	}

	existingOverIssue := entity.OverIssue{}
	getErr := overIssueRepo.DB.Preload(clause.Associations).Where("job_item_id = ?", overIssue.JobItemID).Take(&existingOverIssue).Error
	if getErr != nil {
		creationErr := overIssueRepo.DB.Create(&overIssue).Error
		if creationErr != nil {
			log.Println(creationErr)
			return nil, creationErr
		}
	} else {
		updationErr := overIssueRepo.DB.Table(entity.OverIssue{}.Tablename()).Where("job_item_id = ?", overIssue.JobItemID).Updates(overIssue).Error
		if updationErr != nil {
			return nil, updationErr
		}
	}

	return overIssue, nil
}

func (overIssueRepo *OverIssueRepo) List(jobID string) ([]entity.OverIssue, error) {
	overIssues := []entity.OverIssue{}
	rawQuery := "SELECT * FROM over_issues WHERE job_item_id IN (SELECT id FROM job_items WHERE job_id='" + jobID + "')"
	getErr := overIssueRepo.DB.
		Preload("UnitOfMeasure.Factory").
		Preload("UnitOfMeasure.Factory.Address").
		Preload("UnitOfMeasure.Factory.CreatedBy").
		Preload("UnitOfMeasure.Factory.CreatedBy.UserRole").
		Preload("UnitOfMeasure.Factory.UpdatedBy").
		Preload("UnitOfMeasure.Factory.UpdatedBy.UserRole").
		Preload("UnitOfMeasure.CreatedBy").
		Preload("UnitOfMeasure.CreatedBy.UserRole").
		Preload("UnitOfMeasure.UpdatedBy").
		Preload("UnitOfMeasure.UpdatedBy.UserRole").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload("JobItem.Material.UnitOfMeasure").
		Preload("JobItem.Material.UnitOfMeasure.Factory").
		Preload("JobItem.Material.UnitOfMeasure.Factory.CreatedBy").
		Preload("JobItem.Material.UnitOfMeasure.Factory.CreatedBy.UserRole").
		Preload("JobItem.Material.UnitOfMeasure.Factory.UpdatedBy").
		Preload("JobItem.Material.UnitOfMeasure.Factory.UpdatedBy.UserRole").
		Preload("JobItem.Material.UnitOfMeasure.CreatedBy").
		Preload("JobItem.Material.UnitOfMeasure.CreatedBy.UserRole").
		Preload("JobItem.Material.UnitOfMeasure.UpdatedBy").
		Preload("JobItem.Material.UnitOfMeasure.UpdatedBy.UserRole").
		Preload("JobItem.Material.UnitOfMeasure.Factory.Address").
		Preload("JobItem.UnitOfMeasure.Factory").
		Preload("JobItem.UnitOfMeasure.Factory.CreatedBy").
		Preload("JobItem.UnitOfMeasure.Factory.CreatedBy.UserRole").
		Preload("JobItem.UnitOfMeasure.Factory.UpdatedBy").
		Preload("JobItem.UnitOfMeasure.Factory.UpdatedBy.UserRole").
		Preload("JobItem.UnitOfMeasure.CreatedBy").
		Preload("JobItem.UnitOfMeasure.CreatedBy.UserRole").
		Preload("JobItem.UnitOfMeasure.UpdatedBy").
		Preload("JobItem.UnitOfMeasure.UpdatedBy.UserRole").
		Preload("JobItem.UnitOfMeasure.Factory.Address").
		Preload("JobItem.CreatedBy.UserRole").
		Preload("JobItem.UpdatedBy.UserRole").
		Preload("JobItem.CreatedBy").
		Preload("JobItem.UpdatedBy").
		Preload("JobItem.Material.CreatedBy").
		Preload("JobItem.Material.CreatedBy.UserRole").
		Preload("JobItem.Material.UpdatedBy").
		Preload("JobItem.Material.UpdatedBy.UserRole").
		Preload(clause.Associations).Raw(rawQuery).Find(&overIssues).Error
	if getErr != nil {
		return nil, getErr
	}

	return overIssues, nil
}

func (overIssueRepo *OverIssueRepo) Update(id string, update *entity.OverIssue) (*entity.OverIssue, error) {
	existingOverIssue := entity.OverIssue{}
	getErr := overIssueRepo.DB.Preload(clause.Associations).Where("id = ?", id).Take(&existingOverIssue).Error
	if getErr != nil {
		return nil, getErr
	}

	updationErr := overIssueRepo.DB.Table(entity.OverIssue{}.Tablename()).Where("id = ?", id).Updates(update).Error
	if updationErr != nil {
		return nil, updationErr
	}

	updated := entity.OverIssue{}
	overIssueRepo.DB.Preload(clause.Associations).Where("id = ?", id).Take(&updated)

	return &updated, nil
}
