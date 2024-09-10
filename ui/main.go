package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID    uint `gorm:"primarykey"`
	Name  string
	Email string
}

func createDatabase() (*gorm.DB, error) {
	// 创建 SQLite 数据库连接
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 自动迁移
	db.AutoMigrate(&User{})

	// 插入种子数据
	seedData := []User{
		{Name: "Bruce", Email: "bruce@example.com"},
		{Name: "Alice", Email: "alice@example.com"},
	}
	for _, user := range seedData {
		db.FirstOrCreate(&user, User{Name: user.Name})
	}

	return db, nil
}

func main() {
	var db *gorm.DB

	// 创建 Fyne 应用程序
	myApp := app.New()
	myWindow := myApp.NewWindow("CRUD Example")

	// 输入框组件
	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Enter Name")

	emailEntry := widget.NewEntry()
	emailEntry.SetPlaceHolder("Enter Email")

	resultLabel := widget.NewLabel("")

	// 创建数据库按钮
	createDbBtn := widget.NewButton("Create Database", func() {
		var err error
		db, err = createDatabase()
		if err != nil {
			resultLabel.SetText("Failed to create database")
		} else {
			resultLabel.SetText("Database created with seed data")
		}
	})

	// CRUD 按钮
	createBtn := widget.NewButton("Create", func() {
		if db != nil {
			user := User{Name: nameEntry.Text, Email: emailEntry.Text}
			db.Create(&user)
			resultLabel.SetText(fmt.Sprintf("User Created: %s", user.Name))
		} else {
			resultLabel.SetText("Database not created")
		}
	})

	readBtn := widget.NewButton("Read", func() {
		if db != nil {
			var user User
			db.First(&user, "name = ?", nameEntry.Text)
			if user.ID != 0 {
				resultLabel.SetText(fmt.Sprintf("Found User: %s, Email: %s", user.Name, user.Email))
			} else {
				resultLabel.SetText("User Not Found")
			}
		} else {
			resultLabel.SetText("Database not created")
		}
	})

	updateBtn := widget.NewButton("Update", func() {
		if db != nil {
			var user User
			db.First(&user, "name = ?", nameEntry.Text)
			if user.ID != 0 {
				db.Model(&user).Update("Email", emailEntry.Text)
				resultLabel.SetText(fmt.Sprintf("Updated User: %s's Email", user.Name))
			} else {
				resultLabel.SetText("User Not Found")
			}
		} else {
			resultLabel.SetText("Database not created")
		}
	})

	deleteBtn := widget.NewButton("Delete", func() {
		if db != nil {
			var user User
			db.First(&user, "name = ?", nameEntry.Text)
			if user.ID != 0 {
				db.Delete(&user)
				resultLabel.SetText(fmt.Sprintf("Deleted User: %s", user.Name))
			} else {
				resultLabel.SetText("User Not Found")
			}
		} else {
			resultLabel.SetText("Database not created")
		}
	})

	// 布局
	content := container.NewVBox(
		widget.NewLabel("User CRUD Operations"),
		nameEntry,
		emailEntry,
		createDbBtn,
		createBtn,
		readBtn,
		updateBtn,
		deleteBtn,
		resultLabel,
	)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(400, 400))
	myWindow.ShowAndRun()
}
