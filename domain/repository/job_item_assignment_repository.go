package repository

import "eazyweigh/domain/entity"

type JobItemAssignmentRepository interface {
	Create(jobItemAssignment *entity.JobItemAssignment) (*entity.JobItemAssignment, error)
	Get(id string) (*entity.JobItemAssignment, error)
	List(conditions string) ([]entity.JobItemAssignment, error)
	Update(id string, update *entity.JobItemAssignment) (*entity.JobItemAssignment, error)
}
