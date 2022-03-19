package repository

import "eazyweigh/domain/entity"

type ShiftScheduleRepository interface {
	Create(shiftSchedule *entity.ShiftSchedule) (*entity.ShiftSchedule, error)
	Get(username string) ([]entity.ShiftSchedule, error)
	List(conditions string) ([]entity.ShiftSchedule, error)
	Update(id string, shiftSchedule *entity.ShiftSchedule) (*entity.ShiftSchedule, error)
}
