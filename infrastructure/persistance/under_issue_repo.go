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

	creationErr := underIssueRepo.DB.Create(&underIssue).Error
	if creationErr != nil {
		return nil, creationErr
	}

	return underIssue, nil
}

func (underIssueRepo *UnderIssueRepo) List(jobID string) ([]entity.UnderIssue, error) {
	underIssues := []entity.UnderIssue{}
	getErr := underIssueRepo.DB.Preload(clause.Associations).Where("job_id = ?", jobID).Find(&underIssues).Error
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
