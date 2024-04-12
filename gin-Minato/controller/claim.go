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
	ClaimId        string `json:"claim_id"`
	ClaimUser      string `json:"claim_user"`
	ClaimInsurance int    `json:"claim_insurance"`
	Callnumber     string `json:"callnumber"`
	CarId          string `json:"car_id"`
	Region         string `json:"region"`
	Createtime     string `json:"createtime"`
}
type ClaimId struct {
	ClaimController
	ClaimId      string  `json:"claim_id"`
	Compensation float64 `json:"compensation"`
}

type AddressUpdate struct {
	ClaimController
	ClaimId string `json:"claim_id"`
	Address string `json:"address"`
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
		newclaim.ClaimId,
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
	userid := ctx.Param("userid")
	result, code := model.ShowClaim(userid)
	if code != 200 {
		config.ReturnFalse(ctx, code, "查询失败")
	} else {
		config.ReturnSuccess(ctx, code, "查询成功", result, 1)
	}
}

func (c ClaimController) ShowAllClaims(ctx *gin.Context) {
	result, code := model.ShowAllClaim()
	if code != 200 {
		config.ReturnFalse(ctx, code, "查询失败")
	} else {
		config.ReturnSuccess(ctx, code, "查询成功", result, 1)
	}
}

func (c ClaimController) ShowClaimDetail(ctx *gin.Context) {
	claimid := ctx.Param("claimid")
	result, code := model.ClaimDetail(claimid)
	if code != 200 {
		config.ReturnFalse(ctx, code, "查询失败")
	} else {
		config.ReturnSuccess(ctx, code, "查询成功", result, 1)
	}
}

func (c ClaimController) AddressUpdate(ctx *gin.Context) {
	claim := AddressUpdate{}
	err := ctx.BindJSON(&claim)
	if err != nil {
		fmt.Println("申报模块数据绑定失效：", err)
		config.ReturnFalse(ctx, 400, "数据绑定失败")
		return
	}
	code := model.AddressUpdate(claim.ClaimId, claim.Address)
	if code == 200 {
		config.ReturnSuccess(ctx, code, "进度加一", claim, 1)
	} else {
		config.ReturnFalse(ctx, code, "进度添加失败")
	}
}
