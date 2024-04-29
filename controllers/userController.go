package controllers

import (
	"BACKEND-GO/database"

	"BACKEND-GO/utils"
	"fmt"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"

	"net/http"
)

func MailServices(complaint database.User) error {
	message := gomail.NewMessage()
	message.SetHeader("From", "aditya3.collegeboard@gmail.com")
	message.SetHeader("To", complaint.Email, "aditya30joshi@gmail.com")
	message.SetHeader("Subject", "Hostel Ease Log In")

	// Construct the email body with dynamic complaint details
	body := fmt.Sprintf("Dear student, thank you for logging in to the hostel ease. If it's not you reach out to us asap through hostelvit@gmail.com\n\n")
	body += fmt.Sprintf("*Hostel Team*\n")
	body += fmt.Sprintf("*Block 4*")
	message.SetBody("text/plain", body)

	//message.Attach("/home/Alex/lolcat.jpg")

	// Initialize SMTP dialer
	dialer := gomail.NewDialer("smtp.gmail.com", 587, "aditya3.collegeboard@gmail.com", "ehnxaubjqelkotks") // Update with your SMTP server details

	// Send email
	if err := dialer.DialAndSend(message); err != nil {
		panic(err)
	}

	fmt.Println("Email sent successfully!")
	return nil
}

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
		newStudent.FullName = newUser.FullName
		newStudent.Email = newUser.Email

		if err := database.DB.Create(&newStudent).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create student record"})
			return
		}
	} else if newUser.Type == "warden" {
		var newWarden database.Warden
		newWarden.Warden_Id = newUser.UserID
		newWarden.BlockID = newUser.BlockID

		if err := database.DB.Create(&newWarden).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create warden record"})
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
	jwtToken, err := utils.GenerateJWT(int(user.UserID), user.Type, user.BlockID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate JWT token"})
		return
	}
	if err := MailServices(user); err != nil {
		fmt.Println("Failed to send email notification:", err)
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
