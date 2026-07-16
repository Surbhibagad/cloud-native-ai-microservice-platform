package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/Surbhibagad/cloud-native-ai-microservice-platform/services/auth-service/internal/config"
	"github.com/Surbhibagad/cloud-native-ai-microservice-platform/services/auth-service/internal/database"
	"github.com/Surbhibagad/cloud-native-ai-microservice-platform/services/auth-service/internal/handlers"
	"github.com/Surbhibagad/cloud-native-ai-microservice-platform/services/auth-service/internal/repository"
	"github.com/Surbhibagad/cloud-native-ai-microservice-platform/services/auth-service/internal/routes"
	"github.com/Surbhibagad/cloud-native-ai-microservice-platform/services/auth-service/internal/services"
)

func main() {

	cfg := config.LoadConfig()

	db := database.Connect(cfg)
	defer db.Close(context.Background())

	// Repository
	userRepo := repository.NewUserRepository(db)

	// Service
	authService := services.NewAuthService(userRepo)

	// Handler
	authHandler := handlers.NewAuthHandler(authService, cfg.JWTSecret)

	router := gin.Default()

	// Health Check
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Auth Service Running 🚀",
		})
	})

	// API Routes
	routes.SetupRoutes(router, authHandler, cfg)

	log.Println("Server running on port", cfg.Port)

	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}