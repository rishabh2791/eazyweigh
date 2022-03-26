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

type UserFactoryInterface struct {
	appStore *application.AppStore
	logger   hclog.Logger
}

func NewUserFactoryInterface(appStore *application.AppStore, logger hclog.Logger) *UserFactoryInterface {
	return &UserFactoryInterface{
		appStore: appStore,
		logger:   logger,
	}
}

func (userFactoryInterface *UserFactoryInterface) Create(ctx *gin.Context) {
	response := value_objects.Response{}

	model := entity.UserFactory{}
	jsonErr := json.NewDecoder(ctx.Request.Body).Decode(&model)
	if jsonErr != nil {
		response.Status = false
		response.Message = jsonErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	created, creationErr := userFactoryInterface.appStore.UserFactoryApp.Create(&model)
	if creationErr != nil {
		response.Status = false
		response.Message = creationErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "User Factory Created."
	response.Payload = created

	ctx.JSON(http.StatusOK, response)
}

func (userFactoryInterface *UserFactoryInterface) Get(ctx *gin.Context) {
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

	uoms, getErr := userFactoryInterface.appStore.UserFactoryApp.Get(utilities.ConvertJSONToSQL(conditions))
	if getErr != nil {
		response.Status = false
		response.Message = getErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "User Factories Found"
	response.Payload = uoms

	ctx.JSON(http.StatusOK, response)
}
