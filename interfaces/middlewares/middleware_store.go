package middlewares

import (
	"eazyweigh/application"

	"github.com/hashicorp/go-hclog"
)

type MiddlewareStore struct {
	logger               hclog.Logger
	appStore             *application.AppStore
	AuthMiddleware       *AuthMiddleware
	PermissionMiddleware *PermissionMiddleware
}

func NewMiddlewareStore(logger hclog.Logger, appStore *application.AppStore) *MiddlewareStore {
	middlewareStore := MiddlewareStore{}
	middlewareStore.appStore = appStore
	middlewareStore.logger = logger
	middlewareStore.AuthMiddleware = NewAuthMiddleware(logger, *appStore)
	middlewareStore.PermissionMiddleware = NewPermissionMiddleware(logger, *appStore)
	return &middlewareStore
}
