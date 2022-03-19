package repository

import "eazyweigh/domain/entity"

type TerminalRepository interface {
	Create(termial *entity.Terminal) (*entity.Terminal, error)
	Get(id string) (*entity.Terminal, error)
	List(conditions string) ([]entity.Terminal, error)
	Update(id string, update *entity.Terminal) (*entity.Terminal, error)
}
