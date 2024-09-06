package controllers

import (
	"fmt"
	"gin/intro/mvc_static_login/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		var requestUser models.User

		if err := c.ShouldBindJSON(&requestUser); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
				"code":  1,
				"msg":   "请求参数错误",
			})
			return
		}

		if err := db.Where("username = ? AND password = ?", requestUser.Username, requestUser.Password).First(&user).Error; err != nil {
			log.Println(err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "欢迎, " + user.Username})
	}
}

func Register(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestUser models.User

		if err := c.ShouldBindJSON(&requestUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Create(&requestUser).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "注册失败"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "注册成功"})
	}
}

func AddRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		var json struct {
			Number1 string `json:"number1"`
			Number2 string `json:"number2"`
		}

		fmt.Println("Received request at /add")
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
			return
		}

		fmt.Printf("Received JSON: %+v\n", json)

		number1 := json.Number1
		number2 := json.Number2

		if number1 == "" || number2 == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "number1 and number2 are required"})
			return
		}

		num1, err1 := strconv.ParseFloat(number1, 64)
		num2, err2 := strconv.ParseFloat(number2, 64)
		if err1 != nil || err2 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid number format"})
			return
		}

		result := num1 + num2
		c.JSON(http.StatusOK, gin.H{"result": result})
	}
}

func GetUsers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []models.User
		if err := db.Find(&users).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取用户列表"})
			return
		}
		c.HTML(http.StatusOK, "users.html", gin.H{"users": users})
	}
}

func ShowLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func ShowRegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func ShowIndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Yujun Chen's Personal Website",
	})
}

func ShowReactPage(c *gin.Context) {
	c.HTML(http.StatusOK, "react.html", nil)
}
