package persistance

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"

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
		return nil, validationErr
	}

	creationErr := overIssueRepo.DB.Create(&overIssue).Error
	if creationErr != nil {
		return nil, creationErr
	}

	return overIssue, nil
}

func (overIssueRepo *OverIssueRepo) List(jobID string) ([]entity.OverIssue, error) {
	overIssues := []entity.OverIssue{}
	getErr := overIssueRepo.DB.Preload(clause.Associations).Where("job_id = ?", jobID).Find(&overIssues).Error
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
