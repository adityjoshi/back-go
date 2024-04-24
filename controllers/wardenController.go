package controllers

import (
	"BACKEND-GO/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostWarden(c *gin.Context) {
	var warden database.Warden

	// Bind request body to struct
	if err := c.ShouldBindJSON(&warden); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create the warden record in the database
	if err := database.DB.Create(&warden).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the inserted warden
	c.JSON(http.StatusCreated, warden)
}
