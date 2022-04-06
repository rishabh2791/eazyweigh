package persistance

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
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
		"actual_weight": jobItem.ActualWeight + jobItemWeight.Weight,
	}

	if (jobItem.ActualWeight+jobItemWeight.Weight) <= jobItem.UpperBound && (jobItem.ActualWeight+jobItemWeight.Weight) >= jobItem.LowerBound {
		update["complete"] = true
	}

	updationErr := jobItemWeighingRepo.DB.Model(&entity.JobItem{}).Where("id = ?", jobItemWeight.JOBItemID).Updates(update).Error
	if updationErr != nil {
		return nil, updationErr
	}

	return jobItemWeight, nil
}
