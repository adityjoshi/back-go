package routes

import (
	"BACKEND-GO/controllers"
	"BACKEND-GO/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoute(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/register", controllers.Register)
	incomingRoutes.POST("/login", controllers.Login)
	// incomingRoutes.GET("/warden", middleware.AuthorizeWarden(), func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "This function is accessible by wardens only",
	// 	})
	// 	c.Next()
	// })
	// 	incomingRoutes.GET("/worker", func(c *gin.Context) {
	// 		c.JSON(200, gin.H{
	// 			"message": "This function is accessible by worker only",
	// 		})
	// 		c.Next()
	// 	})
	incomingRoutes.GET("/student", middleware.AuthorizeStudent(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "This function is accessible by students only",
		})
		c.Next()
	})

	// incomingRoutes.GET("/complaint", middleware.AuthorizeComplaint(), func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "This function is accessible by complaint only",
	// 	})
	// 	c.Next()
	// })

	// incomingRoutes.POST("/compl", middleware.AuthorizeComplaint(), func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "This function is by complaint only",
	// 	})
	// 	c.Next()
	// })

	// }

}
