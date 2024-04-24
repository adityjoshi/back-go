package routes

import (
	"github.com/gin-gonic/gin"
)

func WardenRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/warden")

}
