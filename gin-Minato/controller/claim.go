package controller

import (
	"fmt"
	"gin-Minato/config"
	"gin-Minato/model"
	"github.com/gin-gonic/gin"
)

type ClaimController struct {
}
type NewClaim struct {
	ClaimController
	ClaimUser      string `json:"claim_user"`
	ClaimInsurance int    `json:"claim_insurance"`
	Callnumber     string `json:"callnumber"`
	CarId          string `json:"car_id"`
	Region         string `json:"region"`
	Createtime     string `json:"createtime"`
}
type ClaimId struct {
	ClaimController
	ClaimId      int     `json:"claim_id"`
	Compensation float64 `json:"compensation"`
}

func (c ClaimController) AddClaim(ctx *gin.Context) {
	newclaim := NewClaim{}
	err := ctx.BindJSON(&newclaim)
	if err != nil {
		fmt.Println("申报模块数据绑定失效：", err)
		config.ReturnFalse(ctx, 400, "数据绑定失败")
		return
	}
	code := model.AddClaim(
		config.EncryMd5(newclaim.ClaimUser),
		newclaim.ClaimInsurance,
		newclaim.Callnumber,
		newclaim.CarId,
		newclaim.Region,
		newclaim.Createtime)
	if code == 200 {
		config.ReturnSuccess(ctx, code, "创建成功", newclaim, 1)
	} else {
		config.ReturnFalse(ctx, code, "创建数据集失败")
		return
	}
}

func (c ClaimController) UpdateClaim(ctx *gin.Context) {
	claim := ClaimId{}
	err := ctx.BindJSON(&claim)
	if err != nil {
		fmt.Println("申报模块数据绑定失效：", err)
		config.ReturnFalse(ctx, 400, "数据绑定失败")
		return
	}
	code := model.UpdateStatus(claim.ClaimId, claim.Compensation)
	if code == 200 {
		config.ReturnSuccess(ctx, code, "进度加一", claim, 1)
	} else {
		config.ReturnFalse(ctx, code, "进度添加失败")
	}
}

func (c ClaimController) ShowClaim(ctx *gin.Context) {
	result, code := model.ShowClaim()
	if code != 200 {
		config.ReturnFalse(ctx, code, "查询失败")
	} else {
		config.ReturnSuccess(ctx, code, "查询成功", result, 1)
	}
}
