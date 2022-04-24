package server

import (
	"eazyweigh/interfaces"
	"eazyweigh/interfaces/middlewares"

	"github.com/gin-gonic/gin"
)

type ShiftScheduleRouter struct {
	router         *gin.RouterGroup
	interfaceStore *interfaces.InterfaceStore
	middlewares    *middlewares.MiddlewareStore
}

func NewShiftScheduleRouter(router *gin.RouterGroup, iStore *interfaces.InterfaceStore, middlewares *middlewares.MiddlewareStore) *ShiftScheduleRouter {
	return &ShiftScheduleRouter{
		router:         router,
		interfaceStore: iStore,
		middlewares:    middlewares,
	}
}

func (shiftScheduleRouter *ShiftScheduleRouter) ServeRoutes() {
	shiftScheduleRouter.router.POST("/create/", shiftScheduleRouter.middlewares.AuthMiddleware.ValidateAccessToken(), shiftScheduleRouter.interfaceStore.ShiftScheduleInterface.Create)
	shiftScheduleRouter.router.POST("/create/multi/", shiftScheduleRouter.middlewares.AuthMiddleware.ValidateAccessToken(), shiftScheduleRouter.interfaceStore.ShiftScheduleInterface.CreateMultiple)
	shiftScheduleRouter.router.DELETE("/:id/", shiftScheduleRouter.middlewares.AuthMiddleware.ValidateAccessToken(), shiftScheduleRouter.interfaceStore.ShiftScheduleInterface.Delete)
	shiftScheduleRouter.router.POST("/", shiftScheduleRouter.middlewares.AuthMiddleware.ValidateAccessToken(), shiftScheduleRouter.interfaceStore.ShiftScheduleInterface.List)
}
