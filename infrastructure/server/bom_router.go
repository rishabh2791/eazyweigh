package server

import (
	"eazyweigh/interfaces"
	"eazyweigh/interfaces/middlewares"

	"github.com/gin-gonic/gin"
)

type BOMRouter struct {
	router         *gin.RouterGroup
	interfaceStore *interfaces.InterfaceStore
	middlewares    *middlewares.MiddlewareStore
}

func NewBOMRouter(router *gin.RouterGroup, iStore *interfaces.InterfaceStore, middlewares *middlewares.MiddlewareStore) *BOMRouter {
	return &BOMRouter{
		router:         router,
		interfaceStore: iStore,
		middlewares:    middlewares,
	}
}

func (bomRouter *BOMRouter) ServeRoutes() {
	bomRouter.router.POST("/create/", bomRouter.middlewares.AuthMiddleware.ValidateAccessToken(), bomRouter.interfaceStore.BOMInterface.Create)
	bomRouter.router.POST("/create/multi/", bomRouter.middlewares.AuthMiddleware.ValidateAccessToken(), bomRouter.interfaceStore.BOMInterface.CreateMultiple)
	bomRouter.router.GET("/:id/", bomRouter.middlewares.AuthMiddleware.ValidateAccessToken(), bomRouter.interfaceStore.BOMInterface.Get)
	bomRouter.router.POST("/", bomRouter.middlewares.AuthMiddleware.ValidateAccessToken(), bomRouter.interfaceStore.BOMInterface.List)
	bomRouter.router.PATCH("/:id/", bomRouter.middlewares.AuthMiddleware.ValidateAccessToken(), bomRouter.interfaceStore.BOMInterface.Update)
}
