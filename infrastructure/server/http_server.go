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
	}

	utilities.GinLogger() // Comment out for terminal logging
	httpServer.Router = gin.Default()
	httpServer.InterfaceStore = interfaceStore
	httpServer.MiddlewareStore = middlewareStore
	httpServer.AppStore = appStore
	return &httpServer
}

func (httpServer *HTTPServer) Serve() {
	// All Routers Here
}
