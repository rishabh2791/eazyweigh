package repository

import "eazyweigh/domain/entity"

type ProcessRepository interface {
	Create(process *entity.Process) (*entity.Process, error)
	Get(materialID string) (*entity.Process, error)
	List(conditions string) ([]entity.Process, error)
	Update(id string, process *entity.Process) (*entity.Process, error)
}
