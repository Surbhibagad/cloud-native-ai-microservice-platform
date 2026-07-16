package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/Surbhibagad/cloud-native-ai-microservice-platform/services/auth-service/internal/utils"
)

func JWTMiddleware(jwtSecret string) gin.HandlerFunc {

	return func(c *gin.Context) {

		// Read Authorization header
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header is missing",
			})
			c.Abort()
			return
		}

		// Expected format:
		// Authorization: Bearer <token>

		tokenParts := strings.Split(authHeader, " ")

		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid Authorization header",
			})
			c.Abort()
			return
		}

		token := tokenParts[1]

		// Validate JWT
		userID, err := utils.ValidateJWT(token, jwtSecret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid or expired token",
			})
			c.Abort()
			return
		}

		// Store user ID in Gin context
		c.Set("userID", userID)

		// Continue to next handler
		c.Next()
	}
}