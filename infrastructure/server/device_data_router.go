package server

import (
	"eazyweigh/interfaces"
	"eazyweigh/interfaces/middlewares"

	"github.com/gin-gonic/gin"
)

type DeviceDataRouter struct {
	router         *gin.RouterGroup
	interfaceStore *interfaces.InterfaceStore
	middlewares    *middlewares.MiddlewareStore
}

func NewDeviceDataRouter(router *gin.RouterGroup, interfaceStore *interfaces.InterfaceStore, middlewares *middlewares.MiddlewareStore) *DeviceDataRouter {
	return &DeviceDataRouter{
		router:         router,
		interfaceStore: interfaceStore,
		middlewares:    middlewares,
	}
}

func (router *DeviceDataRouter) ServeRoutes() {
	router.router.POST("/create/", router.middlewares.AuthMiddleware.ValidateAccessToken(), router.middlewares.PermissionMiddleware.HasPermission(), router.interfaceStore.DeviceDataInterface.Create)
	router.router.GET("/:deviceID/", router.middlewares.AuthMiddleware.ValidateAccessToken(), router.middlewares.PermissionMiddleware.HasPermission(), router.interfaceStore.DeviceDataInterface.List)
}
