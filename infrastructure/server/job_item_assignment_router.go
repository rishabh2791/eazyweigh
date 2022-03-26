package server

import (
	"eazyweigh/interfaces"
	"eazyweigh/interfaces/middlewares"

	"github.com/gin-gonic/gin"
)

type JobItemAssignmentRouter struct {
	router         *gin.RouterGroup
	interfaceStore *interfaces.InterfaceStore
	middlewares    *middlewares.MiddlewareStore
}

func NewJobItemAssignmentRouter(router *gin.RouterGroup, iStore *interfaces.InterfaceStore, middlewares *middlewares.MiddlewareStore) *JobItemAssignmentRouter {
	return &JobItemAssignmentRouter{
		router:         router,
		interfaceStore: iStore,
		middlewares:    middlewares,
	}
}

func (jobItemAssignmentRouter *JobItemAssignmentRouter) ServeRoutes() {
	jobItemAssignmentRouter.router.POST("/create/", jobItemAssignmentRouter.middlewares.AuthMiddleware.ValidateAccessToken(), jobItemAssignmentRouter.interfaceStore.JobItemAssignmentInterface.Create)
	jobItemAssignmentRouter.router.POST("/create/multi/", jobItemAssignmentRouter.middlewares.AuthMiddleware.ValidateAccessToken(), jobItemAssignmentRouter.interfaceStore.JobItemAssignmentInterface.CreateMultiple)
	jobItemAssignmentRouter.router.GET("/:id/", jobItemAssignmentRouter.middlewares.AuthMiddleware.ValidateAccessToken(), jobItemAssignmentRouter.interfaceStore.JobItemAssignmentInterface.Get)
	jobItemAssignmentRouter.router.POST("/", jobItemAssignmentRouter.middlewares.AuthMiddleware.ValidateAccessToken(), jobItemAssignmentRouter.interfaceStore.JobItemAssignmentInterface.List)
	jobItemAssignmentRouter.router.PATCH("/:id/", jobItemAssignmentRouter.middlewares.AuthMiddleware.ValidateAccessToken(), jobItemAssignmentRouter.interfaceStore.JobItemAssignmentInterface.Update)
}
