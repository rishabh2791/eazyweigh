package middlewares

import (
	"eazyweigh/application"

	"github.com/hashicorp/go-hclog"
)

type MiddlewareStore struct {
	logger   hclog.Logger
	appStore *application.AppStore
}

func NewMiddlewareStore(logger hclog.Logger, appStore *application.AppStore) *MiddlewareStore {
	middlewareStore := MiddlewareStore{}
	middlewareStore.appStore = appStore
	middlewareStore.logger = logger
	return &middlewareStore
}
