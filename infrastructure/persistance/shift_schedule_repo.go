package persistance

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
)

type ShiftScheduleRepo struct {
	DB     *gorm.DB
	Logger hclog.Logger
}

var _ repository.ShiftScheduleRepository = &ShiftScheduleRepo{}

func NewShiftScheduleRepo(db *gorm.DB, logger hclog.Logger) *ShiftScheduleRepo {
	return &ShiftScheduleRepo{
		DB:     db,
		Logger: logger,
	}
}

func (shift *ShiftScheduleRepo) Create(shiftSchedule *entity.ShiftSchedule) (*entity.ShiftSchedule, error) {
	validationErr := shiftSchedule.Validate()
	if validationErr != nil {
		return nil, validationErr
	}

	creationErr := shift.DB.Create(&shiftSchedule).Error
	if creationErr != nil {
		return nil, creationErr
	}

	return shiftSchedule, nil
}

func (shift *ShiftScheduleRepo) List(conditions string) ([]entity.ShiftSchedule, error) {
	shiftSchedule := []entity.ShiftSchedule{}

	getErr := shift.DB.Where(conditions).Find(&shiftSchedule).Error
	if getErr != nil {
		return nil, getErr
	}

	return shiftSchedule, nil
}

func (shift *ShiftScheduleRepo) Delete(id string) error {
	deleteErr := shift.DB.Where("id = ?", id).Delete(&entity.ShiftSchedule{}).Error
	return deleteErr
}
