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

type BOMItemInterface struct {
	appStore *application.AppStore
	logger   hclog.Logger
}

func NewBOMItemInterface(appStore *application.AppStore, logger hclog.Logger) *BOMItemInterface {
	return &BOMItemInterface{
		appStore: appStore,
		logger:   logger,
	}
}

func (bomItemInterface *BOMItemInterface) Create(ctx *gin.Context) {
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

	model := entity.BOMItem{}
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

	created, creationErr := bomItemInterface.appStore.BOMItemApp.Create(&model)
	if creationErr != nil {
		response.Status = false
		response.Message = creationErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "BOM Item Created."
	response.Payload = created

	ctx.JSON(http.StatusOK, response)
}

func (bomItemInterface *BOMItemInterface) CreateMultiple(ctx *gin.Context) {
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
	user := requestingUser.(*entity.User)

	models := []entity.BOMItem{}
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

		_, creationErr := bomItemInterface.appStore.BOMItemApp.Create(&model)
		if creationErr != nil {
			creationErrors = append(creationErrors, creationErr)
		} else {
			createdModels = append(createdModels, model)
		}
	}

	response.Status = true
	response.Message = "BOM Item Created."
	response.Payload = map[string]interface{}{
		"models": createdModels,
		"errors": creationErrors,
	}

	ctx.JSON(http.StatusOK, response)
}

func (bomItemInterface *BOMItemInterface) Get(ctx *gin.Context) {
	response := value_objects.Response{}
	id := ctx.Param("id")

	bom, getErr := bomItemInterface.appStore.BOMItemApp.Get(id)
	if getErr != nil {
		response.Status = false
		response.Message = getErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "BOM Item Found"
	response.Payload = bom

	ctx.JSON(http.StatusOK, response)
}

func (bomItemInterface *BOMItemInterface) List(ctx *gin.Context) {
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

	bomItems, getErr := bomItemInterface.appStore.BOMItemApp.List(utilities.ConvertJSONToSQL(conditions))
	if getErr != nil {
		response.Status = false
		response.Message = getErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "BOM Items Found"
	response.Payload = bomItems

	ctx.JSON(http.StatusOK, response)
}

func (bomItemInterface *BOMItemInterface) Update(ctx *gin.Context) {
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

	model := entity.BOMItem{}
	jsonErr := json.NewDecoder(ctx.Request.Body).Decode(&model)
	if jsonErr != nil {
		response.Status = false
		response.Message = jsonErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	model.UpdatedByUsername = user.Username

	updated, updationErr := bomItemInterface.appStore.BOMItemApp.Update(id, &model)
	if updationErr != nil {
		response.Status = false
		response.Message = updationErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Return response.
	response.Status = true
	response.Message = "BOM Item Updated."
	response.Payload = updated

	ctx.JSON(http.StatusOK, response)
}
