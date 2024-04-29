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
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

func MailService(complaint database.Complaint, studentEmail string) error {
	message := gomail.NewMessage()
	message.SetHeader("From", "aditya3.collegeboard@gmail.com")
	message.SetHeader("To", "aditya03joshi@gmail.com", studentEmail)
	message.SetHeader("Subject", "New Complaint Filed Of Room Number:", complaint.Room)

	// Construct the email body with dynamic complaint details
	body := fmt.Sprintf("Dear student, thank you for filling the complaint form. Below are the details of your complaint:\n\n")
	body += fmt.Sprintf("Name: %s\n", complaint.Name)
	body += fmt.Sprintf("Description: %s\n", complaint.Description)
	body += fmt.Sprintf("Room: %s\n", complaint.Room)
	body += fmt.Sprintf("Complaint Type: %s\n", complaint.ComplaintIssues)
	body += fmt.Sprintf("Complaint Created: %s\n", complaint.CreatedAt)
	body += fmt.Sprint("Sorry for the incovience we will fix it as soon as possible.\n")
	body += fmt.Sprintf("<b>Hostel Team</b>\n")
	body += fmt.Sprintf("<b>Block 4</b>")
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

func MailServiceDone(complaint database.Complaint, studentEmail string) error {
	message := gomail.NewMessage()
	message.SetHeader("From", "aditya3.collegeboard@gmail.com")
	message.SetHeader("To", "aditya03joshi@gmail.com", studentEmail)
	message.SetHeader("Subject", "Complaint Completed : ", complaint.Room)

	// Construct the email body with dynamic complaint details
	body := fmt.Sprintf("Dear student,your issue has been resolved. Sorry for the inconvenience :\n\n")
	body += fmt.Sprintf("Name: %s\n", complaint.Name)
	body += fmt.Sprintf("Description: %s\n", complaint.Description)
	body += fmt.Sprintf("Room: %s\n", complaint.Room)
	body += fmt.Sprintf("Complaint Type: %s\n", complaint.ComplaintIssues)
	body += fmt.Sprintf("Created At: %s\n", complaint.CreatedAt)
	body += fmt.Sprintf("Completed At: %s\n", complaint.AssignedAt)
	body += fmt.Sprint("Thank you for your patience and understanding. We appreciate your wait and apologize for any inconvenience caused. If you have any further questions or concerns, please feel free to reach out to us. Your satisfaction is our priority.")
	body += fmt.Sprintf("**Hostel Team**\n")
	body += fmt.Sprintf("**Block 4**")
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
	studentEmail, err := GetStudentEmailByID(int(studentID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve student email"})
		return
	}

	if err := MailService(complaint, studentEmail); err != nil {
		fmt.Println("Failed to send email notification:", err)
	}

	c.JSON(http.StatusCreated, complaint)
}
func GetStudentEmailByID(studentID int) (string, error) {
	var student database.Student
	if err := database.DB.Where("student_id = ?", studentID).First(&student).Error; err != nil {
		return "", err
	}
	return student.Email, nil
}

type ComplaintResponse struct {
	Name            string                 `json:"name"`
	Description     string                 `json:"description"`
	BlockID         uint                   `json:"block_id"`
	CreatedAt       time.Time              `json:"created_at"`
	IsCompleted     bool                   `json:"is_completed"`
	ID              uint                   `json:"id"`
	ComplaintIssues database.ComplaintType `json:"complaint_issues"`
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
// 		var allComplaints []ComplaintResponse
// 		if err := database.DB.Table("complaints").Select("name", "description", "block_id", "created_at", "is_completed", "id", "complaint_issues").Order("created_at DESC").Scan(&allComplaints).Error; err != nil {
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
// 		var myComplaints []ComplaintResponse
// 		if err := database.DB.Table("complaints").Select("name", "description", "block_id", "created_at", "is_completed", "id", "complaint_issues").Where("student_id = ?", studentID).Order("created_at DESC").Scan(&myComplaints).Error; err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
// 			return
// 		}
// 		c.JSON(http.StatusOK, myComplaints)
// 	} else {
// 		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
// 	}
// }

// func GetAllComplaintsByUser(c *gin.Context) {
// 	token := c.GetHeader("Authorization")
// 	claims, err := utils.DecodeJWT(token)
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized token"})
// 		return
// 	}

// 	userType := claims["user"].(map[string]interface{})["type"].(string)
// 	if userType == "warden" {
// 		wardenBlockID, ok := claims["user"].(map[string]interface{})["block_id"].(float64)
// 		if !ok {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Block ID not found or invalid type in claims"})
// 			return
// 		}

// 		var allComplaints []ComplaintResponse
// 		if err := database.DB.Table("complaints").Select("name", "description", "block_id", "created_at", "is_completed", "id", "complaint_issues").Where("block_id = ?", wardenBlockID).Order("created_at DESC").Scan(&allComplaints).Error; err != nil {
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
// 		var myComplaints []ComplaintResponse
// 		if err := database.DB.Table("complaints").Select("name", "description", "block_id", "created_at", "is_completed", "id", "complaint_issues").Where("student_id = ?", studentID).Order("created_at DESC").Scan(&myComplaints).Error; err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
// 			return
// 		}
// 		c.JSON(http.StatusOK, myComplaints)
// 	} else {
// 		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
// 	}
// }

func GetAllComplaintsByUser(c *gin.Context) {
	token := c.GetHeader("Authorization")
	claims, err := utils.DecodeJWT(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized token"})
		return
	}

	userType := claims["user"].(map[string]interface{})["type"].(string)
	var filteredComplaints []ComplaintResponse

	if userType == "warden" {
		wardenBlockID, ok := claims["user"].(map[string]interface{})["block_id"].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Block ID not found or invalid type in claims"})
			return
		}

		// Fetch all complaints for warden
		if err := database.DB.Table("complaints").Select("name", "description", "block_id", "created_at", "is_completed", "id", "complaint_issues").Where("block_id = ?", wardenBlockID).Order("created_at DESC").Scan(&filteredComplaints).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
	} else if userType == "student" {
		userID := int(claims["user"].(map[string]interface{})["user_id"].(float64))
		studentID, _, err := DecodeStudent(userID) // Ignoring blockID
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized studentid and block id"})
			return
		}

		// Fetch complaints for student
		if err := database.DB.Table("complaints").Select("name", "description", "block_id", "created_at", "is_completed", "id", "complaint_issues").Where("student_id = ?", studentID).Order("created_at DESC").Scan(&filteredComplaints).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
		return
	}

	// Extract complaint type query parameter
	complaintType := c.Query("complaintType")
	if complaintType != "" {
		// Filter complaints based on complaint type
		var filteredByType []ComplaintResponse
		for _, complaint := range filteredComplaints {
			if string(complaint.ComplaintIssues) == complaintType {
				filteredByType = append(filteredByType, complaint)
			}
		}
		filteredComplaints = filteredByType
	}

	c.JSON(http.StatusOK, filteredComplaints)
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
	studentEmail, err := GetStudentEmail(int(complaint.StudentID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve student email"})
		return
	}
	if err := MailServiceDone(complaint, studentEmail); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email notification"})
		return
	}

	c.JSON(http.StatusOK, complaint)
}
func GetStudentEmail(studentID int) (string, error) {
	var student database.Student
	if err := database.DB.Where("student_id = ?", studentID).First(&student).Error; err != nil {
		return "", err
	}
	return student.Email, nil
}

type UserDetailsResponse struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`

	// Add other fields as needed
}
type StudentDetailsResponse struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	USN      string `json:"usn"`
	BlockID  uint   `json:"block_id"`
	Room     string `json:"room"`
	// Add other fields as needed
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
		var studentDetails database.User
		result := database.DB.Where("user_id = ?", userID).First(&studentDetails)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Student details not found"})
			return
		}
		studentDetailsResponse := StudentDetailsResponse{
			FullName: studentDetails.FullName,
			Email:    studentDetails.Email,
			Phone:    studentDetails.Phone,
			USN:      studentDetails.USN,
			BlockID:  studentDetails.BlockID,
			Room:     studentDetails.Room,
			// Add other fields as needed
		}

		c.JSON(http.StatusOK, studentDetailsResponse)
	} else if userType == "warden" {
		var userDetails database.User
		result := database.DB.Where("user_id = ?", userID).First(&userDetails)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User details not found"})
			return
		}
		userDetailsResponse := UserDetailsResponse{
			FullName: userDetails.FullName,
			Email:    userDetails.Email,
			Phone:    userDetails.Phone,

			// Add other fields as needed
		}
		c.JSON(http.StatusOK, userDetailsResponse)
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
	}
}

func DeleteComplaints() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		claims, err := utils.DecodeJWT(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized token"})
			return
		}

		userType := claims["user"].(map[string]interface{})["type"].(string)
		id := c.Param("id")

		if userType == "warden" {
			// Perform deletion
			result := database.DB.Where("id = ?", id).Delete(&database.Complaint{})
			if result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete complaint"})
				return
			}
			if result.RowsAffected == 0 {
				c.JSON(http.StatusNotFound, gin.H{"error": "Complaint not found"})
				return
			}

			c.JSON(http.StatusOK, gin.H{"message": "Complaint deleted"})
			return
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "Permission Denied"})
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

func InitiateEmailNotifications() {
	ticker := time.NewTicker(5 * time.Second) // Check every 5 seconds
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			checkAndSendOverdueNotifications()
		}
	}
}

// checkAndSendOverdueNotifications checks for overdue complaints and sends email notifications if needed
func checkAndSendOverdueNotifications() {
	// Query the database for complaints that are not completed and whose completion time has exceeded the threshold
	complaints := getOverdueComplaints()

	for _, complaint := range complaints {
		// Send email notification for each overdue complaint
		err := SendEmailNotification(complaint)
		if err != nil {
			fmt.Println("Error sending email notification for complaint:", complaint.ID, "-", err)
		} else {
			fmt.Println("Email notification sent for overdue complaint:", complaint.ID)
		}
	}
}

// getOverdueComplaints retrieves complaints that are overdue for completion
func getOverdueComplaints() []database.Complaint {
	var overdueComplaints []database.Complaint

	// Fetch complaints where IsCompleted is false and (CurrentTime - CreatedAt) > Threshold
	threshold := time.Now().Add(-time.Second * 5) // Example threshold: 5 minutes
	database.DB.Where("is_completed = ?", false).Where("created_at < ?", threshold).Find(&overdueComplaints)

	return overdueComplaints
}

// SendEmailNotification sends an email notification for the given complaint
func SendEmailNotification(complaint database.Complaint) error {

	message := gomail.NewMessage()
	message.SetHeader("From", "aditya3.collegeboard@gmail.com")
	message.SetHeader("To", "aditya03joshi@gmail.com")
	message.SetHeader("Subject", "Complaint Overdue Notification")

	// Construct the email body
	body := fmt.Sprintf("Dear user,\n\nYour complaint with ID %d is overdue for completion. Please take necessary action.\n\n", complaint.ID)
	message.SetBody("text/plain", body)

	// Initialize SMTP dialer
	dialer := gomail.NewDialer("smtp.gmail.com", 587, "aditya3.collegeboard@gmail.com", "ehnxaubjqelkotks") // Update with your SMTP server details

	// Send email
	if err := dialer.DialAndSend(message); err != nil {
		return err
	}

	return nil
}

// StartEmailNotificationScheduler starts the Goroutine for periodic email notifications
func StartEmailNotificationScheduler() {
	go InitiateEmailNotifications()
}
