package database

import (
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

	// Migrate the schema
	DB.AutoMigrate(&User{})
}

type User struct {
	UserID   uint   `gorm:"primaryKey"`
	FullName string `gorm:"not null"`
	Email    string `gorm:"not null"`
	Phone    string `gorm:"not null"`
	Password string `gorm:"not null"`
	Type     string `gorm:"not null"`
}

func CloseDatabase() {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			// Handle error
			return
		}
		sqlDB.Close()
	}
}
