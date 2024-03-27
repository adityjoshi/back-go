package middleware

import (
	"net/http"

	"BACKEND-GO/utils" // Update with correct path

	"github.com/gin-gonic/gin"
)

func AuthorizeWarden() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is missing"})
			c.Abort()
			return
		}

		claims, err := utils.DecodeJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		userType, ok := claims["user"].(map[string]interface{})["type"].(string)
		if !ok || userType != "warden" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized for warden"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func AuthorizeStudent() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is missing"})
			c.Abort()
			return
		}

		claims, err := utils.DecodeJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		userType, ok := claims["user"].(map[string]interface{})["type"].(string)
		if !ok || userType != "student" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized for student"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func AuthorizeComplaint() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is missing"})
			c.Abort()
			return
		}

		claims, err := utils.DecodeJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		userType, ok := claims["user"].(map[string]interface{})["type"].(string)
		if !ok || userType != "student" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized for student"})
			c.Abort()
			return
		}

		c.Next()
	}
}
