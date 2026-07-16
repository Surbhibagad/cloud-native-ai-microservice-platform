package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/Surbhibagad/cloud-native-ai-microservice-platform/services/auth-service/internal/config"
	"github.com/Surbhibagad/cloud-native-ai-microservice-platform/services/auth-service/internal/handlers"
	"github.com/Surbhibagad/cloud-native-ai-microservice-platform/services/auth-service/internal/middleware"
)

func SetupRoutes(router *gin.Engine, authHandler *handlers.AuthHandler, cfg *config.Config) {

	auth := router.Group("/auth")
	{
		// Public Routes
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)

		// Protected Routes
		protected := auth.Group("/")
		protected.Use(middleware.JWTMiddleware(cfg.JWTSecret))

		protected.GET("/profile", authHandler.Profile)
	}
}