package application

import "eazyweigh/infrastructure/persistance"

type AppStore struct {
	AddressApp                 *AddressApp
	AuthApp                    *AuthApp
	BOMApp                     *BOMApp
	BOMItemApp                 *BOMItemApp
	CompanyApp                 *CompanyApp
	CommonApp                  *CommonApp
	FactoryApp                 *FactoryApp
	JobApp                     *JobApp
	JobItemApp                 *JobItemApp
	JobItemAssignmentApp       *JobItemAssignmentApp
	JobItemWeighingApp         *JobItemWeighingApp
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
	UserRoleAccessApp          *UserRoleAccessApp
	UserCompanyApp             *UserCompanyApp
	UserFactoryApp             *UserFactoryApp
}

func NewAppStore(repoStore *persistance.RepoStore) *AppStore {
	appStore := AppStore{}
	appStore.AddressApp = NewAddressApp(repoStore.AddressRepo)
	appStore.AuthApp = NewAuthApp(repoStore.AuthRepo)
	appStore.BOMApp = NewBOMApp(repoStore.BOMRepo)
	appStore.BOMItemApp = NewBOMItemApp(repoStore.BOMItemRepo)
	appStore.CompanyApp = NewCompanyApp(repoStore.CompanyRepo)
	appStore.CommonApp = NewCommonApp(repoStore.CommonRepo)
	appStore.FactoryApp = NewFactoryRepository(repoStore.FactoryRepo)
	appStore.JobApp = NewJobApp(repoStore.JobRepo)
	appStore.JobItemApp = NewJobItemApp(repoStore.JobItemRepo)
	appStore.JobItemAssignmentApp = NewJobItemAssignmentApp(repoStore.JobItemAssignmentRepo)
	appStore.JobItemWeighingApp = NewJobItemWeighingApp(repoStore.JobItemWeighingRepo)
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
	appStore.UserRoleAccessApp = NewUserRoleAccessApp(repoStore.UserRoleAccessRepo)
	appStore.UserCompanyApp = NewUserCompanyApp(repoStore.UserCompanyRepo)
	appStore.UserFactoryApp = NewUserFactoryApp(repoStore.UserFactoryRepo)
	return &appStore
}
