package server

import (
	"eazyweigh/interfaces"
	"eazyweigh/interfaces/middlewares"

	"github.com/gin-gonic/gin"
)

type UnderIssueRouter struct {
	router         *gin.RouterGroup
	interfaceStore *interfaces.InterfaceStore
	middlewares    *middlewares.MiddlewareStore
}

func NewUnderIssueRouter(router *gin.RouterGroup, iStore *interfaces.InterfaceStore, middlewares *middlewares.MiddlewareStore) *UnderIssueRouter {
	return &UnderIssueRouter{
		router:         router,
		interfaceStore: iStore,
		middlewares:    middlewares,
	}
}

func (underIssueRouter *UnderIssueRouter) ServeRoutes() {
	underIssueRouter.router.POST("/create/", underIssueRouter.middlewares.AuthMiddleware.ValidateAccessToken(), underIssueRouter.interfaceStore.UnderIssueInterface.Create)
	underIssueRouter.router.POST("/create/multi/", underIssueRouter.middlewares.AuthMiddleware.ValidateAccessToken(), underIssueRouter.interfaceStore.UnderIssueInterface.CreateMultiple)
	underIssueRouter.router.POST("/", underIssueRouter.middlewares.AuthMiddleware.ValidateAccessToken(), underIssueRouter.interfaceStore.UnderIssueInterface.List)
	underIssueRouter.router.PATCH("/:id/", underIssueRouter.middlewares.AuthMiddleware.ValidateAccessToken(), underIssueRouter.interfaceStore.UnderIssueInterface.Update)
}
