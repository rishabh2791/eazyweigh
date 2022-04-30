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

type ScannedDataInterface struct {
	appStore *application.AppStore
	logger   hclog.Logger
}

func NewScannedDataInterface(appStore *application.AppStore, logger hclog.Logger) *ScannedDataInterface {
	return &ScannedDataInterface{
		appStore: appStore,
		logger:   logger,
	}
}

func (scannedDataInterface *ScannedDataInterface) Create(ctx *gin.Context) {
	response := value_objects.Response{}

	// requestingUser, ok := ctx.Get("user")
	// if !ok {
	// 	response.Status = false
	// 	response.Message = "Anonymous User"
	// 	response.Payload = ""

	// 	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
	// 	return
	// }
	// user := requestingUser.(*entity.User)

	model := entity.ScannedData{}
	jsonErr := json.NewDecoder(ctx.Request.Body).Decode(&model)
	if jsonErr != nil {
		response.Status = false
		response.Message = jsonErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	created, creationErr := scannedDataInterface.appStore.ScannedDataApp.Create(&model)
	if creationErr != nil {
		response.Status = false
		response.Message = creationErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "Scanned Data Created."
	response.Payload = created

	ctx.JSON(http.StatusOK, response)
}

func (scannedDataInterface *ScannedDataInterface) Get(ctx *gin.Context) {
	response := value_objects.Response{}
	id := ctx.Param("id")

	scannedData, getErr := scannedDataInterface.appStore.ScannedDataApp.Get(id)
	if getErr != nil {
		response.Status = false
		response.Message = getErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "Scanned Data Found"
	response.Payload = scannedData

	ctx.JSON(http.StatusOK, response)
}

func (scannedDataInterface *ScannedDataInterface) List(ctx *gin.Context) {
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

	scannedData, getErr := scannedDataInterface.appStore.ScannedDataApp.List(utilities.ConvertJSONToSQL(conditions))
	if getErr != nil {
		response.Status = false
		response.Message = getErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "Scanned Data Found"
	response.Payload = scannedData

	ctx.JSON(http.StatusOK, response)
}
