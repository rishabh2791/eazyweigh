package application

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"
)

type UnderIssueApp struct {
	underIssueRepository repository.UnderIssueRepository
}

var _ UnderIssueAppInterface = &UnderIssueApp{}

func NewUnderIssueApp(underIssueReporitory repository.UnderIssueRepository) *UnderIssueApp {
	return &UnderIssueApp{
		underIssueRepository: underIssueReporitory,
	}
}

type UnderIssueAppInterface interface {
	Create(underIssue *entity.UnderIssue) (*entity.UnderIssue, error)
	List(jobID string) ([]entity.UnderIssue, error)
	Update(id string, update *entity.UnderIssue) (*entity.UnderIssue, error)
}

func (underIssueApp *UnderIssueApp) Create(underIssue *entity.UnderIssue) (*entity.UnderIssue, error) {
	return underIssueApp.underIssueRepository.Create(underIssue)
}

func (underIssueApp *UnderIssueApp) List(jobID string) ([]entity.UnderIssue, error) {
	return underIssueApp.underIssueRepository.List(jobID)
}

func (underIssueApp *UnderIssueApp) Update(id string, update *entity.UnderIssue) (*entity.UnderIssue, error) {
	return underIssueApp.underIssueRepository.Update(id, update)
}
