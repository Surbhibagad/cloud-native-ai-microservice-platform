package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/Surbhibagad/cloud-native-ai-microservice-platform/services/auth-service/internal/handlers"
)

func SetupRoutes(router *gin.Engine, authHandler *handlers.AuthHandler) {

	auth := router.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}
}