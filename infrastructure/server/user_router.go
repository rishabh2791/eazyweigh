package server

import (
	"eazyweigh/interfaces"
	"eazyweigh/interfaces/middlewares"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	router         *gin.RouterGroup
	interfaceStore *interfaces.InterfaceStore
	middlewares    *middlewares.MiddlewareStore
}

func NewUserRouter(router *gin.RouterGroup, interfaceStore *interfaces.InterfaceStore, middlewares *middlewares.MiddlewareStore) *UserRouter {
	return &UserRouter{
		router:         router,
		interfaceStore: interfaceStore,
		middlewares:    middlewares,
	}
}

func (router *UserRouter) ServeRoutes() {
	router.router.POST("/create/", router.middlewares.AuthMiddleware.ValidateAccessToken(), router.middlewares.PermissionMiddleware.HasPermission(), router.interfaceStore.UserInterface.Create)
	router.router.POST("/create/multi/", router.middlewares.AuthMiddleware.ValidateAccessToken(), router.middlewares.PermissionMiddleware.HasPermission(), router.interfaceStore.UserInterface.CreateMultiple)
	router.router.GET("/:username/", router.middlewares.AuthMiddleware.ValidateAccessToken(), router.middlewares.PermissionMiddleware.HasPermission(), router.interfaceStore.UserInterface.Get)
	router.router.POST("/", router.middlewares.AuthMiddleware.ValidateAccessToken(), router.middlewares.PermissionMiddleware.HasPermission(), router.interfaceStore.UserInterface.List)
	router.router.PATCH("/:username/", router.middlewares.AuthMiddleware.ValidateAccessToken(), router.middlewares.PermissionMiddleware.HasPermission(), router.interfaceStore.UserInterface.Update)
}
