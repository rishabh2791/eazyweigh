package persistance

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UnderIssueRepo struct {
	DB     *gorm.DB
	Logger hclog.Logger
}

var _ repository.UnderIssueRepository = &UnderIssueRepo{}

func NewUnderIssueRepo(db *gorm.DB, logger hclog.Logger) *UnderIssueRepo {
	return &UnderIssueRepo{
		DB:     db,
		Logger: logger,
	}
}

func (underIssueRepo *UnderIssueRepo) Create(underIssue *entity.UnderIssue) (*entity.UnderIssue, error) {
	validationErr := underIssue.Validate()
	if validationErr != nil {
		return nil, validationErr
	}

	existingUnderIsse := entity.UnderIssue{}
	getErr := underIssueRepo.DB.Preload(clause.Associations).Where("job_item_id = ?", underIssue.JobItemID).Take(&existingUnderIsse).Error
	if getErr != nil {
		creationErr := underIssueRepo.DB.Create(&underIssue).Error
		if creationErr != nil {
			return nil, creationErr
		}
	} else {
		updationErr := underIssueRepo.DB.Table(entity.UnderIssue{}.Tablename()).Where("job_item_id = ?", underIssue.JobItemID).Updates(underIssue).Error
		if updationErr != nil {
			return nil, updationErr
		}
	}

	return underIssue, nil
}

func (underIssueRepo *UnderIssueRepo) List(conditions string) ([]entity.UnderIssue, error) {
	underIssues := []entity.UnderIssue{}
	rawQuery := "SELECT * FROM under_issues WHERE job_item_id IN (SELECT id FROM job_items WHERE " + conditions + ")"

	getErr := underIssueRepo.DB.
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
		Preload("JobItem.Material").
		Preload("JobItem.Material.UnitOfMeasure").
		Preload("JobItem.Material.UnitOfMeasure.Factory").
		Preload("JobItem.Material.UnitOfMeasure.Factory.Address").
		Preload("JobItem.Material.UnitOfMeasure.Factory.CreatedBy").
		Preload("JobItem.Material.UnitOfMeasure.Factory.CreatedBy.UserRole").
		Preload("JobItem.Material.UnitOfMeasure.Factory.UpdatedBy").
		Preload("JobItem.Material.UnitOfMeasure.Factory.UpdatedBy.UserRole").
		Preload("JobItem.Material.UnitOfMeasure.CreatedBy").
		Preload("JobItem.Material.UnitOfMeasure.CreatedBy.UserRole").
		Preload("JobItem.Material.UnitOfMeasure.UpdatedBy").
		Preload("JobItem.Material.UnitOfMeasure.UpdatedBy.UserRole").
		Preload("JobItem.Material.CreatedBy").
		Preload("JobItem.Material.CreatedBy.UserRole").
		Preload("JobItem.Material.UpdatedBy").
		Preload("JobItem.Material.UpdatedBy.UserRole").
		Preload("JobItem.UnitOfMeasure").
		Preload("JobItem.UnitOfMeasure.Factory").
		Preload("JobItem.UnitOfMeasure.Factory.Address").
		Preload("JobItem.UnitOfMeasure.Factory.CreatedBy").
		Preload("JobItem.UnitOfMeasure.Factory.CreatedBy.UserRole").
		Preload("JobItem.UnitOfMeasure.Factory.UpdatedBy").
		Preload("JobItem.UnitOfMeasure.Factory.UpdatedBy.UserRole").
		Preload("JobItem.UnitOfMeasure.CreatedBy").
		Preload("JobItem.UnitOfMeasure.CreatedBy.UserRole").
		Preload("JobItem.UnitOfMeasure.UpdatedBy").
		Preload("JobItem.UnitOfMeasure.UpdatedBy.UserRole").
		Preload("JobItem.CreatedBy").
		Preload("JobItem.CreatedBy.UserRole").
		Preload("JobItem.UpdatedBy").
		Preload("JobItem.UpdatedBy.UserRole").
		Preload("JobItem.JobItemWeighing").
		Preload("JobItem.JobItemWeighing.CreatedBy").
		Preload("JobItem.JobItemWeighing.CreatedBy.UserRole").
		Preload("JobItem.JobItemWeighing.UpdatedBy").
		Preload("JobItem.JobItemWeighing.UpdatedBy.UserRole").
		Preload(clause.Associations).Raw(rawQuery).Find(&underIssues).Error
	if getErr != nil {
		return nil, getErr
	}

	return underIssues, nil
}

func (underIssueRepo *UnderIssueRepo) Update(id string, update *entity.UnderIssue) (*entity.UnderIssue, error) {
	existingunderIssue := entity.UnderIssue{}
	getErr := underIssueRepo.DB.Preload(clause.Associations).Where("id = ?", id).Take(&existingunderIssue).Error
	if getErr != nil {
		return nil, getErr
	}

	updationErr := underIssueRepo.DB.Table(entity.UnderIssue{}.Tablename()).Where("id = ?", id).Updates(update).Error
	if updationErr != nil {
		return nil, updationErr
	}

	updated := entity.UnderIssue{}
	underIssueRepo.DB.Preload(clause.Associations).Where("id = ?", id).Take(&updated)

	return &updated, nil
}
