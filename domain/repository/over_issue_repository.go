package repository

import "eazyweigh/domain/entity"

type OverIssueRepository interface {
	Create(overIssue *entity.OverIssue) (*entity.OverIssue, error)
	List(conditions string) ([]entity.OverIssue, error)
	Update(id string, update *entity.OverIssue) (*entity.OverIssue, error)
}
