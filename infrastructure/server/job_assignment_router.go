package server

import (
	"eazyweigh/interfaces"
	"eazyweigh/interfaces/middlewares"

	"github.com/gin-gonic/gin"
)

type JobAssignmentRouter struct {
	router         *gin.RouterGroup
	interfaceStore *interfaces.InterfaceStore
	middlewares    *middlewares.MiddlewareStore
}

func NewJobAssignmentRouter(router *gin.RouterGroup, iStore *interfaces.InterfaceStore, middlewares *middlewares.MiddlewareStore) *JobAssignmentRouter {
	return &JobAssignmentRouter{
		router:         router,
		interfaceStore: iStore,
		middlewares:    middlewares,
	}
}

func (jobAssignmentRouter *JobAssignmentRouter) ServeRoutes() {
	jobAssignmentRouter.router.POST("/create/", jobAssignmentRouter.middlewares.AuthMiddleware.ValidateAccessToken(), jobAssignmentRouter.interfaceStore.JobAssignmentInterface.Create)
	jobAssignmentRouter.router.POST("/create/multi/", jobAssignmentRouter.middlewares.AuthMiddleware.ValidateAccessToken(), jobAssignmentRouter.interfaceStore.JobAssignmentInterface.CreateMultiple)
	jobAssignmentRouter.router.GET("/:id/", jobAssignmentRouter.middlewares.AuthMiddleware.ValidateAccessToken(), jobAssignmentRouter.interfaceStore.JobAssignmentInterface.Get)
	jobAssignmentRouter.router.POST("/", jobAssignmentRouter.middlewares.AuthMiddleware.ValidateAccessToken(), jobAssignmentRouter.interfaceStore.JobAssignmentInterface.List)
	jobAssignmentRouter.router.PATCH("/:id/", jobAssignmentRouter.middlewares.AuthMiddleware.ValidateAccessToken(), jobAssignmentRouter.interfaceStore.JobAssignmentInterface.Update)
}
