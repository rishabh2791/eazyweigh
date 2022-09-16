package server

import (
	"eazyweigh/interfaces"
	"eazyweigh/interfaces/middlewares"

	"github.com/gin-gonic/gin"
)

type JobRouter struct {
	router         *gin.RouterGroup
	interfaceStore *interfaces.InterfaceStore
	middlewares    *middlewares.MiddlewareStore
}

func NewJobRouter(router *gin.RouterGroup, iStore *interfaces.InterfaceStore, middlewares *middlewares.MiddlewareStore) *JobRouter {
	return &JobRouter{
		router:         router,
		interfaceStore: iStore,
		middlewares:    middlewares,
	}
}

func (jobRouter *JobRouter) ServeRoutes() {
	jobRouter.router.POST("/create/", jobRouter.middlewares.AuthMiddleware.ValidateAccessToken(), jobRouter.interfaceStore.JobInterface.Create)
	jobRouter.router.POST("/create/multi/", jobRouter.middlewares.AuthMiddleware.ValidateAccessToken(), jobRouter.interfaceStore.JobInterface.CreateMultiple)
	jobRouter.router.GET("/:id/", jobRouter.middlewares.AuthMiddleware.ValidateAccessToken(), jobRouter.interfaceStore.JobInterface.Get)
	jobRouter.router.POST("/", jobRouter.middlewares.AuthMiddleware.ValidateAccessToken(), jobRouter.interfaceStore.JobInterface.List)
	jobRouter.router.PATCH("/:jobCode/", jobRouter.middlewares.AuthMiddleware.ValidateAccessToken(), jobRouter.interfaceStore.JobInterface.Update)
	jobRouter.router.GET("/remote/:factory_id/", jobRouter.middlewares.AuthMiddleware.ValidateAccessToken(), jobRouter.interfaceStore.JobInterface.PullFromRemote)
}
