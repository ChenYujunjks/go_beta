package controllers

import (
	"jwt-go/db"
	"jwt-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowUsersPage(c *gin.Context) {
	var users []models.User // 假设你的模型叫 User
	if err := db.DB.Find(&users).Error; err != nil {
		c.String(http.StatusInternalServerError, "数据库查询失败: %v", err)
		return
	}
	c.HTML(http.StatusOK, "users.html", gin.H{
		"users": users,
	})
}

func ShowHomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func ShowLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func ShowRegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{})
}
