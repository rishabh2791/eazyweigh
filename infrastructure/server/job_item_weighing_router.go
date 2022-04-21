package server

import (
	"eazyweigh/interfaces"
	"eazyweigh/interfaces/middlewares"

	"github.com/gin-gonic/gin"
)

type JobItemWeightRouter struct {
	router         *gin.RouterGroup
	interfaceStore *interfaces.InterfaceStore
	middlewares    *middlewares.MiddlewareStore
}

func NewJobItemWeightRouter(router *gin.RouterGroup, iStore *interfaces.InterfaceStore, middlewares *middlewares.MiddlewareStore) *JobItemWeightRouter {
	return &JobItemWeightRouter{
		router:         router,
		interfaceStore: iStore,
		middlewares:    middlewares,
	}
}

func (jobItemWeightRouter *JobItemWeightRouter) ServeRoutes() {
	jobItemWeightRouter.router.POST("/create/", jobItemWeightRouter.middlewares.AuthMiddleware.ValidateAccessToken(), jobItemWeightRouter.interfaceStore.JobItemWeighingInterface.Create)
	jobItemWeightRouter.router.GET("/list/:job_item_id/", jobItemWeightRouter.middlewares.AuthMiddleware.ValidateAccessToken(), jobItemWeightRouter.interfaceStore.JobItemWeighingInterface.List)
}
