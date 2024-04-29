package routes

import (
	"BACKEND-GO/controllers"
	"BACKEND-GO/middleware"

	"github.com/gin-gonic/gin"
)

func ComplaintRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/complaints", middleware.AuthorizeComplaint(), controllers.PostComplaints)
	incomingRoutes.GET("/complaints", controllers.GetAllComplaintsByUser)
	incomingRoutes.GET("/compalints/:id", controllers.GetComplaintByID)
	incomingRoutes.PUT("/complaints/:id", controllers.PutComplaintsByid)
	incomingRoutes.GET("/userDetails/:id", controllers.GetUserDetails)
	incomingRoutes.DELETE("/complaints/:id", controllers.DeleteComplaints())
}
