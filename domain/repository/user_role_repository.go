package repository

import "eazyweigh/domain/entity"

type UserRoleRepository interface {
	Create(userRole *entity.UserRole) (*entity.UserRole, error)
	Get(userRole string) (*entity.UserRole, error)
	List(conditions string) ([]entity.UserRole, error)
	Update(userRole string, update *entity.UserRole) (*entity.UserRole, error)
}
