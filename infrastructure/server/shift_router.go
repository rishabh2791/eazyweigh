package server

import (
	"eazyweigh/interfaces"
	"eazyweigh/interfaces/middlewares"

	"github.com/gin-gonic/gin"
)

type ShiftRouter struct {
	router         *gin.RouterGroup
	interfaceStore *interfaces.InterfaceStore
	middlewares    *middlewares.MiddlewareStore
}

func NewShiftRouter(router *gin.RouterGroup, iStore *interfaces.InterfaceStore, middlewares *middlewares.MiddlewareStore) *ShiftRouter {
	return &ShiftRouter{
		router:         router,
		interfaceStore: iStore,
		middlewares:    middlewares,
	}
}

func (shiftRouter *ShiftRouter) ServeRoutes() {
	shiftRouter.router.POST("/create/", shiftRouter.middlewares.AuthMiddleware.ValidateAccessToken(), shiftRouter.interfaceStore.ShiftInterface.Create)
	shiftRouter.router.GET("/:id/", shiftRouter.middlewares.AuthMiddleware.ValidateAccessToken(), shiftRouter.interfaceStore.ShiftInterface.Get)
	shiftRouter.router.POST("/", shiftRouter.middlewares.AuthMiddleware.ValidateAccessToken(), shiftRouter.interfaceStore.ShiftInterface.List)
}
