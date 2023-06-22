package persistance

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type JobItemRepo struct {
	DB     *gorm.DB
	Logger hclog.Logger
}

var _ repository.JobItemRepository = &JobItemRepo{}

func NewJobItemRepo(db *gorm.DB, logger hclog.Logger) *JobItemRepo {
	return &JobItemRepo{
		DB:     db,
		Logger: logger,
	}
}

func (jobItemRepo *JobItemRepo) Create(jobItem *entity.JobItem) (*entity.JobItem, error) {
	validationErr := jobItem.Validate()
	if validationErr != nil {
		return nil, validationErr
	}

	creationErr := jobItemRepo.DB.Create(&jobItem).Error
	if creationErr != nil {
		return nil, creationErr
	}

	return jobItem, nil
}

func (jobItemRepo *JobItemRepo) Get(conditions string) ([]entity.JobItem, error) {
	jobItems := []entity.JobItem{}

	getErr := jobItemRepo.DB.
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
		Preload(clause.Associations).Where(conditions).Find(&jobItems).Error
	if getErr != nil {
		return nil, getErr
	}

	return jobItems, nil
}

func (jobItemRepo *JobItemRepo) Update(id string, jobItem *entity.JobItem) (*entity.JobItem, error) {
	existingItem := entity.JobItem{}
	getErr := jobItemRepo.DB.Preload(clause.Associations).Where("id = ?", id).Take(&existingItem).Error
	if getErr != nil {
		return nil, getErr
	}

	updationErr := jobItemRepo.DB.Table(entity.JobItem{}.Tablename()).Where("id = ?", id).Updates(jobItem).Error
	if updationErr != nil {
		return nil, updationErr
	}

	updated := entity.JobItem{}
	jobItemRepo.DB.Preload(clause.Associations).Where("id = ?", id).Take(&updated)

	return &updated, nil
}
