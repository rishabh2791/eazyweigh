package application

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"
)

type OverIssueApp struct {
	overIssueRepository repository.OverIssueRepository
}

var _ OverIssueAppInterface = &OverIssueApp{}

func NewOverIssueApp(overIssueRepository repository.OverIssueRepository) *OverIssueApp {
	return &OverIssueApp{
		overIssueRepository: overIssueRepository,
	}
}

type OverIssueAppInterface interface {
	Create(overIssue *entity.OverIssue) (*entity.OverIssue, error)
	List(conditions string) ([]entity.OverIssue, error)
	Update(id string, update *entity.OverIssue) (*entity.OverIssue, error)
}

func (overIssueApp *OverIssueApp) Create(overIssue *entity.OverIssue) (*entity.OverIssue, error) {
	return overIssueApp.overIssueRepository.Create(overIssue)
}

func (overIssueApp *OverIssueApp) List(conditions string) ([]entity.OverIssue, error) {
	return overIssueApp.overIssueRepository.List(conditions)
}

func (overIssueApp *OverIssueApp) Update(id string, update *entity.OverIssue) (*entity.OverIssue, error) {
	return overIssueApp.overIssueRepository.Update(id, update)
}
