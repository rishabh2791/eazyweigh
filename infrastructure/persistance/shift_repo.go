package persistance

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ShiftRepo struct {
	DB     *gorm.DB
	Logger hclog.Logger
}

var _ repository.ShiftRepository = &ShiftRepo{}

func NewShiftRepo(db *gorm.DB, logger hclog.Logger) *ShiftRepo {
	return &ShiftRepo{
		DB:     db,
		Logger: logger,
	}
}

func (shiftRepo *ShiftRepo) Create(shift *entity.Shift) (*entity.Shift, error) {
	validationErr := shift.Validate()
	if validationErr != nil {
		return nil, validationErr
	}

	creationErr := shiftRepo.DB.Create(&shift).Error
	if creationErr != nil {
		return nil, creationErr
	}

	return shift, nil
}

func (shiftRepo *ShiftRepo) Get(id string) (*entity.Shift, error) {
	shift := entity.Shift{}

	getErr := shiftRepo.DB.Preload(clause.Associations).Where("id = ?", id).Take(&shift).Error
	if getErr != nil {
		return nil, getErr
	}

	return &shift, nil
}

func (shiftRepo *ShiftRepo) List(conditions string) ([]entity.Shift, error) {
	shifts := []entity.Shift{}

	getErr := shiftRepo.DB.
		Preload("Factory.Address").
		Preload("Factory.CreatedBy").
		Preload("Factory.CreatedBy.UserRole").
		Preload("Factory.UpdatedBy").
		Preload("Factory.UpdatedBy.UserRole").
		Preload("CreatedBy").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).Where(conditions).Find(&shifts).Error
	if getErr != nil {
		return nil, getErr
	}

	return shifts, nil
}
