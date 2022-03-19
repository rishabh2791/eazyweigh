package server

import (
	"eazyweigh/interfaces"
	"eazyweigh/interfaces/middlewares"

	"github.com/gin-gonic/gin"
)

type CompanyRouter struct {
	router         *gin.RouterGroup
	interfaceStore *interfaces.InterfaceStore
	middlewares    *middlewares.MiddlewareStore
}

func NewCompanyRouter(router *gin.RouterGroup, iStore *interfaces.InterfaceStore, middlewares *middlewares.MiddlewareStore) *CompanyRouter {
	return &CompanyRouter{
		router:         router,
		interfaceStore: iStore,
		middlewares:    middlewares,
	}
}

func (companyRouter *CompanyRouter) ServeRoutes() {
	companyRouter.router.POST("/create/", companyRouter.middlewares.AuthMiddleware.ValidateAccessToken(), companyRouter.interfaceStore.CompanyInterface.Create)
	companyRouter.router.GET("/:id/", companyRouter.middlewares.AuthMiddleware.ValidateAccessToken(), companyRouter.interfaceStore.CompanyInterface.Get)
	companyRouter.router.POST("/", companyRouter.middlewares.AuthMiddleware.ValidateAccessToken(), companyRouter.interfaceStore.CompanyInterface.List)
	companyRouter.router.PATCH("/:id/", companyRouter.middlewares.AuthMiddleware.ValidateAccessToken(), companyRouter.interfaceStore.CompanyInterface.Update)
}
