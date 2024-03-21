package routes

import (
	"BACKEND-GO/controllers"

	"github.com/gin-gonic/gin"
)

func studentRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/student", controllers.PostStudent)
	incomingRoutes.GET("/student/:student_id", controllers.GetStudentByID)
}
