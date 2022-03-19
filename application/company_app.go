package application

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"
)

type CompanyApp struct {
	companyRepository repository.CompanyRepository
}

var _ CompanyAppInterface = &CompanyApp{}

func NewCompanyApp(companyRepository repository.CompanyRepository) *CompanyApp {
	return &CompanyApp{
		companyRepository: companyRepository,
	}
}

type CompanyAppInterface interface {
	Create(company *entity.Company) (*entity.Company, error)
	Get(id int) (*entity.Company, error)
	List(conditions string) ([]entity.Company, error)
	Update(id string, company *entity.Company) (*entity.Company, error)
}

func (companyApp *CompanyApp) Create(company *entity.Company) (*entity.Company, error) {
	return companyApp.companyRepository.Create(company)
}

func (companyApp *CompanyApp) Get(id int) (*entity.Company, error) {
	return companyApp.companyRepository.Get(id)
}

func (companyApp *CompanyApp) List(conditions string) ([]entity.Company, error) {
	return companyApp.companyRepository.List(conditions)
}

func (companyApp *CompanyApp) Update(id string, company *entity.Company) (*entity.Company, error) {
	return companyApp.companyRepository.Update(id, company)
}
