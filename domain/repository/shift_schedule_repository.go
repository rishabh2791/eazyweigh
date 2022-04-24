package repository

import "eazyweigh/domain/entity"

type ShiftScheduleRepository interface {
	Create(shiftSchedule *entity.ShiftSchedule) (*entity.ShiftSchedule, error)
	List(conditions string) ([]entity.ShiftSchedule, error)
	Delete(id string) error
}
