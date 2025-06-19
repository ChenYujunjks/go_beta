package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func PostAdd(c *gin.Context) {
	var req struct {
		Number1 float64 `json:"number1"`
		Number2 float64 `json:"number2"`
	  }
	  if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效 JSON"})
		return
	  }
	  c.JSON(http.StatusOK, gin.H{
		"result": req.Number1 + req.Number2,
	  })
}
