package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/Surbhibagad/cloud-native-ai-microservice-platform/services/auth-service/internal/config"
	"github.com/Surbhibagad/cloud-native-ai-microservice-platform/services/auth-service/internal/database"
)

func main() {

	cfg := config.LoadConfig()

	db := database.Connect(cfg)
	defer db.Close(context.Background())

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Auth Service Running 🚀",
		})
	})

	log.Println("Server running on port", cfg.Port)

	router.Run(":" + cfg.Port)
}