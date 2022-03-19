package server

import (
	"eazyweigh/interfaces"
	"eazyweigh/interfaces/middlewares"

	"github.com/gin-gonic/gin"
)

type JobItemRouter struct {
	router         *gin.RouterGroup
	interfaceStore *interfaces.InterfaceStore
	middlewares    *middlewares.MiddlewareStore
}

func NewJobItemRouter(router *gin.RouterGroup, iStore *interfaces.InterfaceStore, middlewares *middlewares.MiddlewareStore) *JobItemRouter {
	return &JobItemRouter{
		router:         router,
		interfaceStore: iStore,
		middlewares:    middlewares,
	}
}

func (jobItemRouter *JobItemRouter) ServeRoutes() {
	jobItemRouter.router.POST("/create/", jobItemRouter.middlewares.AuthMiddleware.ValidateAccessToken(), jobItemRouter.interfaceStore.JobItemInterface.Create)
	jobItemRouter.router.POST("/create/multi/", jobItemRouter.middlewares.AuthMiddleware.ValidateAccessToken(), jobItemRouter.interfaceStore.JobItemInterface.CreateMultiple)
	jobItemRouter.router.GET("/:id/", jobItemRouter.middlewares.AuthMiddleware.ValidateAccessToken(), jobItemRouter.interfaceStore.JobItemInterface.Get)
	jobItemRouter.router.PATCH("/:id/", jobItemRouter.middlewares.AuthMiddleware.ValidateAccessToken(), jobItemRouter.interfaceStore.JobItemInterface.Update)
}
