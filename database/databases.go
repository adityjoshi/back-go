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
	}
	fmt.Print("database connected successfully ⚡️")

	// Migrate the schema
	DB.AutoMigrate(&User{}, &Block{}, &Student{}, &Category{}, &Complaint{})
}

type User struct {
	UserID   uint   `gorm:"primaryKey"`
	FullName string `gorm:"not null"`
	Email    string `gorm:"not null"`
	Phone    string `gorm:"not null"`
	Password string `gorm:"not null"`
	Type     string `gorm:"not null"`
}

type Block struct {
	BlockId   uint   `gorm:"primaryKey"`
	BlockName string `gorm:"not null"`
}

type Student struct {
	StudentId uint `gorm:"primaryKey"`
	BlockId   uint `gorm:"foreignKey:BlockID;references:Block(BlockID);onDelete:CASCADE"`
	USN       string
	Room      string
}

type Category struct {
	CategoryID   uint   `gorm:"primaryKey"`
	CategoryName string `gorm:"not null"`
}

type Complaint struct {
	ID               uint `gorm:"primaryKey"`
	Name             string
	BlockID          uint
	CategoryID       uint
	StudentID        uint
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
