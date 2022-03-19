package server

import (
	"eazyweigh/interfaces"
	"eazyweigh/interfaces/middlewares"

	"github.com/gin-gonic/gin"
)

type UOMConversionRouter struct {
	router         *gin.RouterGroup
	interfaceStore *interfaces.InterfaceStore
	middlewares    *middlewares.MiddlewareStore
}

func NewUOMConversionRouter(router *gin.RouterGroup, iStore *interfaces.InterfaceStore, middlewares *middlewares.MiddlewareStore) *UOMConversionRouter {
	return &UOMConversionRouter{
		router:         router,
		interfaceStore: iStore,
		middlewares:    middlewares,
	}
}

func (uomConversionRouter *UOMConversionRouter) ServeRoutes() {
	uomConversionRouter.router.POST("/create/", uomConversionRouter.middlewares.AuthMiddleware.ValidateAccessToken(), uomConversionRouter.interfaceStore.UOMConversionInterface.Create)
	uomConversionRouter.router.POST("/create/multi/", uomConversionRouter.middlewares.AuthMiddleware.ValidateAccessToken(), uomConversionRouter.interfaceStore.UOMConversionInterface.CreateMultiple)
	uomConversionRouter.router.GET("/:id/", uomConversionRouter.middlewares.AuthMiddleware.ValidateAccessToken(), uomConversionRouter.interfaceStore.UOMConversionInterface.Get)
	uomConversionRouter.router.POST("/", uomConversionRouter.middlewares.AuthMiddleware.ValidateAccessToken(), uomConversionRouter.interfaceStore.UOMConversionInterface.List)
}
