package repository

import "eazyweigh/domain/entity"

type StepTypeRepository interface {
	Create(stepType *entity.StepType) (*entity.StepType, error)
	Get(id string) (*entity.StepType, error)
	List(conditions string) ([]entity.StepType, error)
	Update(id string, stepType *entity.StepType) (*entity.StepType, error)
}
