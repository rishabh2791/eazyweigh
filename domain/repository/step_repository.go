package repository

import "eazyweigh/domain/entity"

type StepRepository interface {
	Create(step *entity.Step) (*entity.Step, error)
	Get(id string) (*entity.Step, error)
	List(conditions string) ([]entity.Step, error)
	Update(id string, step *entity.Step) (*entity.Step, error)
}
