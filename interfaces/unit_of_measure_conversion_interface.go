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

type UOMConversionInterface struct {
	appStore *application.AppStore
	logger   hclog.Logger
}

func NewUOMConversionInterface(appStore *application.AppStore, logger hclog.Logger) *UOMConversionInterface {
	return &UOMConversionInterface{
		appStore: appStore,
		logger:   logger,
	}
}

func (uomConversionInterface *UOMConversionInterface) Create(ctx *gin.Context) {
	response := value_objects.Response{}

	requestingUser, ok := ctx.Get("user")
	if !ok {
		response.Status = false
		response.Message = "Anonymous User"
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	user := requestingUser.(entity.User)

	model := entity.UnitOfMeasureConversion{}
	jsonErr := json.NewDecoder(ctx.Request.Body).Decode(&model)
	if jsonErr != nil {
		response.Status = false
		response.Message = jsonErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	model.CreatedByUsername = user.Username
	model.UpdatedByUsername = user.Username

	created, creationErr := uomConversionInterface.appStore.UnitOfMeasureConversionApp.Create(&model)
	if creationErr != nil {
		response.Status = false
		response.Message = creationErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "UOM Conversion Created."
	response.Payload = created

	ctx.JSON(http.StatusOK, response)
}

func (uomConversionInterface *UOMConversionInterface) CreateMultiple(ctx *gin.Context) {
	response := value_objects.Response{}
	createdModels := []interface{}{}
	creationErrors := []interface{}{}

	requestingUser, ok := ctx.Get("user")
	if !ok {
		response.Status = false
		response.Message = "Anonymous User"
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	user := requestingUser.(entity.User)

	models := []entity.UnitOfMeasureConversion{}
	jsonErr := json.NewDecoder(ctx.Request.Body).Decode(&models)
	if jsonErr != nil {
		response.Status = false
		response.Message = jsonErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	for _, model := range models {
		model.CreatedByUsername = user.Username
		model.UpdatedByUsername = user.Username

		created, creationErr := uomConversionInterface.appStore.UnitOfMeasureConversionApp.Create(&model)
		if creationErr != nil {
			creationErrors = append(creationErrors, creationErr)
		} else {
			createdModels = append(createdModels, created)
		}
	}

	response.Status = true
	response.Message = "UOM Conversions Created."
	response.Payload = map[string]interface{}{
		"models": createdModels,
		"errors": creationErrors,
	}

	ctx.JSON(http.StatusOK, response)
}

func (uomConversionInterface *UOMConversionInterface) Get(ctx *gin.Context) {
	response := value_objects.Response{}
	id := ctx.Param("id")

	uomConversion, getErr := uomConversionInterface.appStore.UnitOfMeasureConversionApp.Get(id)
	if getErr != nil {
		response.Status = false
		response.Message = getErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "UOM Conversion Found"
	response.Payload = uomConversion

	ctx.JSON(http.StatusOK, response)
}

func (uomConversionInterface *UOMConversionInterface) List(ctx *gin.Context) {
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

	uomConversions, getErr := uomConversionInterface.appStore.UnitOfMeasureConversionApp.List(utilities.ConvertJSONToSQL(conditions))
	if getErr != nil {
		response.Status = false
		response.Message = getErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "UOM Conversions Found"
	response.Payload = uomConversions

	ctx.JSON(http.StatusOK, response)
}
