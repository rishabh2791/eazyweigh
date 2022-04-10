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

type UnderIssueInterface struct {
	appStore *application.AppStore
	logger   hclog.Logger
}

func NewUnderIssueInterface(appStore *application.AppStore, logger hclog.Logger) *UnderIssueInterface {
	return &UnderIssueInterface{
		appStore: appStore,
		logger:   logger,
	}
}

func (underIssueInterface *UnderIssueInterface) Create(ctx *gin.Context) {
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

	model := entity.UnderIssue{}
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

	created, creationErr := underIssueInterface.appStore.UnderIssueApp.Create(&model)
	if creationErr != nil {
		response.Status = false
		response.Message = creationErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "Under Issue Created."
	response.Payload = created

	ctx.JSON(http.StatusOK, response)
}

func (underIssueInterface *UnderIssueInterface) List(ctx *gin.Context) {
	response := value_objects.Response{}
	jobID := ctx.Param("job_id")

	underIssues, getErr := underIssueInterface.appStore.UnderIssueApp.List(jobID)
	if getErr != nil {
		response.Status = false
		response.Message = getErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "Under Issues Found"
	response.Payload = underIssues

	ctx.JSON(http.StatusOK, response)
}

func (underIssueInterface *UnderIssueInterface) Update(ctx *gin.Context) {
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

	model := entity.UnderIssue{}
	jsonErr := json.NewDecoder(ctx.Request.Body).Decode(&model)
	if jsonErr != nil {
		response.Status = false
		response.Message = jsonErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	model.UpdatedByUsername = user.Username

	updated, updationErr := underIssueInterface.appStore.UnderIssueApp.Update(id, &model)
	if updationErr != nil {
		response.Status = false
		response.Message = updationErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Return response.
	response.Status = true
	response.Message = "Under Issue Updated."
	response.Payload = updated

	ctx.JSON(http.StatusOK, response)
}
