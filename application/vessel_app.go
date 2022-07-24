package application

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"
)

type VesselApp struct {
	vesselRepository repository.VesselRepository
}

var _ VesselAppInterface = &VesselApp{}

func NewVesselApp(vesselRepository repository.VesselRepository) *VesselApp {
	return &VesselApp{
		vesselRepository: vesselRepository,
	}
}

func (vesselApp *VesselApp) Create(vessel *entity.Vessel) (*entity.Vessel, error) {
	return vesselApp.vesselRepository.Create(vessel)
}

func (vesselApp *VesselApp) Get(id string) (*entity.Vessel, error) {
	return vesselApp.vesselRepository.Get(id)
}

func (vesselApp *VesselApp) List(conditions string) ([]entity.Vessel, error) {
	return vesselApp.vesselRepository.List(conditions)
}

func (vesselApp *VesselApp) Update(id string, vessel *entity.Vessel) (*entity.Vessel, error) {
	return vesselApp.vesselRepository.Update(id, vessel)
}

type VesselAppInterface interface {
	Create(vessel *entity.Vessel) (*entity.Vessel, error)
	Get(id string) (*entity.Vessel, error)
	List(conditions string) ([]entity.Vessel, error)
	Update(id string, vessel *entity.Vessel) (*entity.Vessel, error)
}
