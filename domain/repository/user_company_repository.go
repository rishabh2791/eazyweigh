package repository

import "eazyweigh/domain/entity"

type UserCompanyRepository interface {
	Create(userCompany *entity.UserCompany) (*entity.UserCompany, error)
	Get(conditions string) ([]entity.UserCompany, error)
}
