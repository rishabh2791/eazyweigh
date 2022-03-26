package server

import (
	"eazyweigh/interfaces"
	"eazyweigh/interfaces/middlewares"

	"github.com/gin-gonic/gin"
)

type UserCompanyRouter struct {
	router         *gin.RouterGroup
	interfaceStore *interfaces.InterfaceStore
	middlewares    *middlewares.MiddlewareStore
}

func NewUserCompanyRouter(router *gin.RouterGroup, iStore *interfaces.InterfaceStore, middlewares *middlewares.MiddlewareStore) *UserCompanyRouter {
	return &UserCompanyRouter{
		router:         router,
		interfaceStore: iStore,
		middlewares:    middlewares,
	}
}

func (userCompanyRouter *UserCompanyRouter) ServeRoutes() {
	userCompanyRouter.router.POST("/create/", userCompanyRouter.middlewares.AuthMiddleware.ValidateAccessToken(), userCompanyRouter.interfaceStore.UserCompanyInterface.Create)
	userCompanyRouter.router.POST("/", userCompanyRouter.middlewares.AuthMiddleware.ValidateAccessToken(), userCompanyRouter.interfaceStore.UserCompanyInterface.Get)
}
