package repository

import "eazyweigh/domain/entity"

type AddressRepository interface {
	Create(address *entity.Address) (*entity.Address, error)
	List(conditions string) ([]entity.Address, error)
	Update(id string, address *entity.Address) (*entity.Address, error)
	Delete(id string) error
}
