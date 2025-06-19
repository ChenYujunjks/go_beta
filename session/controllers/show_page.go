package controllers

import (
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

func GetUpload(c *gin.Context) {
	c.HTML(http.StatusOK, "upload.html", nil)
}

func PostAdd(c *gin.Context) {
	n1 := c.PostForm("number1")
	n2 := c.PostForm("number2")
	num1, err1 := strconv.ParseFloat(n1, 64)
	num2, err2 := strconv.ParseFloat(n2, 64)

	data := gin.H{}
	if err1 != nil || err2 != nil {
		data["Result"] = "无效输入，请填写数字"
	} else {
		data["Result"] = num1 + num2
	}
	c.HTML(http.StatusOK, "index.html", data)
}
