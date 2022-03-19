package repository

import "eazyweigh/domain/entity"

type BOMRepository interface {
	Create(bom *entity.BOM) (*entity.BOM, error)
	Get(id string) (*entity.BOM, error)
	List(conditions string) ([]entity.BOM, error)
	Update(id string, update *entity.BOM) (*entity.BOM, error)
}
