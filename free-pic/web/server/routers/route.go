package routers

import (
	"github.com/gin-gonic/gin"
	"qn/server/controllers"
	"qn/server/middlewares"
)

func LoadRouter(e *gin.Engine) {
	e.GET("/upload", controllers.UpToken)
	e.GET("/api/v1/login", middlewares.JWTAuthMiddleware(), controllers.QueryAllLogin)
	e.POST("/vue-admin-template/user/login", controllers.TrueLogin)
	e.POST("/api/v1/login/add", controllers.PostLogin)
}
