package server

import (
	"eazyweigh/interfaces"
	"eazyweigh/interfaces/middlewares"

	"github.com/gin-gonic/gin"
)

type DeviceRouter struct {
	router         *gin.RouterGroup
	interfaceStore *interfaces.InterfaceStore
	middlewares    *middlewares.MiddlewareStore
}

func NewDeviceRouter(router *gin.RouterGroup, interfaceStore *interfaces.InterfaceStore, middlewares *middlewares.MiddlewareStore) *DeviceRouter {
	return &DeviceRouter{
		router:         router,
		interfaceStore: interfaceStore,
		middlewares:    middlewares,
	}
}

func (router *DeviceRouter) ServeRoutes() {
	router.router.POST("/create/", router.middlewares.AuthMiddleware.ValidateAccessToken(), router.middlewares.PermissionMiddleware.HasPermission(), router.interfaceStore.DeviceInterface.Create)
	router.router.POST("/create/multi/", router.middlewares.AuthMiddleware.ValidateAccessToken(), router.middlewares.PermissionMiddleware.HasPermission(), router.interfaceStore.DeviceInterface.CreateMultiple)
	router.router.GET("/:id/", router.middlewares.AuthMiddleware.ValidateAccessToken(), router.middlewares.PermissionMiddleware.HasPermission(), router.interfaceStore.DeviceInterface.Get)
	router.router.POST("/", router.middlewares.AuthMiddleware.ValidateAccessToken(), router.middlewares.PermissionMiddleware.HasPermission(), router.interfaceStore.DeviceInterface.List)
	router.router.PATCH("/:id/", router.middlewares.AuthMiddleware.ValidateAccessToken(), router.middlewares.PermissionMiddleware.HasPermission(), router.interfaceStore.DeviceInterface.Update)
}
