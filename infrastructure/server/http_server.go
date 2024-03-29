package server

import (
	"eazyweigh/application"
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/config"
	"eazyweigh/infrastructure/utilities"
	"eazyweigh/interfaces"
	"eazyweigh/interfaces/middlewares"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	if serverConfig.IsDebug() {
		httpServer.Router.Static("/public", "./public")
	} else {
		httpServer.Router.Static("/public", "/home/administrator/Development/code/backend/public")
	}
	httpServer.InterfaceStore = interfaceStore
	httpServer.MiddlewareStore = middlewareStore
	httpServer.AppStore = appStore
	return &httpServer
}

func (httpServer *HTTPServer) Serve() {
	httpServer.Router.POST("/image/upload/", ImageUploader)
	addressRouter := NewAddressRouter(httpServer.Router.Group("/address/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	authRouter := NewAuthRouter(httpServer.Router.Group("/auth/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	batchRouter := NewBatchRouter(httpServer.Router.Group("/batch/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	batchRunRouter := NewBatchRunRouter(httpServer.Router.Group("/batchrun/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	bomRouter := NewBOMRouter(httpServer.Router.Group("/bom/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	bomItemRouter := NewBOMItemRouter(httpServer.Router.Group("/bom_item/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	companyRouter := NewCompanyRouter(httpServer.Router.Group("/company/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	commonRouter := NewCommonRouter(httpServer.Router.Group("/common/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	deviceRouter := NewDeviceRouter(httpServer.Router.Group("/device/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	deviceDataRouter := NewDeviceDataRouter(httpServer.Router.Group("/device_data/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	deviceTypeRouter := NewDeviceTypeRouter(httpServer.Router.Group("/device_type/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	factoryRouter := NewFactoryRouter(httpServer.Router.Group("/factory/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	jobRouter := NewJobRouter(httpServer.Router.Group("/job/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	jobItemRouter := NewJobItemRouter(httpServer.Router.Group("/job_item/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	jobAssignmentRouter := NewJobItemWeightRouter(httpServer.Router.Group("/job_item_weight/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	jobItemAssignmentRouter := NewJobItemAssignmentRouter(httpServer.Router.Group("/job_item_assignment/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	materialRouter := NewMaterialRouter(httpServer.Router.Group("/material/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	overIssuerRouter := NewOverIssueRouter(httpServer.Router.Group("/over_issue/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	processRouter := NewProcessRouter(httpServer.Router.Group("/process/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	scannedDataRouter := NewScannedDataRouter(httpServer.Router.Group("/scan/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	shiftRouter := NewShiftRouter(httpServer.Router.Group("/shift/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	shiftScheduleRouter := NewShiftScheduleRouter(httpServer.Router.Group("/shift_schedule/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	underIssueRouter := NewUnderIssueRouter(httpServer.Router.Group("/under_issue/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	stepRouter := NewStepRouter(httpServer.Router.Group("/step/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	stepTypeRouter := NewStepTypeRouter(httpServer.Router.Group("/step_type/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	terminalRouter := NewTerminalRouter(httpServer.Router.Group("/terminal/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	uomRouter := NewUOMRouter(httpServer.Router.Group("/uom/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	uomConversionRouter := NewUOMConversionRouter(httpServer.Router.Group("/uom_conversion/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	userRouter := NewUserRouter(httpServer.Router.Group("/user/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	userRoleRouter := NewUserRoleRouter(httpServer.Router.Group("/user_role/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	userRoleAccessRouter := NewUserRoleAccessRouter(httpServer.Router.Group("/user_role_access/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	userCompanyRouter := NewUserCompanyRouter(httpServer.Router.Group("/user_company/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	userFactoryyRouter := NewUserFactoryRouter(httpServer.Router.Group("/user_factory/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)
	vesselRouter := NewVesselRouter(httpServer.Router.Group("/vessel/"), httpServer.InterfaceStore, httpServer.MiddlewareStore)

	addressRouter.ServeRoutes()
	authRouter.ServeRoutes()
	batchRouter.ServeRoutes()
	batchRunRouter.ServeRoutes()
	bomRouter.ServeRoutes()
	bomItemRouter.ServeRoutes()
	companyRouter.ServeRoutes()
	commonRouter.ServeRoutes()
	deviceRouter.ServeRoutes()
	deviceDataRouter.ServeRoutes()
	deviceTypeRouter.ServeRoutes()
	factoryRouter.ServeRoutes()
	jobRouter.ServeRoutes()
	jobItemRouter.ServeRoutes()
	jobAssignmentRouter.ServeRoutes()
	jobItemAssignmentRouter.ServeRoutes()
	materialRouter.ServeRoutes()
	overIssuerRouter.ServeRoutes()
	processRouter.ServeRoutes()
	scannedDataRouter.ServeRoutes()
	shiftRouter.ServeRoutes()
	shiftScheduleRouter.ServeRoutes()
	stepRouter.ServeRoutes()
	stepTypeRouter.ServeRoutes()
	terminalRouter.ServeRoutes()
	underIssueRouter.ServeRoutes()
	uomRouter.ServeRoutes()
	uomConversionRouter.ServeRoutes()
	userRouter.ServeRoutes()
	userRoleRouter.ServeRoutes()
	userRoleAccessRouter.ServeRoutes()
	userCompanyRouter.ServeRoutes()
	userFactoryyRouter.ServeRoutes()
	vesselRouter.ServeRoutes()
}

func ImageUploader(ctx *gin.Context) {
	response := value_objects.Response{}

	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		response.Status = false
		response.Message = err.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	extension := strings.Split(header.Filename, ".")[len(strings.Split(header.Filename, "."))-1]
	imageFile := "public/profile_pics/" + uuid.New().String() + "." + extension
	filename := "/home/administrator/Development/code/backend/public/profile_pics/" + uuid.New().String() + "." + extension
	out, err := os.Create(filename)
	if err != nil {
		response.Status = false
		response.Message = err.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		response.Status = false
		response.Message = err.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "Image Uploaded"
	response.Payload = imageFile

	ctx.AbortWithStatusJSON(http.StatusOK, response)
	return
}
