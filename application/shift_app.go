package application

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"
)

type ShiftApp struct {
	shiftRepository repository.ShiftRepository
}

var _ ShiftAppInterface = &ShiftApp{}

func NewShiftApp(shiftRepository repository.ShiftRepository) *ShiftApp {
	return &ShiftApp{
		shiftRepository: shiftRepository,
	}
}

type ShiftAppInterface interface {
	Create(shift *entity.Shift) (*entity.Shift, error)
	Get(id string) (*entity.Shift, error)
	List(conditions string) ([]entity.Shift, error)
}

func (shiftApp *ShiftApp) Create(shift *entity.Shift) (*entity.Shift, error) {
	return shiftApp.shiftRepository.Create(shift)
}

func (shiftApp *ShiftApp) Get(id string) (*entity.Shift, error) {
	return shiftApp.shiftRepository.Get(id)
}

func (shiftApp *ShiftApp) List(conditions string) ([]entity.Shift, error) {
	return shiftApp.shiftRepository.List(conditions)
}
