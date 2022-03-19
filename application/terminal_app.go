package application

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"
)

type TerminalApp struct {
	terminalRepository repository.TerminalRepository
}

func NewTerminalApp(terminalRepository repository.TerminalRepository) *TerminalApp {
	return &TerminalApp{
		terminalRepository: terminalRepository,
	}
}

type TerminalAppInterface interface {
	Create(termial *entity.Terminal) (*entity.Terminal, error)
	Get(id string) (*entity.Terminal, error)
	List(conditions string) ([]entity.Terminal, error)
	Update(id string, update *entity.Terminal) (*entity.Terminal, error)
}

func (terminalApp *TerminalApp) Create(termial *entity.Terminal) (*entity.Terminal, error) {
	return terminalApp.terminalRepository.Create(termial)
}

func (terminalApp *TerminalApp) Get(id string) (*entity.Terminal, error) {
	return terminalApp.terminalRepository.Get(id)
}

func (terminalApp *TerminalApp) List(conditions string) ([]entity.Terminal, error) {
	return terminalApp.terminalRepository.List(conditions)
}

func (terminalApp *TerminalApp) Update(id string, update *entity.Terminal) (*entity.Terminal, error) {
	return terminalApp.terminalRepository.Update(id, update)
}
