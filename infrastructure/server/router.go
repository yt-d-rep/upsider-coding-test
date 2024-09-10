package server

import (
	"upsider-base/infrastructure/di"

	"github.com/gin-gonic/gin"
)

func route(router *gin.Engine) {
	// users
	router.POST("/api/users", di.Wire().UserHandler.Register)
	router.POST("/api/login", di.Wire().UserHandler.Login)
}
