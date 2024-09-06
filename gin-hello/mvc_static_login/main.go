package main

import (
	"gin/intro/mvc_static_login/controllers"
	"gin/intro/mvc_static_login/models"
	"log"

	"github.com/gin-gonic/gin"

	//选择轻量级数据库而不是"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// SQLite连接配置
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// 迁移数据库
	db.AutoMigrate(&models.User{})

	r := gin.Default()
	//首先加载templates目录下面的所有模版文件，模版文件扩展名随意
	r.LoadHTMLGlob("views/pages/*")
	//  /assets/images/1.jpg 这个url文件，存储在/public/images/1.jpg
	r.Static("/assets", "views/static")
	/* 配置CORS中间件
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8444"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	*/
	r.POST("/login", controllers.Login(db))
	r.POST("/register", controllers.Register(db))
	r.GET("/users", controllers.GetUsers(db))
	r.GET("/login", controllers.ShowLoginPage)
	r.GET("/register", controllers.ShowRegisterPage)
	r.GET("/", controllers.ShowIndexPage)
	r.GET("/react", controllers.ShowReactPage)
	r.POST("/add", controllers.AddRoute())

	r.Run(":8080")
}
