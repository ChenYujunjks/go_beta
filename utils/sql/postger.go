package sql

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID   uint
	Name string
}

func PostgreSQL_Test() {
	dsn := "host=localhost user=postgres password=1532 dbname=mydatabase port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Successfully connected to the PostgreSQL database")

	// 自动迁移模式
	db.AutoMigrate(&User{})

	// 创建一个用户
	db.Create(&User{Name: "Yujun"})

	// 查询刚刚创建的用户
	var user User
	db.First(&user, 1) // 根据主键查询
	fmt.Println("User:", user.Name)
}
