package interfaces

import (
	"eazyweigh/application"

	"github.com/hashicorp/go-hclog"
)

type InterfaceStore struct {
	logger                     hclog.Logger
	appStore                   *application.AppStore
	AddressInterface           *AddressInterface
	AuthInterface              *AuthInterface
	BOMInterface               *BOMInterface
	BOMItemInterface           *BOMItemInterface
	CompanyInterface           *CompanyInterface
	CommonInterface            *CommonInterface
	FactoryInterface           *FactoryInterface
	JobInterface               *JobInterface
	JobItemInterface           *JobItemInterface
	JobItemAssignmentInterface *JobItemAssignmentInterface
	JobItemWeighingInterface   *JobItemWeighingInterface
	MaterialInterface          *MaterialInterface
	OverIssueInterface         *OverIssueInterface
	ScannedDataInterface       *ScannedDataInterface
	ShiftInterface             *ShiftInterface
	ShiftScheduleInterface     *ShiftScheduleInterface
	TerminalInterface          *TerminalInterface
	UnderIssueInterface        *UnderIssueInterface
	UOMInterface               *UOMInterface
	UOMConversionInterface     *UOMConversionInterface
	UserInterface              *UserInterface
	UserRoleInterface          *UserRoleInterface
	UserRoleAccessInterface    *UserRoleAccessInterface
	UserCompanyInterface       *UserCompanyInterface
	UserFactoryInterface       *UserFactoryInterface
}

func NewInterfaceStore(logger hclog.Logger, appStore *application.AppStore) *InterfaceStore {
	interfaceStore := InterfaceStore{}
	interfaceStore.appStore = appStore
	interfaceStore.logger = logger
	interfaceStore.AddressInterface = NewAddressInterface(appStore, logger)
	interfaceStore.AuthInterface = NewAuthInterface(appStore, logger)
	interfaceStore.BOMInterface = NewBOMInterface(appStore, logger)
	interfaceStore.BOMItemInterface = NewBOMItemInterface(appStore, logger)
	interfaceStore.CompanyInterface = NewCompanyInterface(appStore, logger)
	interfaceStore.CommonInterface = NewCommonInterface(appStore, logger)
	interfaceStore.FactoryInterface = NewFactoryInterface(appStore, logger)
	interfaceStore.JobInterface = NewJobInterface(appStore, logger)
	interfaceStore.JobItemInterface = NewJobItemInterface(appStore, logger)
	interfaceStore.JobItemAssignmentInterface = NewJobItemAssignmentInterface(appStore, logger)
	interfaceStore.JobItemWeighingInterface = NewJobItemWeighingInterface(appStore, logger)
	interfaceStore.MaterialInterface = NewMaterialInterface(appStore, logger)
	interfaceStore.OverIssueInterface = NewOverIssueInterface(appStore, logger)
	interfaceStore.ScannedDataInterface = NewScannedDataInterface(appStore, logger)
	interfaceStore.ShiftInterface = NewShiftInterface(appStore, logger)
	interfaceStore.ShiftScheduleInterface = NewShiftScheduleInterface(appStore, logger)
	interfaceStore.TerminalInterface = NewTerminalInterface(appStore, logger)
	interfaceStore.UnderIssueInterface = NewUnderIssueInterface(appStore, logger)
	interfaceStore.UOMInterface = NewUOMInterface(appStore, logger)
	interfaceStore.UOMConversionInterface = NewUOMConversionInterface(appStore, logger)
	interfaceStore.UserInterface = NewUserInterface(appStore, logger)
	interfaceStore.UserRoleInterface = NewUserRoleInterface(appStore, logger)
	interfaceStore.UserRoleAccessInterface = NewUserRoleAccessInterface(appStore, logger)
	interfaceStore.UserCompanyInterface = NewUserCompanyInterface(appStore, logger)
	interfaceStore.UserFactoryInterface = NewUserFactoryInterface(appStore, logger)
	return &interfaceStore
}
