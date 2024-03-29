package controllers

import (
	"BACKEND-GO/database"
	"BACKEND-GO/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostComplaints(c *gin.Context) {
	token := c.GetHeader("Authorization")
	claims, err := utils.DecodeJWT(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized token"})
		return
	}

	userType := claims["user"].(map[string]interface{})["type"].(string)
	if userType != "student" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized student type"})
		return
	}

	userID := int(claims["user"].(map[string]interface{})["user_id"].(float64))
	studentID, blockID, err := DecodeStudent(userID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized studentid and block id"})
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

func GetAllComplaintsByUser(c *gin.Context) {
	token := c.GetHeader("Authorization")
	claims, err := utils.DecodeJWT(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized token"})
		return
	}

	userType := claims["user"].(map[string]interface{})["type"].(string)
	if userType == "warden" {
		var allComplaints []database.Complaint
		if err := database.DB.Order("created_at DESC").Find(&allComplaints).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		c.JSON(http.StatusOK, allComplaints)
	} else if userType == "student" {
		userID := int(claims["user"].(map[string]interface{})["user_id"].(float64))
		studentID, _, err := DecodeStudent(userID) // Ignoring blockID
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized studentid and block id"})
			return
		}
		var myComplaints []database.Complaint
		if err := database.DB.Where("student_id = ?", studentID).Order("created_at DESC").Find(&myComplaints).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		c.JSON(http.StatusOK, myComplaints)
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
	}
}

func DecodeStudent(userID int) (studentID, blockID uint, err error) {
	var student database.Student
	if err := database.DB.Where("student_id = ?", userID).Find(&student).Error; err != nil {
		fmt.Println("Error fetching student details:", err)
		return 0, 0, err
	}
	fmt.Println("Student details:", student)
	return student.StudentID, student.BlockID, nil
}
