package application

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"
)

type UserCompanyApp struct {
	userCompanyRepository repository.UserCompanyRepository
}

var _ UserCompanyAppInterface = &UserCompanyApp{}

func NewUserCompanyApp(userCompanyRepository repository.UserCompanyRepository) *UserCompanyApp {
	return &UserCompanyApp{
		userCompanyRepository: userCompanyRepository,
	}
}

type UserCompanyAppInterface interface {
	Create(userCompany *entity.UserCompany) (*entity.UserCompany, error)
	Get(companyID string) ([]entity.UserCompany, error)
}

func (userCompanyApp *UserCompanyApp) Create(userCompany *entity.UserCompany) (*entity.UserCompany, error) {
	return userCompanyApp.userCompanyRepository.Create(userCompany)
}

func (userCompanyApp *UserCompanyApp) Get(companyID string) ([]entity.UserCompany, error) {
	return userCompanyApp.userCompanyRepository.Get(companyID)
}
