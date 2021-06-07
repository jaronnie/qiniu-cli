package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"qn/server/service"
)

func RetRank(c *gin.Context) {
	member := c.PostForm("key")
	ranks := service.GetRank(member)
	c.IndentedJSON(200, gin.H {
		"ranks": ranks,
	})

	fmt.Println()
}
