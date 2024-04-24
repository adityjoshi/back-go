package controllers

import (
	"BACKEND-GO/database"
	"BACKEND-GO/utils"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

// func GetAllComplaintsByUser(c *gin.Context) {
// 	token := c.GetHeader("Authorization")
// 	claims, err := utils.DecodeJWT(token)
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized token"})
// 		return
// 	}

// 	userType := claims["user"].(map[string]interface{})["type"].(string)
// 	if userType == "warden" {
// 		var allComplaints []database.Complaint
// 		if err := database.DB.Order("created_at DESC").Find(&allComplaints).Error; err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
// 			return
// 		}
// 		c.JSON(http.StatusOK, allComplaints)
// 	} else if userType == "student" {
// 		userID := int(claims["user"].(map[string]interface{})["user_id"].(float64))
// 		studentID, _, err := DecodeStudent(userID) // Ignoring blockID
// 		if err != nil {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized studentid and block id"})
// 			return
// 		}
// 		var myComplaints []database.Complaint
// 		if err := database.DB.Where("student_id = ?", studentID).Order("created_at DESC").Find(&myComplaints).Error; err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
// 			return
// 		}
// 		c.JSON(http.StatusOK, myComplaints)
// 		} else {
// 		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
// 	}
// }

//

type ComplaintResponse struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	BlockID     uint      `json:"block_id"`
	CreatedAt   time.Time `json:"created_at"`
	IsCompleted bool      `json:"is_completed"`
	ID          uint      `json:"id"`
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
		var allComplaints []ComplaintResponse
		if err := database.DB.Table("complaints").Select("name", "description", "block_id", "created_at", "is_completed", "id").Order("created_at DESC").Scan(&allComplaints).Error; err != nil {
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
		var myComplaints []ComplaintResponse
		if err := database.DB.Table("complaints").Select("name", "description", "block_id", "created_at", "is_completed", "id").Where("student_id = ?", studentID).Order("created_at DESC").Scan(&myComplaints).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		c.JSON(http.StatusOK, myComplaints)
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
	}
}

func GetComplaintByID(c *gin.Context) {
	// Get the complaint ID from the URL parameter
	complaintID := c.Param("id")

	// Retrieve the complaint from the database by its ID
	var complaint database.Complaint
	if err := database.DB.First(&complaint, complaintID).Error; err != nil {
		// If complaint not found, return 404 Not Found
		c.JSON(http.StatusNotFound, gin.H{"error": "Complaint not found"})
		return
	}

	// If complaint found, return it as JSON response
	c.JSON(http.StatusOK, complaint)
}

func PutComplaintsByid(c *gin.Context) {
	token := c.GetHeader("Authorization")
	claims, err := utils.DecodeJWT(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized token"})
		return
	}

	userType := claims["user"].(map[string]interface{})["type"].(string)
	if userType != "warden" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized user type"})
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid complaint ID"})
		return
	}

	var complaint database.Complaint
	result := database.DB.First(&complaint, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Complaint not found"})
		return
	}

	// Toggle is_completed field
	//complaint.IsCompleted = !complaint.IsCompleted
	var updatePayload map[string]bool
	if err := c.BindJSON(&updatePayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Update is_completed field if it exists in the request payload
	if val, ok := updatePayload["is_completed"]; ok {
		complaint.IsCompleted = val
	}

	// Update assigned_at timestamp to current time
	complaint.AssignedAt = time.Now()

	// Save updated complaint to database
	result = database.DB.Save(&complaint)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update complaint"})
		return
	}

	c.JSON(http.StatusOK, complaint)
}

func GetUserDetails(c *gin.Context) {
	token := c.GetHeader("Authorization")
	claims, err := utils.DecodeJWT(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized token"})
		return
	}

	userType := claims["user"].(map[string]interface{})["type"].(string)
	userID := int(claims["user"].(map[string]interface{})["user_id"].(float64))

	if userType == "student" {
		var studentDetails database.Student
		if err := database.DB.Table("users").Select("full_name", "email", "phone", "usn", "block_id", "room").
			Joins("JOIN students ON users.user_id = students.student_id").
			Where("users.user_id = ?", userID).
			First(&studentDetails).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			}
			return
		}
		c.JSON(http.StatusOK, studentDetails)
	} else if userType == "warden" {
		var wardenDetails database.User
		if err := database.DB.Table("users").Select("full_name", "email", "phone").
			Where("users.user_id = ?", userID).
			First(&wardenDetails).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Warden not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			}
			return
		}
		c.JSON(http.StatusOK, wardenDetails)
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
