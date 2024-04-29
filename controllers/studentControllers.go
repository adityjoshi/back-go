package controllers

import (
	"BACKEND-GO/database"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PostStudent(c *gin.Context) {
	var student database.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := database.DB.Create(&student).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, student)
}

// func GetStudentByID(c *gin.Context) {
// 	studentID := c.Param("student_id")
// 	var student database.Student
// 	fmt.Print(studentID)

// 	err := database.DB.Find(&student, studentID).Error
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, student)
// }

// func GetStudentByID(c *gin.Context) {
// 	studentID := c.Param("student_id")
// 	var student database.Student
// 	fmt.Print(studentID)

// 	err := database.DB.Find(&student, "student_id = ?", studentID).Error
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, student)
// }

func GetStudentByID(c *gin.Context) {
	// Get the student ID from the request parameter
	studentIDStr := c.Param("student_id")

	// Convert the student ID string to uint
	studentID, err := strconv.ParseUint(studentIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	// Initialize a variable to store student details
	var studentDetails database.Student

	// Query the database to fetch student details by student ID
	if err := database.DB.Select("full_name, email, phone, usn, block_id, room").Where("student_id = ?", studentID).Find(&studentDetails).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	// Return the student details in the response
	c.JSON(http.StatusOK, studentDetails)
}
