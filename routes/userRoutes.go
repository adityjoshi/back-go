package routes

import (
	"github.com/gin-gonic/gin"
)

func UserRoute(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/register")
	incomingRoutes.POST("/login")
	incomingRoutes.GET("/warden", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hey": "i am a function",
		})
		c.Next()
	})
	incomingRoutes.GET("/worker")
	incomingRoutes.GET("/student")
	incomingRoutes.GET("/complaint")

}
