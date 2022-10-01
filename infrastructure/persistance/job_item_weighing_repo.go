package persistance

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"
	"math"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type JobItemWeighingRepo struct {
	DB     *gorm.DB
	Logger hclog.Logger
}

var _ repository.JobItemWeighingRepository = &JobItemWeighingRepo{}

func NewJobItemWeighingRepo(db *gorm.DB, logger hclog.Logger) *JobItemWeighingRepo {
	return &JobItemWeighingRepo{
		DB:     db,
		Logger: logger,
	}
}

func (jobItemWeighingRepo *JobItemWeighingRepo) Create(jobItemWeight *entity.JobItemWeighing) (*entity.JobItemWeighing, error) {
	validationErr := jobItemWeight.Validate()
	if validationErr != nil {
		return nil, validationErr
	}

	creationErr := jobItemWeighingRepo.DB.Create(&jobItemWeight).Error
	if creationErr != nil {
		return nil, creationErr
	}

	jobItem := entity.JobItem{}
	getErr := jobItemWeighingRepo.DB.Where("id = ?", jobItemWeight.JOBItemID).Take(&jobItem).Error
	if getErr != nil {
		return nil, getErr
	}

	update := map[string]interface{}{
		"actual_weight":       jobItem.ActualWeight + jobItemWeight.Weight,
		"updated_by_username": jobItemWeight.UpdatedByUsername,
	}

	if math.Round((float64(jobItem.ActualWeight)+float64(jobItemWeight.Weight))*1000)/1000 <= math.Round(float64(jobItem.UpperBound)*1000)/1000 && math.Round((float64(jobItem.ActualWeight)+float64(jobItemWeight.Weight))*1000)/1000 >= math.Round(float64(jobItem.LowerBound)*1000)/1000 {
		update["complete"] = true
	}

	updationErr := jobItemWeighingRepo.DB.Model(&entity.JobItem{}).Where("id = ?", jobItemWeight.JOBItemID).Updates(update).Error
	if updationErr != nil {
		return nil, updationErr
	}

	return jobItemWeight, nil
}

func (jobItemWeighingRepo *JobItemWeighingRepo) List(jobItemID string) ([]entity.JobItemWeighing, error) {
	jobItemWeighings := []entity.JobItemWeighing{}

	getErr := jobItemWeighingRepo.DB.
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
		Preload("JobItem.Material.CreatedBy").
		Preload("JobItem.Material.CreatedBy.UserRole").
		Preload("JobItem.Material.UpdatedBy").
		Preload("JobItem.Material.UpdatedBy.UserRole").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).Where("job_item_id = ?", jobItemID).Find(&jobItemWeighings).Error
	if getErr != nil {
		return nil, getErr
	}

	return jobItemWeighings, nil
}

func (jobItemWeighingRepo *JobItemWeighingRepo) Update(id string, jobItemWeighing *entity.JobItemWeighing) (*entity.JobItemWeighing, error) {
	existingJobItemWeighing := entity.JobItemWeighing{}
	getErr := jobItemWeighingRepo.DB.Preload(clause.Associations).Where("id = ?", id).Take(&existingJobItemWeighing).Error
	if getErr != nil {
		return nil, getErr
	}

	updationErr := jobItemWeighingRepo.DB.Table(entity.JobItemWeighing{}.Tablename()).Where("id = ?", id).Updates(jobItemWeighing).Error
	if updationErr != nil {
		return nil, updationErr
	}

	updated := entity.JobItemWeighing{}
	jobItemWeighingRepo.DB.Preload(clause.Associations).Where("id = ?", id).Take(&updated)

	jobItemID := existingJobItemWeighing.JOBItemID
	jobItemWeighings := []entity.JobItemWeighing{}
	weighingsErr := jobItemWeighingRepo.DB.Where("job_item_id = ? AND verified = FALSE", jobItemID).Find(&jobItemWeighings).Error
	if weighingsErr != nil {
		return nil, weighingsErr
	}
	if len(jobItemWeighings) == 0 {
		jobItemWeighingRepo.DB.Table(entity.JobItem{}.Tablename()).Where("id = ?", jobItemID).Update("verified", true)
	}

	return &updated, nil
}

func (jobItemWeighingRepo *JobItemWeighingRepo) Details(conditions string) ([]entity.WeighingBatch, error) {
	data := []entity.WeighingBatch{}
	queryString := "SELECT * FROM job_item_weighings LEFT JOIN (SELECT job_items.id as job_item_id, job_items.material_id as job_item_material_id, job_items.required_weight as required_weight, job_items.actual_weight as actual_weight,jobs.job_code as job_code, jobs.material_id as job_material_id FROM job_items LEFT JOIN jobs ON job_items.job_id = jobs.id) as job_items ON job_items.job_item_id = job_item_weighings.job_item_id"
	if conditions != "" || len(conditions) != 0 {
		queryString += " WHERE " + conditions
	} else {
		queryString += ";"
	}
	getErr := jobItemWeighingRepo.DB.Raw(queryString).Find(&data).Error
	return data, getErr
}
