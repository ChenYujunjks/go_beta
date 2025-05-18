package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// User 定义用户结构
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GetUser 获取单个用户信息
// @Summary 获取用户
// @Description 通过用户 ID 获取用户信息
// @Tags 用户
// @Accept  json
// @Produce  json
// @Param id path int true "用户ID"
// @Success 200 {object} User
// @Router /user/{id} [get]
func GetUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, User{
		ID:   1,
		Name: "Yitan " + id,
	})
}
