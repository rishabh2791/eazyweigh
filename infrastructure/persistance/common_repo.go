package persistance

import (
	"eazyweigh/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
)

type CommonRepo struct {
	DB     *gorm.DB
	Logger hclog.Logger
}

var _ repository.CommonRepository = &CommonRepo{}

func NewCommonRepo(db *gorm.DB, logger hclog.Logger) *CommonRepo {
	return &CommonRepo{
		DB:     db,
		Logger: logger,
	}
}

func (commonRepo *CommonRepo) GetTables() ([]string, error) {
	tables := []string{}

	getErr := commonRepo.DB.Table("information_schema.tables").Where("table_schema = ?", "eazyweigh").Pluck("table_name", &tables).Error

	if getErr != nil {
		return nil, getErr
	}

	return tables, nil
}
