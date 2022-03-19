package repository

import "eazyweigh/domain/entity"

type UnitOfMeasureConversionRepository interface {
	Create(conversion *entity.UnitOfMeasureConversion) (*entity.UnitOfMeasureConversion, error)
	Get(id string) (*entity.UnitOfMeasureConversion, error)
	List(conditions string) ([]entity.UnitOfMeasureConversion, error)
}
