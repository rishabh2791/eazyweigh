package middlewares

import (
	"eazyweigh/application"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-hclog"
)

type PermissionMiddleware struct {
	Logger   hclog.Logger
	AppStore *application.AppStore
}

func NewPermissionMiddleware(logger hclog.Logger, appStore application.AppStore) *PermissionMiddleware {
	permissionMiddleware := PermissionMiddleware{}
	permissionMiddleware.Logger = logger
	permissionMiddleware.AppStore = &appStore
	return &permissionMiddleware
}

func (permission *PermissionMiddleware) HasPermission() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Handle Permissions
		ctx.Next()
	}
}

func (permission *PermissionMiddleware) HasObjectPermission(ctx *gin.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}

func (permission *PermissionMiddleware) HasTablePermission(ctx *gin.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}
