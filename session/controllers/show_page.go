package controllers

import (
	"fmt"
	"jwt-go/db"
	"jwt-go/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShowUsersPage(c *gin.Context) {
	var users []models.User
	if err := db.DB.Find(&users).Error; err != nil {
		c.String(http.StatusInternalServerError, "数据库查询失败: %v", err)
		return
	}
	fmt.Println("User count:", len(users))
	for _, u := range users {
		fmt.Printf("ID=%d, Username=%s\n", u.ID, u.Username)
	}
	c.HTML(http.StatusOK, "users.html", gin.H{
		"users": users,
	})
}

func ShowHomePage(c *gin.Context) {
	number1 := c.Query("number1")
	number2 := c.Query("number2")

	data := gin.H{}

	if number1 != "" && number2 != "" {
		num1, err1 := strconv.ParseFloat(number1, 64)
		num2, err2 := strconv.ParseFloat(number2, 64)

		if err1 == nil && err2 == nil {
			data["Result"] = num1 + num2
		} else {
			data["Result"] = "输入格式错误"
		}
	}

	c.HTML(http.StatusOK, "index.html", data)
}

func ShowLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func ShowRegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{})
}

func GetUpload(c *gin.Context) {
	c.HTML(http.StatusOK, "upload.html", nil)
}

func GetSliceAny(c *gin.Context) {
	data := []any{"string", 123, true, map[string]any{"ITIS A MAP!!": "value"}}
	c.JSON(http.StatusOK, data)
}
