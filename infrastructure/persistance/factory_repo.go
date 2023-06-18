package persistance

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
)

type FactoryRepo struct {
	DB     *gorm.DB
	Logger hclog.Logger
}

var _ repository.FactoryRepository = &FactoryRepo{}

func NewFactoryRepo(db *gorm.DB, logger hclog.Logger) *FactoryRepo {
	return &FactoryRepo{
		DB:     db,
		Logger: logger,
	}
}

func (factoryRepo *FactoryRepo) Create(factory *entity.Factory) (*entity.Factory, error) {
	validationErr := factory.Validate()
	if validationErr != nil {
		return nil, validationErr
	}

	creationErr := factoryRepo.DB.Create(&factory).Error
	if creationErr != nil {
		return nil, creationErr
	}

	return factory, nil
}

func (factoryRepo *FactoryRepo) Get(id string) (*entity.Factory, error) {
	factory := entity.Factory{}

	getErr := factoryRepo.DB.Where("id = ?", id).Take(&factory).Error
	if getErr != nil {
		return nil, getErr
	}

	return &factory, nil
}

func (factoryRepo *FactoryRepo) List(conditions string) ([]entity.Factory, error) {
	factories := []entity.Factory{}

	getErr := factoryRepo.DB.Where(conditions).Find(&factories).Error
	if getErr != nil {
		return nil, getErr
	}

	return factories, nil
}

func (factoryRepo *FactoryRepo) Update(id string, factory *entity.Factory) (*entity.Factory, error) {
	existingFactory := entity.Factory{}

	getErr := factoryRepo.DB.Where("id = ?", id).Take(&existingFactory).Error
	if getErr != nil {
		return nil, getErr
	}

	updationErr := factoryRepo.DB.Table(entity.Factory{}.Tablename()).Where("id = ?", id).Updates(factory).Error
	if updationErr != nil {
		return nil, getErr
	}

	updated := entity.Factory{}
	factoryRepo.DB.Where("id = ?", id).Take(&updated)

	return &updated, nil
}
