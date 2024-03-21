package controllers

import (
	"BACKEND-GO/database"
	"net/http"

	"github.com/gin-gonic/gin"
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

func GetStudentByID(c *gin.Context) {
	studentID := c.Param("student_id")
	var student database.Student

	err := database.DB.First(&student, studentID).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	c.JSON(http.StatusOK, student)
}
