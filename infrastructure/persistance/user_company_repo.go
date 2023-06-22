package persistance

import (
	"eazyweigh/domain/entity"
	"eazyweigh/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserCompanyRepo struct {
	DB     *gorm.DB
	Logger hclog.Logger
}

var _ repository.UserCompanyRepository = &UserCompanyRepo{}

func NewUserCompanyRepo(db *gorm.DB, logger hclog.Logger) *UserCompanyRepo {
	return &UserCompanyRepo{
		DB:     db,
		Logger: logger,
	}
}

func (userCompanyRepo *UserCompanyRepo) Create(userCompany *entity.UserCompany) (*entity.UserCompany, error) {
	validationErr := userCompany.Validate()
	if validationErr != nil {
		return nil, validationErr
	}

	creationErr := userCompanyRepo.DB.Create(&userCompany).Error
	if creationErr != nil {
		return nil, creationErr
	}
	return userCompany, nil
}

func (userCompanyRepo *UserCompanyRepo) Get(conditions string) ([]entity.UserCompany, error) {
	companyUsers := []entity.UserCompany{}
	getErr := userCompanyRepo.DB.
		Preload("User.UserRole").
		Preload("Company.CreatedBy").
		Preload("Company.UpdatedBy").
		Preload("Company.CreatedBy.UserRole").
		Preload("Company.UpdatedBy.UserRole").
		Preload(clause.Associations).Where(conditions).Find(&companyUsers).Error
	if getErr != nil {
		return nil, getErr
	}

	return companyUsers, nil
}
