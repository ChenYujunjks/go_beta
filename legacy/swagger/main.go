// @title Swagger 示例项目
// @version 1.0
// @description 这是一个使用 Gin + Swaggo 的示例项目
// @host localhost:8080
// @BasePath /api/v1

package main

import (
	_ "my-api-project/docs"
	"my-api-project/handlers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
    r := gin.Default()

    v1 := r.Group("/api/v1")
    {
        v1.GET("/user/:id", handlers.GetUser)
    }

    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    r.Run(":8080")
}
