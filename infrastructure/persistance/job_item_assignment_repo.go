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

	existingAssignment := entity.JobItemAssignment{}
	getErr := jobItemAssignmentRepo.DB.Preload(clause.Associations).Where("job_item_id = ?", jobItemAssignment.JobItemID).Take(&existingAssignment).Error
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
	getErr := jobItemAssignmentRepo.DB.
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
		Preload("ShiftSchedule.Shift").
		Preload("ShiftSchedule.Shift.Factory.Address").
		Preload("ShiftSchedule.Shift.Factory.CreatedBy").
		Preload("ShiftSchedule.Shift.Factory.CreatedBy.UserRole").
		Preload("ShiftSchedule.Shift.Factory.UpdatedBy").
		Preload("ShiftSchedule.Shift.Factory.UpdatedBy.UserRole").
		Preload("ShiftSchedule.Shift.CreatedBy").
		Preload("ShiftSchedule.Shift.CreatedBy.UserRole").
		Preload("ShiftSchedule.Shift.UpdatedBy").
		Preload("ShiftSchedule.Shift.UpdatedBy.UserRole").
		Preload("ShiftSchedule.User").
		Preload("ShiftSchedule.User.UserRole").
		Preload("ShiftSchedule.CreatedBy").
		Preload("ShiftSchedule.CreatedBy.UserRole").
		Preload("ShiftSchedule.UpdatedBy").
		Preload("ShiftSchedule.UpdatedBy.UserRole").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload("JobItem.JobItemWeighing").
		Preload("JobItem.JobItemWeighing.CreatedBy").
		Preload("JobItem.JobItemWeighing.CreatedBy.UserRole").
		Preload("JobItem.JobItemWeighing.UpdatedBy").
		Preload("JobItem.JobItemWeighing.UpdatedBy.UserRole").
		Preload(clause.Associations).Where("id = ?", id).Take(&jobItemAssignment).Error
	if getErr != nil {
		return nil, getErr
	}

	return &jobItemAssignment, nil
}

func (jobItemAssignmentRepo *JobItemAssignmentRepo) List(conditions string) ([]entity.JobItemAssignment, error) {
	jobItemAssignments := []entity.JobItemAssignment{}
	allJobItemAssignments := []entity.JobItemAssignment{}

	getErr := jobItemAssignmentRepo.DB.
		Preload("ShiftSchedule.Shift").
		Preload("ShiftSchedule.Shift.Factory.Address").
		Preload("ShiftSchedule.Shift.Factory.CreatedBy").
		Preload("ShiftSchedule.Shift.Factory.CreatedBy.UserRole").
		Preload("ShiftSchedule.Shift.Factory.UpdatedBy").
		Preload("ShiftSchedule.Shift.Factory.UpdatedBy.UserRole").
		Preload("ShiftSchedule.Shift.CreatedBy").
		Preload("ShiftSchedule.Shift.CreatedBy.UserRole").
		Preload("ShiftSchedule.Shift.UpdatedBy").
		Preload("ShiftSchedule.Shift.UpdatedBy.UserRole").
		Preload("ShiftSchedule.User").
		Preload("ShiftSchedule.User.UserRole").
		Preload("ShiftSchedule.CreatedBy").
		Preload("ShiftSchedule.CreatedBy.UserRole").
		Preload("ShiftSchedule.UpdatedBy").
		Preload("ShiftSchedule.UpdatedBy.UserRole").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).Where(conditions).Find(&jobItemAssignments).Error
		for _, jobItemAssignment:= range jobItemAssignments{
			thisJobItemAssignment := entity.JobItemAssignment{}
			thisJobItemAssignment = jobItemAssignment
			jobItemID := jobItemAssignment.JobItemID			
			jobItem := entity.JobItem{}
			jobItemErr := jobItemAssignmentRepo.DB.
				Preload("Material.UnitOfMeasure").
				Preload("Material.UnitOfMeasure.Factory").
				Preload("Material.UnitOfMeasure.Factory.CreatedBy").
				Preload("Material.UnitOfMeasure.Factory.CreatedBy.UserRole").
				Preload("Material.UnitOfMeasure.Factory.UpdatedBy").
				Preload("Material.UnitOfMeasure.Factory.UpdatedBy.UserRole").
				Preload("Material.UnitOfMeasure.CreatedBy").
				Preload("Material.UnitOfMeasure.CreatedBy.UserRole").
				Preload("Material.UnitOfMeasure.UpdatedBy").
				Preload("Material.UnitOfMeasure.UpdatedBy.UserRole").
				Preload("Material.UnitOfMeasure.Factory.Address").
				Preload("UnitOfMeasure.Factory").
				Preload("UnitOfMeasure.Factory.CreatedBy").
				Preload("UnitOfMeasure.Factory.CreatedBy.UserRole").
				Preload("UnitOfMeasure.Factory.UpdatedBy").
				Preload("UnitOfMeasure.Factory.UpdatedBy.UserRole").
				Preload("UnitOfMeasure.CreatedBy").
				Preload("UnitOfMeasure.CreatedBy.UserRole").
				Preload("UnitOfMeasure.UpdatedBy").
				Preload("UnitOfMeasure.UpdatedBy.UserRole").
				Preload("UnitOfMeasure.Factory.Address").
				Preload("CreatedBy.UserRole").
				Preload("UpdatedBy.UserRole").
				Preload("Material.CreatedBy").
				Preload("Material.CreatedBy.UserRole").
				Preload("Material.UpdatedBy").
				Preload("Material.UpdatedBy.UserRole").
				Preload(clause.Associations).Where("id = ?",jobItemID).Take(&jobItem).Error			
			if jobItemErr != nil {
				
			}
			thisJobItemAssignment.JobItem=&jobItem
			allJobItemAssignments = append(allJobItemAssignments, thisJobItemAssignment)
		}
	if getErr != nil {
		return nil, getErr
	}
	return allJobItemAssignments, nil
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
	jobItemAssignmentRepo.DB.Preload(clause.Associations).Where("id = ?", id).Take(&updated)

	return &updated, nil
}
