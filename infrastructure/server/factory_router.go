package server

import (
	"eazyweigh/interfaces"
	"eazyweigh/interfaces/middlewares"

	"github.com/gin-gonic/gin"
)

type FactoryRouter struct {
	router         *gin.RouterGroup
	interfaceStore *interfaces.InterfaceStore
	middlewares    *middlewares.MiddlewareStore
}

func NewFactoryRouter(router *gin.RouterGroup, iStore *interfaces.InterfaceStore, middlewares *middlewares.MiddlewareStore) *FactoryRouter {
	return &FactoryRouter{
		router:         router,
		interfaceStore: iStore,
		middlewares:    middlewares,
	}
}

func (factoryRouter *FactoryRouter) ServeRoutes() {
	factoryRouter.router.POST("/create/", factoryRouter.middlewares.AuthMiddleware.ValidateAccessToken(), factoryRouter.interfaceStore.FactoryInterface.Create)
	factoryRouter.router.POST("/create/multi/", factoryRouter.middlewares.AuthMiddleware.ValidateAccessToken(), factoryRouter.interfaceStore.FactoryInterface.CreateMultiple)
	factoryRouter.router.GET("/:id/", factoryRouter.middlewares.AuthMiddleware.ValidateAccessToken(), factoryRouter.interfaceStore.FactoryInterface.Get)
	factoryRouter.router.POST("/", factoryRouter.middlewares.AuthMiddleware.ValidateAccessToken(), factoryRouter.interfaceStore.FactoryInterface.List)
	factoryRouter.router.PATCH("/:id/", factoryRouter.middlewares.AuthMiddleware.ValidateAccessToken(), factoryRouter.interfaceStore.FactoryInterface.Update)
}
