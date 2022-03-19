package server

import (
	"eazyweigh/application"
	"eazyweigh/infrastructure/config"
	"eazyweigh/infrastructure/utilities"
	"eazyweigh/interfaces"
	"eazyweigh/interfaces/middlewares"

	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	Router          *gin.Engine
	MiddlewareStore *middlewares.MiddlewareStore
	InterfaceStore  *interfaces.InterfaceStore
	AppStore        *application.AppStore
}

func NewHTTPServer(serverConfig config.ServerConfig, appStore *application.AppStore, interfaceStore *interfaces.InterfaceStore, middlewareStore *middlewares.MiddlewareStore) *HTTPServer {
	httpServer := HTTPServer{}

	if !serverConfig.IsDebug() {
		gin.SetMode(gin.ReleaseMode)
		utilities.GinLogger()
	} else {
		utilities.NewConsoleLogger()
	}

	httpServer.Router = gin.Default()
	httpServer.InterfaceStore = interfaceStore
	httpServer.MiddlewareStore = middlewareStore
	httpServer.AppStore = appStore
	return &httpServer
}

func (httpServer *HTTPServer) Serve() {
	addressRouter := NewAddressRouter(httpServer.Router.Group("/address/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	authRouter := NewAuthRouter(httpServer.Router.Group("/auth/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	bomRouter := NewAddressRouter(httpServer.Router.Group("/bom/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	bomItemRouter := NewAddressRouter(httpServer.Router.Group("/bom_item/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	companyRouter := NewAddressRouter(httpServer.Router.Group("/company/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	factoryRouter := NewAddressRouter(httpServer.Router.Group("/factory/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	jobRouter := NewAddressRouter(httpServer.Router.Group("/job/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	jobItemRouter := NewAddressRouter(httpServer.Router.Group("/job_item/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	jobAssignmentRouter := NewAddressRouter(httpServer.Router.Group("/job_assignment/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	materialRouter := NewAddressRouter(httpServer.Router.Group("/material/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	overIssuerRouter := NewOverIssueRouter(httpServer.Router.Group("/over_issue/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	scannedDataRouter := NewScannedDataRouter(httpServer.Router.Group("/scan/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	shiftRouter := NewAddressRouter(httpServer.Router.Group("/shift/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	shiftScheduleRouter := NewAddressRouter(httpServer.Router.Group("/shift_schedule/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	underIssueRouter := NewUnderIssueRouter(httpServer.Router.Group("/under_issue/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	terminalRouter := NewAddressRouter(httpServer.Router.Group("/terminal/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	uomRouter := NewAddressRouter(httpServer.Router.Group("/uom/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	uomConversionRouter := NewAddressRouter(httpServer.Router.Group("/uom_conversion/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	userRouter := NewUserRouter(httpServer.Router.Group("/user/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	userRoleRouter := NewAddressRouter(httpServer.Router.Group("/user_role/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)

	addressRouter.ServeRoutes()
	authRouter.ServeRoutes()
	bomRouter.ServeRoutes()
	bomItemRouter.ServeRoutes()
	companyRouter.ServeRoutes()
	factoryRouter.ServeRoutes()
	jobRouter.ServeRoutes()
	jobItemRouter.ServeRoutes()
	jobAssignmentRouter.ServeRoutes()
	materialRouter.ServeRoutes()
	overIssuerRouter.ServeRoutes()
	scannedDataRouter.ServeRoutes()
	shiftRouter.ServeRoutes()
	shiftScheduleRouter.ServeRoutes()
	terminalRouter.ServeRoutes()
	underIssueRouter.ServeRoutes()
	uomRouter.ServeRoutes()
	uomConversionRouter.ServeRoutes()
	userRouter.ServeRoutes()
	userRoleRouter.ServeRoutes()
}
