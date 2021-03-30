package routers

import (
	"github.com/gin-gonic/gin"
	"qn/server/controllers"
)

func LoadRouter(e *gin.Engine) {
	e.GET("/upload", controllers.UpToken)
}