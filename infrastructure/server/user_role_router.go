package server

import (
	"eazyweigh/interfaces"
	"eazyweigh/interfaces/middlewares"

	"github.com/gin-gonic/gin"
)

type UserRoleRouter struct {
	router         *gin.RouterGroup
	interfaceStore *interfaces.InterfaceStore
	middlewares    *middlewares.MiddlewareStore
}

func NewUserRoleRouter(router *gin.RouterGroup, iStore *interfaces.InterfaceStore, middlewares *middlewares.MiddlewareStore) *UserRoleRouter {
	return &UserRoleRouter{
		router:         router,
		interfaceStore: iStore,
		middlewares:    middlewares,
	}
}

func (userRoleRouter *UserRoleRouter) ServeRoutes() {
	userRoleRouter.router.POST("/create/", userRoleRouter.middlewares.AuthMiddleware.ValidateAccessToken(), userRoleRouter.interfaceStore.UserRoleInterface.Create)
	userRoleRouter.router.POST("/create/multi/", userRoleRouter.middlewares.AuthMiddleware.ValidateAccessToken(), userRoleRouter.interfaceStore.UserRoleInterface.CreateMultiple)
	userRoleRouter.router.GET("/:id/", userRoleRouter.middlewares.AuthMiddleware.ValidateAccessToken(), userRoleRouter.interfaceStore.UserRoleInterface.Get)
	userRoleRouter.router.POST("/", userRoleRouter.middlewares.AuthMiddleware.ValidateAccessToken(), userRoleRouter.interfaceStore.UserRoleInterface.List)
	userRoleRouter.router.PATCH("/:id/", userRoleRouter.middlewares.AuthMiddleware.ValidateAccessToken(), userRoleRouter.interfaceStore.UserRoleInterface.Update)
}
