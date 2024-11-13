package router

import (
	"bluebell/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// 注册路由
	r.POST("/signup", controller.SignUpHandler)

	return r
}
