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

type ShiftInterface struct {
	appStore *application.AppStore
	logger   hclog.Logger
}

func NewShiftInterface(appStore *application.AppStore, logger hclog.Logger) *ShiftInterface {
	return &ShiftInterface{
		appStore: appStore,
		logger:   logger,
	}
}

func (shiftInterface *ShiftInterface) Create(ctx *gin.Context) {
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

	model := entity.Shift{}
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

	created, creationErr := shiftInterface.appStore.ShiftApp.Create(&model)
	if creationErr != nil {
		response.Status = false
		response.Message = creationErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "Shift Created."
	response.Payload = created

	ctx.JSON(http.StatusOK, response)
}

func (shiftInterface *ShiftInterface) Get(ctx *gin.Context) {
	response := value_objects.Response{}
	id := ctx.Param("id")

	shift, getErr := shiftInterface.appStore.ShiftApp.Get(id)
	if getErr != nil {
		response.Status = false
		response.Message = getErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "Shift Found"
	response.Payload = shift

	ctx.JSON(http.StatusOK, response)
}

func (shiftInterface *ShiftInterface) List(ctx *gin.Context) {
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

	shifts, getErr := shiftInterface.appStore.ShiftApp.List(utilities.ConvertJSONToSQL(conditions))
	if getErr != nil {
		response.Status = false
		response.Message = getErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "Shifts Found"
	response.Payload = shifts

	ctx.JSON(http.StatusOK, response)
}
