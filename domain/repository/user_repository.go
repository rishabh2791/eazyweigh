package repository

import "eazyweigh/domain/entity"

type UserRepository interface {
	Create(user *entity.User, action string) (*entity.User, error)
	Get(username string) (*entity.User, error)
	List(conditions string) ([]entity.User, error)
	Update(username string, update map[string]interface{}, user string) (*entity.User, error)
}
