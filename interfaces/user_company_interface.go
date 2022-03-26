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

type UserCompanyInterface struct {
	appStore *application.AppStore
	logger   hclog.Logger
}

func NewUserCompanyInterface(appStore *application.AppStore, logger hclog.Logger) *UserCompanyInterface {
	return &UserCompanyInterface{
		appStore: appStore,
		logger:   logger,
	}
}

func (userCompanyInterface *UserCompanyInterface) Create(ctx *gin.Context) {
	response := value_objects.Response{}

	model := entity.UserCompany{}
	jsonErr := json.NewDecoder(ctx.Request.Body).Decode(&model)
	if jsonErr != nil {
		response.Status = false
		response.Message = jsonErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	created, creationErr := userCompanyInterface.appStore.UserCompanyApp.Create(&model)
	if creationErr != nil {
		response.Status = false
		response.Message = creationErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "User Company Created."
	response.Payload = created

	ctx.JSON(http.StatusOK, response)
}

func (userCompanyInterface *UserCompanyInterface) Get(ctx *gin.Context) {
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

	uoms, getErr := userCompanyInterface.appStore.UserCompanyApp.Get(utilities.ConvertJSONToSQL(conditions))
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
