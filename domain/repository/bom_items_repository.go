package repository

import "eazyweigh/domain/entity"

type BOMItemsRepository interface {
	Create(bomItems *entity.BOMItem) (*entity.BOMItem, error)
	Get(id string) (*entity.BOMItem, error)
	List(conditions string) ([]entity.BOMItem, error)
	Update(id string, update *entity.BOMItem) (*entity.BOMItem, error)
}
