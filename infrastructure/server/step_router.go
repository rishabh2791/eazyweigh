package server

import (
	"eazyweigh/interfaces"
	"eazyweigh/interfaces/middlewares"

	"github.com/gin-gonic/gin"
)

type StepRouter struct {
	router         *gin.RouterGroup
	interfaceStore *interfaces.InterfaceStore
	middlewares    *middlewares.MiddlewareStore
}

func NewStepRouter(router *gin.RouterGroup, interfaceStore *interfaces.InterfaceStore, middlewares *middlewares.MiddlewareStore) *StepRouter {
	return &StepRouter{
		router:         router,
		interfaceStore: interfaceStore,
		middlewares:    middlewares,
	}
}

func (router *StepRouter) ServeRoutes() {
	router.router.POST("/create/", router.middlewares.AuthMiddleware.ValidateAccessToken(), router.middlewares.PermissionMiddleware.HasPermission(), router.interfaceStore.StepInterface.Create)
	router.router.POST("/create/multi/", router.middlewares.AuthMiddleware.ValidateAccessToken(), router.middlewares.PermissionMiddleware.HasPermission(), router.interfaceStore.StepInterface.CreateMultiple)
	router.router.GET("/:id/", router.middlewares.AuthMiddleware.ValidateAccessToken(), router.middlewares.PermissionMiddleware.HasPermission(), router.interfaceStore.StepInterface.Get)
	router.router.POST("/", router.middlewares.AuthMiddleware.ValidateAccessToken(), router.middlewares.PermissionMiddleware.HasPermission(), router.interfaceStore.StepInterface.List)
	router.router.PATCH("/:id/", router.middlewares.AuthMiddleware.ValidateAccessToken(), router.middlewares.PermissionMiddleware.HasPermission(), router.interfaceStore.StepInterface.Update)
}
