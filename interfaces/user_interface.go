package interfaces

import (
	"eazyweigh/application"
	"eazyweigh/domain/entity"
	"eazyweigh/domain/value_objects"
	"eazyweigh/infrastructure/utilities"
	"eazyweigh/infrastructure/utilities/security"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-hclog"
)

type UserInterface struct {
	appStore *application.AppStore
	logger   hclog.Logger
}

func NewUserInterface(appStore *application.AppStore, logging hclog.Logger) *UserInterface {
	return &UserInterface{
		appStore: appStore,
		logger:   logging,
	}
}

func (userInterface *UserInterface) Create(ctx *gin.Context) {
	response := value_objects.Response{}

	// Check Requesting User.
	requestingUser, ok := ctx.Get("user")
	if !ok {
		response.Status = false
		response.Message = "Anonymous User."
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	user := requestingUser.(*entity.User)

	// Get new entry details from request body.
	model := entity.User{}
	jsonErr := json.NewDecoder(ctx.Request.Body).Decode(&model)
	if jsonErr != nil {
		response.Status = false
		response.Message = jsonErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	// Complete model details
	model.CreatedByUsername = user.Username
	model.UpdatedByUsername = user.Username

	// Create entry in database.
	created, creationErr := userInterface.appStore.UserApp.Create(&model, "register")
	if creationErr != nil {
		response.Status = false
		response.Message = creationErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	// Return response.
	response.Status = true
	response.Message = "User Created."
	response.Payload = created

	ctx.JSON(http.StatusOK, response)
}

func (userInterface *UserInterface) CreateMultiple(ctx *gin.Context) {
	response := value_objects.Response{}
	createdModels := []interface{}{}
	creationErrors := []interface{}{}

	// Check Requesting User.
	requestingUser, ok := ctx.Get("user")
	if !ok {
		response.Status = false
		response.Message = "Anonymous User."
		response.Payload = map[string]interface{}{
			"models": []string{},
			"errors": []string{"Anonymous User."},
		}

		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	user := requestingUser.(*entity.User)

	// Get new entry details from request body.
	models := []entity.User{}
	jsonErr := json.NewDecoder(ctx.Request.Body).Decode(&models)
	if jsonErr != nil {
		response.Status = false
		response.Message = jsonErr.Error()
		response.Payload = map[string]interface{}{
			"models": []string{},
			"errors": []string{jsonErr.Error()},
		}

		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	for _, model := range models {
		// Complete model details
		model.CreatedByUsername = user.Username
		model.UpdatedByUsername = user.Username

		// Create entry in database.
		_, creationErr := userInterface.appStore.UserApp.Create(&model, "register")
		if creationErr != nil {
			creationErrors = append(creationErrors, creationErr)
		} else {
			createdModels = append(createdModels, model)
		}
	}

	// Return response.
	response.Status = true
	response.Message = "User Created."
	response.Payload = map[string]interface{}{
		"models": createdModels,
		"errors": creationErrors,
	}

	ctx.JSON(http.StatusOK, response)
}

func (userInterface *UserInterface) Get(ctx *gin.Context) {
	response := value_objects.Response{}

	username := ctx.Param("username")
	user, err := userInterface.appStore.UserApp.Get(username)
	if err != nil {
		response.Status = false
		response.Message = err.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "User Found"
	response.Payload = user

	ctx.JSON(http.StatusOK, response)
}

func (userInterface *UserInterface) List(ctx *gin.Context) {
	response := value_objects.Response{}

	conditions := map[string]interface{}{}

	jsonError := json.NewDecoder(ctx.Request.Body).Decode(&conditions)
	if jsonError != nil {
		response.Status = false
		response.Message = jsonError.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	users, err := userInterface.appStore.UserApp.List(utilities.ConvertJSONToSQL(conditions))
	if err != nil {
		response.Status = false
		response.Message = err.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response.Status = true
	response.Message = "Users Found"
	response.Payload = users

	ctx.JSON(http.StatusOK, response)
}

func (userInterface *UserInterface) Update(ctx *gin.Context) {
	response := value_objects.Response{}
	username := ctx.Param("username")

	// Check Requesting User.
	requestingUser, ok := ctx.Get("user")
	if !ok {
		response.Status = false
		response.Message = "Anonymous User."
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	user := requestingUser.(*entity.User)

	// Get new entry details from request body.
	model := map[string]interface{}{}
	jsonErr := json.NewDecoder(ctx.Request.Body).Decode(&model)
	if jsonErr != nil {
		response.Status = false
		response.Message = jsonErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	log.Println(model)

	// Complete model details
	model["updated_by_username"] = user.Username
	if len(model["password"].(string)) != 0 {
		hashedPass, err := security.Hash(model["password"].(string))
		if err != nil {
		} else {
			model["password"] = hashedPass
		}
	}

	log.Println(model)

	// Create entry in database.
	updated, creationErr := userInterface.appStore.UserApp.Update(username, model, user.Username)
	if creationErr != nil {
		response.Status = false
		response.Message = creationErr.Error()
		response.Payload = ""

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// Return response.
	response.Status = true
	response.Message = "User Updated."
	response.Payload = updated

	ctx.JSON(http.StatusOK, response)
}
