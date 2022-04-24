package application

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"
)

type ShiftScheduleApp struct {
	shiftScheduleRepository repository.ShiftScheduleRepository
}

var _ ShiftScheduleAppInterface = &ShiftScheduleApp{}

func NewShiftScheduleApp(shiftScheduleRepository repository.ShiftScheduleRepository) *ShiftScheduleApp {
	return &ShiftScheduleApp{
		shiftScheduleRepository: shiftScheduleRepository,
	}
}

type ShiftScheduleAppInterface interface {
	Create(shiftSchedule *entity.ShiftSchedule) (*entity.ShiftSchedule, error)
	List(conditions string) ([]entity.ShiftSchedule, error)
	Delete(id string) error
}

func (shift *ShiftScheduleApp) Create(shiftSchedule *entity.ShiftSchedule) (*entity.ShiftSchedule, error) {
	return shift.shiftScheduleRepository.Create(shiftSchedule)
}

func (shift *ShiftScheduleApp) List(conditions string) ([]entity.ShiftSchedule, error) {
	return shift.shiftScheduleRepository.List(conditions)
}

func (shift *ShiftScheduleApp) Delete(id string) error {
	return shift.shiftScheduleRepository.Delete(id)
}
