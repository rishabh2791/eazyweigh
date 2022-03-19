package repository

import "eazyweigh/domain/entity"

type MaterialRepository interface {
	Create(material *entity.Material) (*entity.Material, error)
	Get(id string) (*entity.Material, error)
	List(conditions string) ([]entity.Material, error)
	Update(id string, update *entity.Material) (*entity.Material, error)
}
