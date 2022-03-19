package server

import (
	"eazyweigh/interfaces"
	"eazyweigh/interfaces/middlewares"

	"github.com/gin-gonic/gin"
)

type MaterialRouter struct {
	router         *gin.RouterGroup
	interfaceStore *interfaces.InterfaceStore
	middlewares    *middlewares.MiddlewareStore
}

func NewMaterialRouter(router *gin.RouterGroup, iStore *interfaces.InterfaceStore, middlewares *middlewares.MiddlewareStore) *MaterialRouter {
	return &MaterialRouter{
		router:         router,
		interfaceStore: iStore,
		middlewares:    middlewares,
	}
}

func (materialRouter *MaterialRouter) ServeRoutes() {
	materialRouter.router.POST("/create/", materialRouter.middlewares.AuthMiddleware.ValidateAccessToken(), materialRouter.interfaceStore.MaterialInterface.Create)
	materialRouter.router.POST("/create/multi/", materialRouter.middlewares.AuthMiddleware.ValidateAccessToken(), materialRouter.interfaceStore.MaterialInterface.CreateMultiple)
	materialRouter.router.GET("/:id/", materialRouter.middlewares.AuthMiddleware.ValidateAccessToken(), materialRouter.interfaceStore.MaterialInterface.Get)
	materialRouter.router.POST("/", materialRouter.middlewares.AuthMiddleware.ValidateAccessToken(), materialRouter.interfaceStore.MaterialInterface.List)
	materialRouter.router.PATCH("/:id/", materialRouter.middlewares.AuthMiddleware.ValidateAccessToken(), materialRouter.interfaceStore.MaterialInterface.Update)
}
