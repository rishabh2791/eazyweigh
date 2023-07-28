package repository

import "eazyweigh/domain/entity"

type UnderIssueRepository interface {
	Create(underIssue *entity.UnderIssue) (*entity.UnderIssue, error)
	List(conditions string) ([]entity.UnderIssue, error)
	Update(id string, update *entity.UnderIssue) (*entity.UnderIssue, error)
}
