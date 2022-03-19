package repository

import "eazyweigh/domain/entity"

type CompanyRepository interface {
	Create(company *entity.Company) (*entity.Company, error)
	Get(id int) (*entity.Company, error)
	List(conditions string) ([]entity.Company, error)
	Update(id string, company *entity.Company) (*entity.Company, error)
}
