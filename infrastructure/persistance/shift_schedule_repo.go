package persistance

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (shift *ShiftScheduleRepo) Get(username string) ([]entity.ShiftSchedule, error) {
	shiftSchedule := []entity.ShiftSchedule{}

	getErr := shift.DB.Preload(clause.Associations).Where("user_username = ?", username).Find(&shiftSchedule).Error
	if getErr != nil {
		return nil, getErr
	}

	return shiftSchedule, nil
}

func (shift *ShiftScheduleRepo) List(conditions string) ([]entity.ShiftSchedule, error) {
	shiftSchedule := []entity.ShiftSchedule{}

	getErr := shift.DB.Preload(clause.Associations).Where(conditions).Find(&shiftSchedule).Error
	if getErr != nil {
		return nil, getErr
	}

	return shiftSchedule, nil
}

func (shift *ShiftScheduleRepo) Update(id string, shiftSchedule *entity.ShiftSchedule) (*entity.ShiftSchedule, error) {
	existingShift := entity.ShiftSchedule{}

	getErr := shift.DB.Preload(clause.Associations).Where("id = ?", id).Take(&existingShift).Error
	if getErr != nil {
		return nil, getErr
	}

	updationErr := shift.DB.Table(entity.ShiftSchedule{}.Tablename()).Where("id = ?", id).Updates(shiftSchedule).Error
	if updationErr != nil {
		return nil, updationErr
	}

	update := entity.ShiftSchedule{}
	shift.DB.Preload(clause.Associations).Where("id = ?", id).Take(&update)

	return &update, nil
}
