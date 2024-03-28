package controllers

import (
	"BACKEND-GO/database"

	"BACKEND-GO/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"net/http"
)

func Register(c *gin.Context) {
	var newUser database.User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if user already exists
	var existingUser database.User
	if err := database.DB.Where("email = ?", newUser.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User already exists"})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Create the new user record
	newUser.Password = string(hashedPassword)
	if err := database.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	if newUser.Type == "student" {
		var newStudent database.Student
		newStudent.USN = newUser.USN
		newStudent.Room = newUser.Room
		newStudent.StudentID = newUser.UserID
		newStudent.BlockID = newUser.BlockID

		if err := database.DB.Create(&newStudent).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create student record"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	var loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the user by email
	var user database.User
	if err := database.DB.Where("email = ?", loginRequest.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token
	jwtToken, err := utils.GenerateJWT(int(user.UserID), user.Type)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate JWT token"})
		return
	}

	//c.JSON(http.StatusOK, gin.H{"jwtToken": jwtToken})
	c.JSON(http.StatusOK, gin.H{"jwtToken": jwtToken, "userType": user.Type})
}

//

func GetUserType(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization header missing"})
		return
	}

	claims, err := utils.DecodeJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
		return
	}

	userType := claims["user"].(map[string]interface{})["type"].(string)
	c.JSON(http.StatusOK, gin.H{"userType": userType})
}
