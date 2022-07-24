package server

import (
	"eazyweigh/interfaces"
	"eazyweigh/interfaces/middlewares"

	"github.com/gin-gonic/gin"
)

type StepTypeRouter struct {
	router         *gin.RouterGroup
	interfaceStore *interfaces.InterfaceStore
	middlewares    *middlewares.MiddlewareStore
}

func NewStepTypeRouter(router *gin.RouterGroup, interfaceStore *interfaces.InterfaceStore, middlewares *middlewares.MiddlewareStore) *StepTypeRouter {
	return &StepTypeRouter{
		router:         router,
		interfaceStore: interfaceStore,
		middlewares:    middlewares,
	}
}

func (router *StepTypeRouter) ServeRoutes() {
	router.router.POST("/create/", router.middlewares.AuthMiddleware.ValidateAccessToken(), router.middlewares.PermissionMiddleware.HasPermission(), router.interfaceStore.StepTypeInterface.Create)
	router.router.POST("/create/multi/", router.middlewares.AuthMiddleware.ValidateAccessToken(), router.middlewares.PermissionMiddleware.HasPermission(), router.interfaceStore.StepTypeInterface.CreateMultiple)
	router.router.GET("/:id/", router.middlewares.AuthMiddleware.ValidateAccessToken(), router.middlewares.PermissionMiddleware.HasPermission(), router.interfaceStore.StepTypeInterface.Get)
	router.router.POST("/", router.middlewares.AuthMiddleware.ValidateAccessToken(), router.middlewares.PermissionMiddleware.HasPermission(), router.interfaceStore.StepTypeInterface.List)
	router.router.PATCH("/:id/", router.middlewares.AuthMiddleware.ValidateAccessToken(), router.middlewares.PermissionMiddleware.HasPermission(), router.interfaceStore.StepTypeInterface.Update)
}
