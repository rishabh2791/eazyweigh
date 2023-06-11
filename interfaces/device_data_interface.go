package interfaces

import (
	"eazyweigh/application"
	"eazyweigh/domain/entity"
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-hclog"
)

type DeviceDataInterface struct {
	appStore *application.AppStore
	logger   hclog.Logger
}

func NewDeviceDataInterface(appStore *application.AppStore, logger hclog.Logger) *DeviceDataInterface {
	return &DeviceDataInterface{
		appStore: appStore,
		logger:   logger,
	}
}

func (deviceDataInterface *DeviceDataInterface) Create(ctx *gin.Context) {
	response := value_objects.Response{}

	model := entity.DeviceData{}
	jsonErr := json.NewDecoder(ctx.Request.Body).Decode(&model)
	if jsonErr != nil {
		response.Status = false
		response.Message = jsonErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	created, creationErr := deviceDataInterface.appStore.DeviceDataApp.Create(&model)
	if creationErr != nil {
		response.Status = false
		response.Message = creationErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "Device Data Created."
	response.Payload = created

	ctx.JSON(http.StatusOK, response)
}

func (deviceDataInterface *DeviceDataInterface) Get(ctx *gin.Context) {
	response := value_objects.Response{}
	id := ctx.Param("id")

	device, getErr := deviceDataInterface.appStore.DeviceDataApp.Get(id)
	if getErr != nil {
		response.Status = false
		response.Message = getErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "Device Data Found"
	response.Payload = device

	ctx.JSON(http.StatusOK, response)
}

func (deviceDataInterface *DeviceDataInterface) List(ctx *gin.Context) {
	response := value_objects.Response{}

	conditions := map[string]interface{}{}
	jsonError := json.NewDecoder(ctx.Request.Body).Decode(&conditions)
	if jsonError != nil {
		response.Status = false
		response.Message = jsonError.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	device, getErr := deviceDataInterface.appStore.DeviceDataApp.List(utilities.ConvertJSONToSQL(conditions))
	if getErr != nil {
		response.Status = false
		response.Message = getErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "Device Data Found"
	response.Payload = device

	ctx.JSON(http.StatusOK, response)
}
