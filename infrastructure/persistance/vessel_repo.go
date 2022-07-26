package persistance

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type VesselRepo struct {
	DB     *gorm.DB
	Logger hclog.Logger
}

var _ repository.VesselRepository = &VesselRepo{}

func NewVesselRepo(db *gorm.DB, logger hclog.Logger) *VesselRepo {
	return &VesselRepo{
		DB:     db,
		Logger: logger,
	}
}

func (vesselRepo *VesselRepo) Create(vessel *entity.Vessel) (*entity.Vessel, error) {
	validationErr := vessel.Validate()
	if validationErr != nil {
		return nil, validationErr
	}

	creationErr := vesselRepo.DB.Create(&vessel).Error
	return vessel, creationErr
}

func (vesselRepo *VesselRepo) Get(id string) (*entity.Vessel, error) {
	vessel := entity.Vessel{}

	getErr := vesselRepo.DB.
		Preload("Factory.Address").
		Preload("Factory.CreatedBy").
		Preload("Factory.CreatedBy.UserRole").
		Preload("Factory.UpdatedBy").
		Preload("Factory.UpdatedBy.UserRole").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).Where("id = ?", id).First(&vessel).Error
	return &vessel, getErr
}

func (vesselRepo *VesselRepo) List(conditions string) ([]entity.Vessel, error) {
	vessels := []entity.Vessel{}

	getErr := vesselRepo.DB.
		Preload("Factory.Address").
		Preload("Factory.CreatedBy").
		Preload("Factory.CreatedBy.UserRole").
		Preload("Factory.UpdatedBy").
		Preload("Factory.UpdatedBy.UserRole").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).Where(conditions).Find(vessels).Error
	return vessels, getErr
}

func (vesselRepo *VesselRepo) Update(id string, vessel *entity.Vessel) (*entity.Vessel, error) {
	existingVessel := entity.Vessel{}
	getErr := vesselRepo.DB.
		Preload(clause.Associations).Where("id = ?", id).First(&existingVessel).Error
	if getErr != nil {
		return nil, getErr
	}

	updationErr := vesselRepo.DB.Table(vessel.Tablename()).Where("id = ?", id).Updates(&vessel).Error
	if updationErr != nil {
		return nil, updationErr
	}

	updated := entity.Vessel{}
	vesselRepo.DB.
		Preload("Factory.Address").
		Preload("Factory.CreatedBy").
		Preload("Factory.CreatedBy.UserRole").
		Preload("Factory.UpdatedBy").
		Preload("Factory.UpdatedBy.UserRole").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).Where("id = ?", id).First(&updated)
	return &updated, nil
}
