package repository

import "eazyweigh/domain/entity"

type FactoryRepository interface {
	Create(factory *entity.Factory) (*entity.Factory, error)
	Get(id string) (*entity.Factory, error)
	List(conditions string) ([]entity.Factory, error)
	Update(id string, factory *entity.Factory) (*entity.Factory, error)
}
