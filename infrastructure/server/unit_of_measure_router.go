package server

import (
	"eazyweigh/interfaces"
	"eazyweigh/interfaces/middlewares"

	"github.com/gin-gonic/gin"
)

type UOMRouter struct {
	router         *gin.RouterGroup
	interfaceStore *interfaces.InterfaceStore
	middlewares    *middlewares.MiddlewareStore
}

func NewUOMRouter(router *gin.RouterGroup, iStore *interfaces.InterfaceStore, middlewares *middlewares.MiddlewareStore) *UOMRouter {
	return &UOMRouter{
		router:         router,
		interfaceStore: iStore,
		middlewares:    middlewares,
	}
}

func (uomRouter *UOMRouter) ServeRoutes() {
	uomRouter.router.POST("/create/", uomRouter.middlewares.AuthMiddleware.ValidateAccessToken(), uomRouter.interfaceStore.UOMInterface.Create)
	uomRouter.router.POST("/create/multi/", uomRouter.middlewares.AuthMiddleware.ValidateAccessToken(), uomRouter.interfaceStore.UOMInterface.CreateMultiple)
	uomRouter.router.GET("/:id/", uomRouter.middlewares.AuthMiddleware.ValidateAccessToken(), uomRouter.interfaceStore.UOMInterface.Get)
	uomRouter.router.POST("/", uomRouter.middlewares.AuthMiddleware.ValidateAccessToken(), uomRouter.interfaceStore.UOMInterface.List)
	uomRouter.router.PATCH("/:id/", uomRouter.middlewares.AuthMiddleware.ValidateAccessToken(), uomRouter.interfaceStore.UOMInterface.Update)
}
