package main

import (
	"github.com/gin-gonic/gin"

	"jwt-go/controllers"
	"jwt-go/db"
	"jwt-go/middlewares"
	"jwt-go/models"
)

func main() {
	r := gin.Default()

	db.InitDB()
	models.AutoMigrate(db.DB)
	r.LoadHTMLGlob("templates/*")

	// 中间件
	r.Use(middlewares.InitSession())

	r.POST("/api/register", controllers.Register)
	r.POST("/api/login", controllers.Login)
	r.POST("/api/logout", controllers.Logout)
	r.GET("/api/me", controllers.GetMe)

	// 页面路由
	r.GET("/", middlewares.AuthRequired(), controllers.ShowHomePage)
	r.GET("/login", controllers.ShowLoginPage)
	r.GET("/register", controllers.ShowRegisterPage)
	r.GET("/admin", controllers.ShowUsersPage)

	r.GET("/slice-any", middlewares.AuthRequired(), controllers.GetSliceAny)
	r.Run(":8080")
}
