package controller

import (
	"fmt"
	"gin-Minato/config"
	"gin-Minato/model"
	"github.com/gin-gonic/gin"
)

type PurchaseController struct {
}

type PurchaseCreate struct {
	PurchaseController
	PurchaseUser      string `json:"purchase_user"`
	PurchaseInsurance int    `json:"purchase_insurance"`
	Address           string `json:"address"`
}

func (p PurchaseController) AddPurchase(ctx *gin.Context) {
	var newPerchase PurchaseCreate
	err := ctx.BindJSON(&newPerchase)
	if err != nil {
		fmt.Println("注册模块数据绑定失效：", err.Error())
		config.ReturnFalse(ctx, 400, "数据绑定失败")
		return
	}
	code := model.AddPurchase(config.EncryMd5(newPerchase.PurchaseUser), newPerchase.PurchaseInsurance, newPerchase.Address)
	if code == 200 {
		config.ReturnSuccess(ctx, code, "购买成功", newPerchase, 1)
	} else {
		config.ReturnFalse(ctx, code, "购买失败！")
	}
}
