package models

import "gorm.io/gorm"

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
