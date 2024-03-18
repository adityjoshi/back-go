package routes

import (
	"github.com/gin-gonic/gin"
)

func UserRoute(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/register")
	incomingRoutes.POST("/login")
	incomingRoutes.GET("/warden", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "This function is accessible by wardens only",
		})
		c.Next()
	})
	incomingRoutes.GET("/worker", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "This function is accessible by worker only",
		})
		c.Next()
	})
	incomingRoutes.GET("/student", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "This function is accessible by students only",
		})
		c.Next()
	})
	incomingRoutes.GET("/complaint", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "This function is accessible by complaint only",
		})
		c.Next()
	})

}
