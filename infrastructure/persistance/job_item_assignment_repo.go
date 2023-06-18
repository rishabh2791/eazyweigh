package persistance

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
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

	existingAssignment := entity.JobItemAssignment{}
	getErr := jobItemAssignmentRepo.DB.Where("job_item_id = ?", jobItemAssignment.JobItemID).Take(&existingAssignment).Error
	if getErr != nil {
		creationErr := jobItemAssignmentRepo.DB.Create(jobItemAssignment).Error
		if creationErr != nil {
			return nil, creationErr
		}
	} else {
		updationErr := jobItemAssignmentRepo.DB.Table(entity.JobItemAssignment{}.Tablename()).Where("job_item_id= ?", jobItemAssignment.JobItemID).Updates(jobItemAssignment).Error
		if updationErr != nil {
			return nil, updationErr
		}
	}

	updateErr := jobItemAssignmentRepo.DB.Model(&entity.JobItem{}).Where("id = ?", jobItemAssignment.JobItemID).Update("assigned", true).Error
	if updateErr != nil {
		return nil, updateErr
	}

	return jobItemAssignment, nil
}

func (jobItemAssignmentRepo *JobItemAssignmentRepo) Get(id string) (*entity.JobItemAssignment, error) {
	jobItemAssignment := entity.JobItemAssignment{}
	getErr := jobItemAssignmentRepo.DB.Where("id = ?", id).Take(&jobItemAssignment).Error
	if getErr != nil {
		return nil, getErr
	}

	return &jobItemAssignment, nil
}

func (jobItemAssignmentRepo *JobItemAssignmentRepo) List(conditions string) ([]entity.JobItemAssignment, error) {
	jobItemAssignments := []entity.JobItemAssignment{}
	allJobItemAssignments := []entity.JobItemAssignment{}

	getErr := jobItemAssignmentRepo.DB.Where(conditions).Find(&jobItemAssignments).Error
	for _, jobItemAssignment := range jobItemAssignments {
		thisJobItemAssignment := entity.JobItemAssignment{}
		thisJobItemAssignment = jobItemAssignment
		jobItemID := jobItemAssignment.JobItemID
		jobItem := entity.JobItem{}
		jobItemAssignmentRepo.DB.Where("id = ?", jobItemID).Take(&jobItem)

		thisJobItemAssignment.JobItem = &jobItem
		allJobItemAssignments = append(allJobItemAssignments, thisJobItemAssignment)
	}
	if getErr != nil {
		return nil, getErr
	}
	return allJobItemAssignments, nil
}

func (jobItemAssignmentRepo *JobItemAssignmentRepo) Update(id string, update *entity.JobItemAssignment) (*entity.JobItemAssignment, error) {
	existingjobItemAssignment := entity.JobItemAssignment{}
	getErr := jobItemAssignmentRepo.DB.Where("id = ?", id).Take(&existingjobItemAssignment).Error
	if getErr != nil {
		return nil, getErr
	}

	updationErr := jobItemAssignmentRepo.DB.Table(entity.JobItemAssignment{}.Tablename()).Where("id = ?", id).Updates(update).Error
	if updationErr != nil {
		return nil, updationErr
	}

	updated := entity.JobItemAssignment{}
	jobItemAssignmentRepo.DB.Where("id = ?", id).Take(&updated)

	return &updated, nil
}
