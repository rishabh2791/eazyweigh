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
	bomRouter := NewBOMRouter(httpServer.Router.Group("/bom/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	bomItemRouter := NewBOMItemRouter(httpServer.Router.Group("/bom_item/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	companyRouter := NewCompanyRouter(httpServer.Router.Group("/company/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	factoryRouter := NewFactoryRouter(httpServer.Router.Group("/factory/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	jobRouter := NewJobRouter(httpServer.Router.Group("/job/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	jobItemRouter := NewJobItemRouter(httpServer.Router.Group("/job_item/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	jobAssignmentRouter := NewJobAssignmentRouter(httpServer.Router.Group("/job_assignment/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	jobItemAssignmentRouter := NewJobItemAssignmentRouter(httpServer.Router.Group("/job_item_assignment/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	materialRouter := NewMaterialRouter(httpServer.Router.Group("/material/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	overIssuerRouter := NewOverIssueRouter(httpServer.Router.Group("/over_issue/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	scannedDataRouter := NewScannedDataRouter(httpServer.Router.Group("/scan/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	shiftRouter := NewShiftRouter(httpServer.Router.Group("/shift/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	shiftScheduleRouter := NewShiftScheduleRouter(httpServer.Router.Group("/shift_schedule/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	underIssueRouter := NewUnderIssueRouter(httpServer.Router.Group("/under_issue/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	terminalRouter := NewTerminalRouter(httpServer.Router.Group("/terminal/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	uomRouter := NewUOMRouter(httpServer.Router.Group("/uom/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	uomConversionRouter := NewUOMConversionRouter(httpServer.Router.Group("/uom_conversion/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	userRouter := NewUserRouter(httpServer.Router.Group("/user/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	userRoleRouter := NewUserRoleRouter(httpServer.Router.Group("/user_role/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	userCompanyRouter := NewUserCompanyRouter(httpServer.Router.Group("/user_company/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	userFactoryyRouter := NewUserFactoryRouter(httpServer.Router.Group("/user_factory/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)

	addressRouter.ServeRoutes()
	authRouter.ServeRoutes()
	bomRouter.ServeRoutes()
	bomItemRouter.ServeRoutes()
	companyRouter.ServeRoutes()
	factoryRouter.ServeRoutes()
	jobRouter.ServeRoutes()
	jobItemRouter.ServeRoutes()
	jobAssignmentRouter.ServeRoutes()
	jobItemAssignmentRouter.ServeRoutes()
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
	userCompanyRouter.ServeRoutes()
	userFactoryyRouter.ServeRoutes()
}
