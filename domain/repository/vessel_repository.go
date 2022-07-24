package repository

import "eazyweigh/domain/entity"

type VesselRepository interface {
	Create(vessel *entity.Vessel) (*entity.Vessel, error)
	Get(id string) (*entity.Vessel, error)
	List(conditions string) ([]entity.Vessel, error)
	Update(id string, vessel *entity.Vessel) (*entity.Vessel, error)
}
