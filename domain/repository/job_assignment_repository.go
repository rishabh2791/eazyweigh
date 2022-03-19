package repository

import "eazyweigh/domain/entity"

type JobAssignmentRepository interface {
	Create(jobAssignment *entity.JobAssignment) (*entity.JobAssignment, error)
	Get(id string) (*entity.JobAssignment, error)
	List(conditions string) ([]entity.JobAssignment, error)
	Update(id string, update *entity.JobAssignment) (*entity.JobAssignment, error)
}
