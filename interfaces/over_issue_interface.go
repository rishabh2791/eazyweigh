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

type OverIssueInterface struct {
	appStore *application.AppStore
	logger   hclog.Logger
}

func NewOverIssueInterface(appStore *application.AppStore, logger hclog.Logger) *OverIssueInterface {
	return &OverIssueInterface{
		appStore: appStore,
		logger:   logger,
	}
}

func (overIssueInterface *OverIssueInterface) Create(ctx *gin.Context) {
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

	model := entity.OverIssue{}
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

	created, creationErr := overIssueInterface.appStore.OverIssueApp.Create(&model)
	if creationErr != nil {
		response.Status = false
		response.Message = creationErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "Over Issue Created."
	response.Payload = created

	ctx.JSON(http.StatusOK, response)
}

func (overIssueInterface *OverIssueInterface) List(ctx *gin.Context) {
	response := value_objects.Response{}
	jobID := ctx.Param("job_id")

	overIssues, getErr := overIssueInterface.appStore.OverIssueApp.List(jobID)
	if getErr != nil {
		response.Status = false
		response.Message = getErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "Over Issues Found"
	response.Payload = overIssues

	ctx.JSON(http.StatusOK, response)
}

func (overIssueInterface *OverIssueInterface) Update(ctx *gin.Context) {
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

	model := entity.OverIssue{}
	jsonErr := json.NewDecoder(ctx.Request.Body).Decode(&model)
	if jsonErr != nil {
		response.Status = false
		response.Message = jsonErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	model.UpdatedByUsername = user.Username

	updated, updationErr := overIssueInterface.appStore.OverIssueApp.Update(id, &model)
	if updationErr != nil {
		response.Status = false
		response.Message = updationErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Return response.
	response.Status = true
	response.Message = "Over Issue Updated."
	response.Payload = updated

	ctx.JSON(http.StatusOK, response)
}

func (overIssueInterface *OverIssueInterface) Approve(ctx *gin.Context) {
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

	model := entity.OverIssue{
		Approved:          true,
		UpdatedByUsername: user.Username,
	}

	updated, updationErr := overIssueInterface.appStore.OverIssueApp.Update(id, &model)
	if updationErr != nil {
		response.Status = false
		response.Message = updationErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Return response.
	response.Status = true
	response.Message = "over Issue Approved."
	response.Payload = updated

	ctx.JSON(http.StatusOK, response)
}

func (overIssueInterface *OverIssueInterface) Reject(ctx *gin.Context) {
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

	model := entity.OverIssue{
		Approved:          false,
		UpdatedByUsername: user.Username,
	}

	updated, updationErr := overIssueInterface.appStore.OverIssueApp.Update(id, &model)
	if updationErr != nil {
		response.Status = false
		response.Message = updationErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Return response.
	response.Status = true
	response.Message = "over Issue Rejected."
	response.Payload = updated

	ctx.JSON(http.StatusOK, response)
}
