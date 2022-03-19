package repository

import "eazyweigh/domain/entity"

type UnitOfMeasureRepository interface {
	Create(uom *entity.UnitOfMeasure) (*entity.UnitOfMeasure, error)
	Get(id string) (*entity.UnitOfMeasure, error)
	List(conditions string) ([]entity.UnitOfMeasure, error)
	Update(id string, uom *entity.UnitOfMeasure) (*entity.UnitOfMeasure, error)
}
