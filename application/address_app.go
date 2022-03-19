package application

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"
)

type AddressApp struct {
	addressRepo repository.AddressRepository
}

var _ AddressAppInterface = &AddressApp{}

func NewAddressApp(addressRepo repository.AddressRepository) *AddressApp {
	return &AddressApp{
		addressRepo: addressRepo,
	}
}

type AddressAppInterface interface {
	Create(address *entity.Address) (*entity.Address, error)
	List(conditions string) ([]entity.Address, error)
	Update(id string, address *entity.Address) (*entity.Address, error)
	Delete(id string) error
}

func (addressApp *AddressApp) Create(address *entity.Address) (*entity.Address, error) {
	return addressApp.addressRepo.Create(address)
}

func (addressApp *AddressApp) List(conditions string) ([]entity.Address, error) {
	return addressApp.addressRepo.List(conditions)
}

func (addressApp *AddressApp) Update(id string, address *entity.Address) (*entity.Address, error) {
	return addressApp.addressRepo.Update(id, address)
}

func (addressApp *AddressApp) Delete(id string) error {
	return addressApp.addressRepo.Delete(id)
}
