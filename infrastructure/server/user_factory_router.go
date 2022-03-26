package server

import (
	"eazyweigh/interfaces"
	"eazyweigh/interfaces/middlewares"

	"github.com/gin-gonic/gin"
)

type UserFactoryRouter struct {
	router         *gin.RouterGroup
	interfaceStore *interfaces.InterfaceStore
	middlewares    *middlewares.MiddlewareStore
}

func NewUserFactoryRouter(router *gin.RouterGroup, iStore *interfaces.InterfaceStore, middlewares *middlewares.MiddlewareStore) *UserFactoryRouter {
	return &UserFactoryRouter{
		router:         router,
		interfaceStore: iStore,
		middlewares:    middlewares,
	}
}

func (userFactoryRouter *UserFactoryRouter) ServeRoutes() {
	userFactoryRouter.router.POST("/create/", userFactoryRouter.middlewares.AuthMiddleware.ValidateAccessToken(), userFactoryRouter.interfaceStore.UserFactoryInterface.Create)
	userFactoryRouter.router.POST("/", userFactoryRouter.middlewares.AuthMiddleware.ValidateAccessToken(), userFactoryRouter.interfaceStore.UserFactoryInterface.Get)
}
