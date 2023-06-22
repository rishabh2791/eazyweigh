package persistance

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AddressRepo struct {
	DB     *gorm.DB
	Logger hclog.Logger
}

var _ repository.AddressRepository = &AddressRepo{}

func NewAddressRepo(db *gorm.DB, logger hclog.Logger) *AddressRepo {
	return &AddressRepo{
		DB:     db,
		Logger: logger,
	}

}

func (addressRepo *AddressRepo) Create(address *entity.Address) (*entity.Address, error) {
	validationErr := address.Validate()
	if validationErr != nil {
		return nil, validationErr
	}
	creationErr := addressRepo.DB.Create(&address).Error
	if creationErr != nil {
		return nil, creationErr
	}
	return address, nil
}

func (addressRepo *AddressRepo) List(conditions string) ([]entity.Address, error) {
	addresses := []entity.Address{}
	getErr := addressRepo.DB.Preload(clause.Associations).Where(conditions).Find(&addresses).Error
	if getErr != nil {
		return nil, getErr
	}
	return addresses, nil
}

func (addressRepo *AddressRepo) Update(id string, address *entity.Address) (*entity.Address, error) {
	existingAddress := entity.Address{}

	err := addressRepo.DB.Where("id = ?", id).Take(&existingAddress).Error
	if err != nil {
		return nil, err
	}

	updationErr := addressRepo.DB.Table(entity.Address{}.Tablename()).Where("id = ?", id).Updates(address).Error
	if updationErr != nil {
		return nil, updationErr
	}

	updated := entity.Address{}
	addressRepo.DB.Where("id = ?", id).Take(&updated)

	return &updated, nil
}

func (addressRepo *AddressRepo) Delete(id string) error {
	existingAddress := entity.Address{}

	err := addressRepo.DB.Where("id = ?", id).Take(&existingAddress).Error
	if err != nil {
		return err
	}
	deletionErr := addressRepo.DB.Where("id = ?", id).Delete(&existingAddress).Error
	return deletionErr
}
