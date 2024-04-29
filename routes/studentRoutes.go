package routes

import (
	"BACKEND-GO/controllers"

	"github.com/gin-gonic/gin"
)

func StudentRoutes(incomingRoutes *gin.Engine) {
	//incomingRoutes.POST("/student", controllers.PostStudent)
	//incomingRoutes.GET("/student", controllers.GetStudentByID)
	incomingRoutes.GET("/student/:student_id", controllers.GetStudentByID)

}
