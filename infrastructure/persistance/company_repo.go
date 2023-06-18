package persistance

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
)

type CompanyRepo struct {
	DB     *gorm.DB
	Logger hclog.Logger
}

var _ repository.CompanyRepository = &CompanyRepo{}

func NewCompanyRepo(db *gorm.DB, logger hclog.Logger) *CompanyRepo {
	return &CompanyRepo{
		DB:     db,
		Logger: logger,
	}
}

func (companyRepo *CompanyRepo) Create(company *entity.Company) (*entity.Company, error) {
	validationErr := company.Validate()
	if validationErr != nil {
		return nil, validationErr
	}

	creationErr := companyRepo.DB.Create(&company).Error
	if creationErr != nil {
		return nil, creationErr
	}

	return company, nil
}

func (companyRepo *CompanyRepo) Get(id int) (*entity.Company, error) {
	company := entity.Company{}

	getErr := companyRepo.DB.Where("id = ?", id).Take(&company).Error
	if getErr != nil {
		return nil, getErr
	}

	return &company, nil
}

func (companyRepo *CompanyRepo) List(conditions string) ([]entity.Company, error) {
	companies := []entity.Company{}

	getErr := companyRepo.DB.Find(&companies).Error
	if getErr != nil {
		return nil, getErr
	}

	return companies, nil
}

func (companyRepo *CompanyRepo) Update(id string, company *entity.Company) (*entity.Company, error) {
	existingCompany := entity.Company{}

	getErr := companyRepo.DB.Where("id = ?", id).Take(&existingCompany).Error
	if getErr != nil {
		return nil, getErr
	}

	updationErr := companyRepo.DB.Table(entity.Company{}.Tablename()).Where("id = ?", id).Updates(company).Error
	if updationErr != nil {
		return nil, updationErr
	}

	updated := entity.Company{}
	companyRepo.DB.Take(&updated)

	return &updated, nil
}
