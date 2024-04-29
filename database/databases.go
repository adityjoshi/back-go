package database

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	dsn := "host=localhost user=postgres password=aditya dbname=hostel port=5432"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Print("database connected successfully ⚡️")
	}

	// Migrate the schema
	DB.AutoMigrate(&User{}, &Block{}, &Student{}, &Category{}, &Complaint{}, &Warden{})
}

type User struct {
	UserID   uint   `gorm:"primaryKey"`
	FullName string `gorm:"not null"`
	Email    string `gorm:"not null"`
	Phone    string `gorm:"not null"`
	Password string `gorm:"not null"`
	Type     string `gorm:"not null"`
	BlockID  uint
	USN      string
	Room     string
}

type Block struct {
	BlockId   uint `gorm:"primaryKey;foreignKey:BlockId;references:User(BlockID);onDelete:CASCADE"`
	BlockName string
}

type Student struct {
	StudentID uint `gorm:"primaryKey;foreignKey:StudentId;references:User(UserID);onDelete:CASCADE"`
	FullName  string
	Email     string
	Phone     string
	USN       string
	BlockID   uint `gorm:"foreignKey:BlockID;references:User(BlockID);onDelete:CASCADE"`
	Room      string
}

type Warden struct {
	Warden_Id uint `gorm:"primaryKey;foreignKey:Warden_Id;references:User(UserID);onDelete:CASCADE"`
	BlockID   uint `gorm:"foreignKey:BlockID;references:Block(BlockID);onDelete:CASCADE"`
}

type Category struct {
	CategoryID   uint   `gorm:"primaryKey"`
	CategoryName string `gorm:"not null"`
}
type ComplaintType string

const (
	WIFI        ComplaintType = "WIFI"
	ELECTRICITY ComplaintType = "ELECTRICITY"
	OTHER       ComplaintType = "OTHER"
)

type Complaint struct {
	ID               uint `gorm:"primaryKey"`
	Name             string
	BlockID          uint `gorm:"foreignKey:BlockID;references:Block(BlockID);onDelete:CASCADE"`
	CategoryID       uint
	StudentID        uint `gorm:"foreignKey:StudentID;references:Student(StudentID);onDelete:CASCADE"`
	ComplaintIssues  ComplaintType
	AssignedWorkerID uint
	WardenID         uint
	Description      string
	Room             string
	IsCompleted      bool
	CreatedAt        time.Time
	AssignedAt       time.Time
}

func CloseDatabase() {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			return
		}
		sqlDB.Close()
	}
}
