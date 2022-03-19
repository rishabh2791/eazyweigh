package server

import (
	"eazyweigh/interfaces"
	"eazyweigh/interfaces/middlewares"

	"github.com/gin-gonic/gin"
)

type TerminalRouter struct {
	router         *gin.RouterGroup
	interfaceStore *interfaces.InterfaceStore
	middlewares    *middlewares.MiddlewareStore
}

func NewTerminalRouter(router *gin.RouterGroup, iStore *interfaces.InterfaceStore, middlewares *middlewares.MiddlewareStore) *TerminalRouter {
	return &TerminalRouter{
		router:         router,
		interfaceStore: iStore,
		middlewares:    middlewares,
	}
}

func (terminalRouter *TerminalRouter) ServeRoutes() {
	terminalRouter.router.POST("/create/", terminalRouter.middlewares.AuthMiddleware.ValidateAccessToken(), terminalRouter.interfaceStore.TerminalInterface.Create)
	terminalRouter.router.GET("/:id/", terminalRouter.middlewares.AuthMiddleware.ValidateAccessToken(), terminalRouter.interfaceStore.TerminalInterface.Get)
	terminalRouter.router.POST("/", terminalRouter.middlewares.AuthMiddleware.ValidateAccessToken(), terminalRouter.interfaceStore.TerminalInterface.List)
	terminalRouter.router.PATCH("/:id/", terminalRouter.middlewares.AuthMiddleware.ValidateAccessToken(), terminalRouter.interfaceStore.TerminalInterface.Update)
}
