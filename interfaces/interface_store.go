package interfaces

import (
	"eazyweigh/application"

	"github.com/hashicorp/go-hclog"
)

type InterfaceStore struct {
	logger   hclog.Logger
	appStore *application.AppStore
}

func NewInterfaceStore(logger hclog.Logger, appStore *application.AppStore) *InterfaceStore {
	interfaceStore := InterfaceStore{}
	interfaceStore.appStore = appStore
	interfaceStore.logger = logger
	return &interfaceStore
}
