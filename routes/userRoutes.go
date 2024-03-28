package routes

import (
	"BACKEND-GO/controllers"
	"BACKEND-GO/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoute(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/register", controllers.Register)
	incomingRoutes.POST("/login", controllers.Login)
	incomingRoutes.GET("/student", middleware.AuthorizeStudent(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "This function is accessible by students only",
		})
		c.Next()
	})
	incomingRoutes.GET("/userType", controllers.GetUserType)

}
