package models

import "gorm.io/gorm"

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string
}

// models/migrate.go
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
