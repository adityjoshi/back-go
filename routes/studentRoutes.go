package routes

import (
	"github.com/gin-gonic/gin"
)

func studentRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/student")
	incomingRoutes.POST("/student/:student_id")
}
