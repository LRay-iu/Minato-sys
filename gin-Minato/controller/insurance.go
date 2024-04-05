package controller

import (
	"gin-Minato/config"
	"gin-Minato/model"
	"github.com/gin-gonic/gin"
)

type InsuranceController struct {
}

func (i InsuranceController) ShowInsurance(ctx *gin.Context) {
	result, code := model.Showinsurance()
	if code != 200 {
		config.ReturnFalse(ctx, code, "查询失败")
	} else {
		config.ReturnSuccess(ctx, code, "查询成功", result, 1)
	}

}
