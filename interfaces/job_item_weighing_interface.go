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

func (jobItemWeighingInterface *JobItemWeighingInterface) List(ctx *gin.Context) {
	response := value_objects.Response{}
	id := ctx.Param("job_item_id")

	jobItemWeighings, getErr := jobItemWeighingInterface.appStore.JobItemWeighingApp.List(id)
	if getErr != nil {
		response.Status = false
		response.Message = getErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "Job Items Found"
	response.Payload = jobItemWeighings

	ctx.JSON(http.StatusOK, response)
}

func (jobItemWeighingInterface *JobItemWeighingInterface) Update(ctx *gin.Context) {
	response := value_objects.Response{}
	id := ctx.Param("id")

	requestingUser, ok := ctx.Get("user")
	if !ok {
		response.Status = false
		response.Message = "Anonymous User."
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
	model.UpdatedByUsername = user.Username

	updated, updationErr := jobItemWeighingInterface.appStore.JobItemWeighingApp.Update(id, &model)
	if updationErr != nil {
		response.Status = false
		response.Message = updationErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Return response.
	response.Status = true
	response.Message = "Job Weighing Updated."
	response.Payload = updated

	ctx.JSON(http.StatusOK, response)
}

func (jobItemWeighingInterface *JobItemWeighingInterface) Details(ctx *gin.Context) {
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

	jobItems, getErr := jobItemWeighingInterface.appStore.JobItemWeighingApp.Details(utilities.ConvertJSONToSQL(conditions))
	if getErr != nil {
		response.Status = false
		response.Message = getErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "Job Item Weighing Batches Found"
	response.Payload = jobItems

	ctx.JSON(http.StatusOK, response)
}

func (jobItemWeighingInterface *JobItemWeighingInterface) Materials(ctx *gin.Context) {
	response := value_objects.Response{}
	materialID := ctx.Param("material_id")

	jobItems, getErr := jobItemWeighingInterface.appStore.JobItemWeighingApp.MaterialDetails(materialID)
	if getErr != nil {
		response.Status = false
		response.Message = getErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "Job Item Weighing Materials Found"
	response.Payload = jobItems

	ctx.JSON(http.StatusOK, response)
}
