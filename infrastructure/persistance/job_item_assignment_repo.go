package persistance

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type JobItemAssignmentRepo struct {
	DB     *gorm.DB
	Logger hclog.Logger
}

var _ repository.JobItemAssignmentRepository = &JobItemAssignmentRepo{}

func NewJobItemAssignmentRepo(db *gorm.DB, logger hclog.Logger) *JobItemAssignmentRepo {
	return &JobItemAssignmentRepo{
		DB:     db,
		Logger: logger,
	}
}
func (jobItemAssignmentRepo *JobItemAssignmentRepo) Create(jobItemAssignment *entity.JobItemAssignment) (*entity.JobItemAssignment, error) {
	validationErr := jobItemAssignment.Validate()
	if validationErr != nil {
		return nil, validationErr
	}

	creationErr := jobItemAssignmentRepo.DB.Create(jobItemAssignment).Error
	if creationErr != nil {
		return nil, creationErr
	}

	return jobItemAssignment, nil
}

func (jobItemAssignmentRepo *JobItemAssignmentRepo) Get(id string) (*entity.JobItemAssignment, error) {
	jobItemAssignment := entity.JobItemAssignment{}
	getErr := jobItemAssignmentRepo.DB.Preload(clause.Associations).Where("id = ?", id).Take(&jobItemAssignment).Error
	if getErr != nil {
		return nil, getErr
	}

	return &jobItemAssignment, nil
}

func (jobItemAssignmentRepo *JobItemAssignmentRepo) List(conditions string) ([]entity.JobItemAssignment, error) {
	jobItemAssignments := []entity.JobItemAssignment{}
	getErr := jobItemAssignmentRepo.DB.Preload(clause.Associations).Where(conditions).Find(jobItemAssignments).Error
	if getErr != nil {
		return nil, getErr
	}

	return jobItemAssignments, nil
}

func (jobItemAssignmentRepo *JobItemAssignmentRepo) Update(id string, update *entity.JobItemAssignment) (*entity.JobItemAssignment, error) {
	existingjobItemAssignment := entity.JobItemAssignment{}
	getErr := jobItemAssignmentRepo.DB.Preload(clause.Associations).Where("id = ?", id).Take(&existingjobItemAssignment).Error
	if getErr != nil {
		return nil, getErr
	}

	updationErr := jobItemAssignmentRepo.DB.Table(entity.JobItemAssignment{}.Tablename()).Where("id = ?", id).Updates(update).Error
	if updationErr != nil {
		return nil, updationErr
	}

	updated := entity.JobItemAssignment{}
	jobItemAssignmentRepo.DB.Preload(clause.Associations).Take(&updated)

	return &updated, nil
}
