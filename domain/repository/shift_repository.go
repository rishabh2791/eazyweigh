package repository

import "eazyweigh/domain/entity"

type ShiftRepository interface {
	Create(shift *entity.Shift) (*entity.Shift, error)
	Get(id string) (*entity.Shift, error)
	List(conditions string) ([]entity.Shift, error)
}
