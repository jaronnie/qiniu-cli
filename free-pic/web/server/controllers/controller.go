package controllers

import (
	"github.com/gin-gonic/gin"
	"qn/server/service"
)


func UpToken(c *gin.Context) {
	token := service.GenerateToken()
	c.JSON(200, gin.H{
		"token":token,
	})
}



