package server

import (
	"eazyweigh/interfaces"
	"eazyweigh/interfaces/middlewares"

	"github.com/gin-gonic/gin"
)

type ScannedDataRouter struct {
	router         *gin.RouterGroup
	interfaceStore *interfaces.InterfaceStore
	middlewares    *middlewares.MiddlewareStore
}

func NewScannedDataRouter(router *gin.RouterGroup, iStore *interfaces.InterfaceStore, middlewares *middlewares.MiddlewareStore) *ScannedDataRouter {
	return &ScannedDataRouter{
		router:         router,
		interfaceStore: iStore,
		middlewares:    middlewares,
	}
}

func (scannedDataRouter *ScannedDataRouter) ServeRoutes() {
	scannedDataRouter.router.POST("/create/", scannedDataRouter.middlewares.AuthMiddleware.ValidateAccessToken(), scannedDataRouter.interfaceStore.ScannedDataInterface.Create)
	scannedDataRouter.router.GET("/:id/", scannedDataRouter.middlewares.AuthMiddleware.ValidateAccessToken(), scannedDataRouter.interfaceStore.ScannedDataInterface.Get)
	scannedDataRouter.router.POST("/", scannedDataRouter.middlewares.AuthMiddleware.ValidateAccessToken(), scannedDataRouter.interfaceStore.ScannedDataInterface.List)
}
