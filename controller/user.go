package controller

import (
	"bluebell/logic"
	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	// 参数校验

	logic.SignUp()

	c.JSON(200, "OK!")
}
