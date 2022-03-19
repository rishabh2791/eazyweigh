package server

import (
	"eazyweigh/interfaces"
	"eazyweigh/interfaces/middlewares"

	"github.com/gin-gonic/gin"
)

type AddressRouter struct {
	router         *gin.RouterGroup
	interfaceStore *interfaces.InterfaceStore
	middlewares    *middlewares.MiddlewareStore
}

func NewAddressRouter(router *gin.RouterGroup, iStore *interfaces.InterfaceStore, middlewares *middlewares.MiddlewareStore) *AddressRouter {
	return &AddressRouter{
		router:         router,
		interfaceStore: iStore,
		middlewares:    middlewares,
	}
}

func (addressRouter *AddressRouter) ServeRoutes() {
	addressRouter.router.POST("/create/", addressRouter.middlewares.AuthMiddleware.ValidateAccessToken(), addressRouter.interfaceStore.AddressInterface.Create)
	addressRouter.router.POST("/", addressRouter.middlewares.AuthMiddleware.ValidateAccessToken(), addressRouter.interfaceStore.AddressInterface.List)
	addressRouter.router.PATCH("/:id/", addressRouter.middlewares.AuthMiddleware.ValidateAccessToken(), addressRouter.interfaceStore.AddressInterface.Update)
	addressRouter.router.DELETE("/:id/", addressRouter.middlewares.AuthMiddleware.ValidateAccessToken(), addressRouter.interfaceStore.AddressInterface.Delete)
}
