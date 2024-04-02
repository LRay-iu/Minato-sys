package router

import (
	"gin-Minato/controller"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.POST("/login", controller.UserController{}.Login)
	r.POST("/register", controller.UserController{}.Register)
	return r
}
