package application

import "eazyweigh/domain/repository"

type CommonApp struct {
	commonRepository repository.CommonRepository
}

var _ CommonAppInterface = &CommonApp{}

func NewCommonApp(commonRepository repository.CommonRepository) *CommonApp {
	return &CommonApp{
		commonRepository: commonRepository,
	}
}

func (commonApp *CommonApp) GetTables() ([]string, error) {
	return commonApp.commonRepository.GetTables()
}

type CommonAppInterface interface {
	GetTables() ([]string, error)
}
