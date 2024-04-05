package router

import (
	"gin-Minato/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/login", controller.UserController{}.Login)
	r.POST("/register", controller.UserController{}.Register)
	r.POST("/addclaim", controller.ClaimController{}.AddClaim)
	r.POST("/updatestatus", controller.ClaimController{}.UpdateClaim)
	r.GET("/showclaim", controller.ClaimController{}.ShowClaim)
	r.POST("/buyinsurance", controller.PurchaseController{}.AddPurchase)
	r.GET("/showinsurance", controller.InsuranceController{}.ShowInsurance)
	r.POST("/update", controller.FileController{}.Updatefile)
	return r
}
