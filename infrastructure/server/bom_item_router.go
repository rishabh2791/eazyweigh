package server

import (
	"eazyweigh/interfaces"
	"eazyweigh/interfaces/middlewares"

	"github.com/gin-gonic/gin"
)

type BOMItemRouter struct {
	router         *gin.RouterGroup
	interfaceStore *interfaces.InterfaceStore
	middlewares    *middlewares.MiddlewareStore
}

func NewBOMItemRouter(router *gin.RouterGroup, iStore *interfaces.InterfaceStore, middlewares *middlewares.MiddlewareStore) *BOMItemRouter {
	return &BOMItemRouter{
		router:         router,
		interfaceStore: iStore,
		middlewares:    middlewares,
	}
}

func (bomItemRouter *BOMItemRouter) ServeRoutes() {
	bomItemRouter.router.POST("/create/", bomItemRouter.middlewares.AuthMiddleware.ValidateAccessToken(), bomItemRouter.interfaceStore.BOMItemInterface.Create)
	bomItemRouter.router.POST("/create/multi/", bomItemRouter.middlewares.AuthMiddleware.ValidateAccessToken(), bomItemRouter.interfaceStore.BOMItemInterface.CreateMultiple)
	bomItemRouter.router.GET("/:id/", bomItemRouter.middlewares.AuthMiddleware.ValidateAccessToken(), bomItemRouter.interfaceStore.BOMItemInterface.Get)
	bomItemRouter.router.POST("/", bomItemRouter.middlewares.AuthMiddleware.ValidateAccessToken(), bomItemRouter.interfaceStore.BOMItemInterface.List)
	bomItemRouter.router.PATCH("/:id/", bomItemRouter.middlewares.AuthMiddleware.ValidateAccessToken(), bomItemRouter.interfaceStore.BOMItemInterface.Update)
}
