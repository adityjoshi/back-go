package controllers

import (
	"BACKEND-GO/database"
	"BACKEND-GO/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostComplaints(c *gin.Context) {
	token := c.GetHeader("Authorization")
	claims, err := utils.DecodeJWT(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userType := claims["user"].(map[string]interface{})["type"].(string)
	if userType != "student" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userID := int(claims["user"].(map[string]interface{})["user_id"].(float64))
	studentID, blockID, err := DecodeStudent(userID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var complaint database.Complaint
	if err := c.ShouldBindJSON(&complaint); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	complaint.StudentID = studentID
	complaint.BlockID = blockID

	if err := database.DB.Create(&complaint).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusCreated, complaint)
}

// DecodeStudent decodes the student ID and block ID based on the provided user ID
func DecodeStudent(userID int) (studentID, blockID uint, err error) {
	var student database.Student
	if err := database.DB.Where("student_id = ?", userID).First(&student).Error; err != nil {
		return 0, 0, err
	}
	return student.StudentID, student.BlockID, nil
}
