package application

import "eazyweigh/infrastructure/persistance"

type AppStore struct {
	AddressApp                 *AddressApp
	AuthApp                    *AuthApp
	BOMApp                     *BOMApp
	BOMItemApp                 *BOMItemApp
	CompanyApp                 *CompanyApp
	FactoryApp                 *FactoryApp
	JobApp                     *JobApp
	JobItemApp                 *JobItemApp
	JobAssignmentApp           *JobAssignmentApp
	MaterialApp                *MaterialApp
	OverIssueApp               *OverIssueApp
	ScannedDataApp             *ScannedDataApp
	ShiftApp                   *ShiftApp
	ShiftScheduleApp           *ShiftScheduleApp
	UnitOfMeasureApp           *UnitOfMeasureApp
	TerminalApp                *TerminalApp
	UnderIssueApp              *UnderIssueApp
	UnitOfMeasureConversionApp *UnitOfMeasureConversionApp
	UserApp                    *UserApp
	UserRoleApp                *UserRoleApp
}

func NewAppStore(repoStore *persistance.RepoStore) *AppStore {
	appStore := AppStore{}
	appStore.AddressApp = NewAddressApp(repoStore.AddressRepo)
	appStore.AuthApp = NewAuthApp(repoStore.AuthRepo)
	appStore.BOMApp = NewBOMApp(repoStore.BOMRepo)
	appStore.BOMItemApp = NewBOMItemApp(repoStore.BOMItemRepo)
	appStore.CompanyApp = NewCompanyApp(repoStore.CompanyRepo)
	appStore.FactoryApp = NewFactoryRepository(repoStore.FactoryRepo)
	appStore.JobApp = NewJobApp(repoStore.JobRepo)
	appStore.JobItemApp = NewJobItemApp(repoStore.JobItemRepo)
	appStore.JobAssignmentApp = NewJobAssignmentApp(repoStore.JobAssignmentRepo)
	appStore.MaterialApp = NewMaterialApp(repoStore.MaterialRepo)
	appStore.OverIssueApp = NewOverIssueApp(repoStore.OverIssueRepo)
	appStore.ScannedDataApp = NewScannedDataApp(repoStore.ScannedDataRepo)
	appStore.ShiftApp = NewShiftApp(repoStore.ShiftRepo)
	appStore.ShiftScheduleApp = NewShiftScheduleApp(repoStore.ShiftScheduleRepo)
	appStore.TerminalApp = NewTerminalApp(repoStore.TerminalRepo)
	appStore.UnderIssueApp = NewUnderIssueApp(repoStore.UnderIssueRepo)
	appStore.UnitOfMeasureApp = NewUnitOfMeasureApp(repoStore.UnitOfMeasureRepo)
	appStore.UnitOfMeasureConversionApp = NewUnitOfMeasureConversionApp(repoStore.UnitOfMeasureConversionRepo)
	appStore.UserApp = NewUserApp(repoStore.UserRepo)
	appStore.UserRoleApp = NewUserRoleApp(repoStore.UserRoleRepo)
	return &appStore
}
