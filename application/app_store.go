package application

import "eazyweigh/infrastructure/persistance"

type AppStore struct {
	AddressApp                 *AddressApp
	AuthApp                    *AuthApp
	BatchApp                   *BatchApp
	BOMApp                     *BOMApp
	BOMItemApp                 *BOMItemApp
	CompanyApp                 *CompanyApp
	CommonApp                  *CommonApp
	DeviceApp                  *DeviceApp
	DeviceDataApp              *DeviceDataApp
	DeviceTypeApp              *DeviceTypeApp
	FactoryApp                 *FactoryApp
	JobApp                     *JobApp
	JobItemApp                 *JobItemApp
	JobItemAssignmentApp       *JobItemAssignmentApp
	JobItemWeighingApp         *JobItemWeighingApp
	MaterialApp                *MaterialApp
	OverIssueApp               *OverIssueApp
	ProcessApp                 *ProcessApp
	ScannedDataApp             *ScannedDataApp
	ShiftApp                   *ShiftApp
	ShiftScheduleApp           *ShiftScheduleApp
	StepApp                    *StepApp
	StepTypeApp                *StepTypeApp
	UnitOfMeasureApp           *UnitOfMeasureApp
	TerminalApp                *TerminalApp
	UnderIssueApp              *UnderIssueApp
	UnitOfMeasureConversionApp *UnitOfMeasureConversionApp
	UserApp                    *UserApp
	UserRoleApp                *UserRoleApp
	UserRoleAccessApp          *UserRoleAccessApp
	UserCompanyApp             *UserCompanyApp
	UserFactoryApp             *UserFactoryApp
	VesselApp                  *VesselApp
}

func NewAppStore(repoStore *persistance.RepoStore) *AppStore {
	appStore := AppStore{}
	appStore.AddressApp = NewAddressApp(repoStore.AddressRepo)
	appStore.AuthApp = NewAuthApp(repoStore.AuthRepo)
	appStore.BatchApp = NewBatchApp(repoStore.BatchRepo)
	appStore.BOMApp = NewBOMApp(repoStore.BOMRepo)
	appStore.BOMItemApp = NewBOMItemApp(repoStore.BOMItemRepo)
	appStore.CompanyApp = NewCompanyApp(repoStore.CompanyRepo)
	appStore.CommonApp = NewCommonApp(repoStore.CommonRepo)
	appStore.DeviceApp = NewDeviceApp(repoStore.DeviceRepo)
	appStore.DeviceDataApp = NewDeviceDataApp(repoStore.DeviceDataRepo)
	appStore.DeviceTypeApp = NewDeviceTypeApp(repoStore.DeviceTypeRepo)
	appStore.FactoryApp = NewFactoryRepository(repoStore.FactoryRepo)
	appStore.JobApp = NewJobApp(repoStore.JobRepo)
	appStore.JobItemApp = NewJobItemApp(repoStore.JobItemRepo)
	appStore.JobItemAssignmentApp = NewJobItemAssignmentApp(repoStore.JobItemAssignmentRepo)
	appStore.JobItemWeighingApp = NewJobItemWeighingApp(repoStore.JobItemWeighingRepo)
	appStore.MaterialApp = NewMaterialApp(repoStore.MaterialRepo)
	appStore.OverIssueApp = NewOverIssueApp(repoStore.OverIssueRepo)
	appStore.ProcessApp = NewProcessApp(repoStore.ProcessRepo)
	appStore.ScannedDataApp = NewScannedDataApp(repoStore.ScannedDataRepo)
	appStore.ShiftApp = NewShiftApp(repoStore.ShiftRepo)
	appStore.ShiftScheduleApp = NewShiftScheduleApp(repoStore.ShiftScheduleRepo)
	appStore.StepApp = NewStepApp(repoStore.StepRepo)
	appStore.StepTypeApp = NewStepTypeApp(repoStore.StepTypeRepo)
	appStore.TerminalApp = NewTerminalApp(repoStore.TerminalRepo)
	appStore.UnderIssueApp = NewUnderIssueApp(repoStore.UnderIssueRepo)
	appStore.UnitOfMeasureApp = NewUnitOfMeasureApp(repoStore.UnitOfMeasureRepo)
	appStore.UnitOfMeasureConversionApp = NewUnitOfMeasureConversionApp(repoStore.UnitOfMeasureConversionRepo)
	appStore.UserApp = NewUserApp(repoStore.UserRepo)
	appStore.UserRoleApp = NewUserRoleApp(repoStore.UserRoleRepo)
	appStore.UserRoleAccessApp = NewUserRoleAccessApp(repoStore.UserRoleAccessRepo)
	appStore.UserCompanyApp = NewUserCompanyApp(repoStore.UserCompanyRepo)
	appStore.UserFactoryApp = NewUserFactoryApp(repoStore.UserFactoryRepo)
	appStore.VesselApp = NewVesselApp(repoStore.VesselRepo)
	return &appStore
}
