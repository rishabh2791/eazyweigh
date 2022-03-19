package persistance

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type JobAssignmentRepo struct {
	DB     *gorm.DB
	Logger hclog.Logger
}

var _ repository.JobAssignmentRepository = &JobAssignmentRepo{}

func NewJobAssignmentRepo(db *gorm.DB, logger hclog.Logger) *JobAssignmentRepo {
	return &JobAssignmentRepo{
		DB:     db,
		Logger: logger,
	}
}

func (jobAssignmentRepo *JobAssignmentRepo) Create(jobAssignment *entity.JobAssignment) (*entity.JobAssignment, error) {
	validationErr := jobAssignment.Validate()
	if validationErr != nil {
		return nil, validationErr
	}

	creationErr := jobAssignmentRepo.DB.Create(jobAssignment).Error
	if creationErr != nil {
		return nil, creationErr
	}

	return jobAssignment, nil
}

func (jobAssignmentRepo *JobAssignmentRepo) Get(id string) (*entity.JobAssignment, error) {
	jobAssignment := entity.JobAssignment{}
	getErr := jobAssignmentRepo.DB.Preload(clause.Associations).Where("id = ?", id).Take(&jobAssignment).Error
	if getErr != nil {
		return nil, getErr
	}

	return &jobAssignment, nil
}

func (jobAssignmentRepo *JobAssignmentRepo) List(conditions string) ([]entity.JobAssignment, error) {
	jobAssignments := []entity.JobAssignment{}
	getErr := jobAssignmentRepo.DB.Preload(clause.Associations).Where(conditions).Find(jobAssignments).Error
	if getErr != nil {
		return nil, getErr
	}

	return jobAssignments, nil
}

func (jobAssignmentRepo *JobAssignmentRepo) Update(id string, update *entity.JobAssignment) (*entity.JobAssignment, error) {
	existingJobAssignment := entity.JobAssignment{}
	getErr := jobAssignmentRepo.DB.Preload(clause.Associations).Where("id = ?", id).Take(&existingJobAssignment).Error
	if getErr != nil {
		return nil, getErr
	}

	updationErr := jobAssignmentRepo.DB.Table(entity.JobAssignment{}.Tablename()).Where("id = ?", id).Updates(update).Error
	if updationErr != nil {
		return nil, updationErr
	}

	updated := entity.JobAssignment{}
	jobAssignmentRepo.DB.Preload(clause.Associations).Take(&updated)

	return &updated, nil
}
