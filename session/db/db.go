package db

import (
	"jwt-go/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// ✅ 自动迁移模型
	if err := models.AutoMigrate(DB); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}
