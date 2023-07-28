package server

import (
	"eazyweigh/interfaces"
	"eazyweigh/interfaces/middlewares"

	"github.com/gin-gonic/gin"
)

type OverIssueRouter struct {
	router         *gin.RouterGroup
	interfaceStore *interfaces.InterfaceStore
	middlewares    *middlewares.MiddlewareStore
}

func NewOverIssueRouter(router *gin.RouterGroup, iStore *interfaces.InterfaceStore, middlewares *middlewares.MiddlewareStore) *OverIssueRouter {
	return &OverIssueRouter{
		router:         router,
		interfaceStore: iStore,
		middlewares:    middlewares,
	}
}

func (overIssueRouter *OverIssueRouter) ServeRoutes() {
	overIssueRouter.router.POST("/create/", overIssueRouter.middlewares.AuthMiddleware.ValidateAccessToken(), overIssueRouter.interfaceStore.OverIssueInterface.Create)
	overIssueRouter.router.POST("/create/multi/", overIssueRouter.middlewares.AuthMiddleware.ValidateAccessToken(), overIssueRouter.interfaceStore.OverIssueInterface.CreateMultiple)
	overIssueRouter.router.POST("/", overIssueRouter.middlewares.AuthMiddleware.ValidateAccessToken(), overIssueRouter.interfaceStore.OverIssueInterface.List)
	overIssueRouter.router.PATCH("/:id/", overIssueRouter.middlewares.AuthMiddleware.ValidateAccessToken(), overIssueRouter.interfaceStore.OverIssueInterface.Update)
}
