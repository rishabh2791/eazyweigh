package interfaces

import (
	"eazyweigh/application"
	"eazyweigh/domain/entity"
	"eazyweigh/domain/value_objects"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-hclog"
)

type JobItemWeighingInterface struct {
	appStore *application.AppStore
	logger   hclog.Logger
}

func NewJobItemWeighingInterface(appStore *application.AppStore, logger hclog.Logger) *JobItemWeighingInterface {
	return &JobItemWeighingInterface{
		appStore: appStore,
		logger:   logger,
	}
}

func (jobItemWeighingInterface *JobItemWeighingInterface) Create(ctx *gin.Context) {
	response := value_objects.Response{}

	requestingUser, ok := ctx.Get("user")
	if !ok {
		response.Status = false
		response.Message = "Anonymous User"
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	user := requestingUser.(*entity.User)

	model := entity.JobItemWeighing{}
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

	created, creationErr := jobItemWeighingInterface.appStore.JobItemWeighingApp.Create(&model)
	if creationErr != nil {
		response.Status = false
		response.Message = creationErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "Job Items Weight Created."
	response.Payload = created

	ctx.JSON(http.StatusOK, response)
}
