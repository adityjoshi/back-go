package routes

import (
	"BACKEND-GO/controllers"

	"github.com/gin-gonic/gin"
)

func complaintRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/complaints", controllers.PostComplaints)

}
