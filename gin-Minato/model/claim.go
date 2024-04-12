package model

import (
	"fmt"
	"gorm.io/gorm"
)

type Claim struct {
	ClaimId        string  `gorm:"column:claim_id;"`
	ClaimUser      string  `gorm:"column:claim_user"`
	ClaimInsurance int     `gorm:"column:claim_insurance"`
	Callnumber     string  `gorm:"column:callnumber"`
	CarId          string  `gorm:"column:car_id"`
	Region         string  `gorm:"column:region"`
	Compensation   float64 `gorm:"column:compensation"`
	Isvisual       int     `gorm:"column:isvisual;default:1"`
	Status         int     `gorm:"column:status;default:1"`
	Address        string  `gorm:"column:address"`
	Createtime     string  `gorm:"column:createtime"`
}

func (Claim) TableName() string {
	return "claim"
}

func FindClaimById(claimid int) (Claim, error) {

	var claim Claim
	err := DB.Where("claim_id=?", claimid).First(&claim).Error
	return claim, err
}

func AddClaim(claimid string, claimuser string, claiminsurance int, Callnumber string, carid string, region string, createtime string) int {
	//_, err := FindClaimById(claimid)
	//if err != nil && err.Error() != "record not found" {
	//	fmt.Println("注册查询出错：", err.Error())
	//}
	//if err.Error() == "record not found" {
	var newClaim = Claim{
		ClaimId:        claimid,
		ClaimUser:      claimuser,
		ClaimInsurance: claiminsurance,
		Callnumber:     Callnumber,
		CarId:          carid,
		Region:         region,
		Createtime:     createtime,
	}
	err := DB.Create(&newClaim)
	if err.Error != nil {
		fmt.Println("创建数据出错：", err.Error)
		return 2001
	}
	return 200
}

// 修改claim的status状态
func UpdateStatus(claimid string, compensation float64) int {
	var claim Claim
	err := DB.Where("claim_id=?", claimid).First(&claim).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("claim记录未查询到:", err.Error())
			return 2003
		} else {
			fmt.Println("claim记录查询出错:", err.Error())
			return 2999
		}
	} else {
		claim.Status++
		claim.Compensation = compensation
		fmt.Println(claim)
		err := DB.Model(&Claim{}).Where("claim_id = ?", claimid).Updates(&claim).Error
		if err != nil {
			fmt.Println("数据库修改：", err.Error)
			return 2004
		} else {
			return 200
		}
	}
}

func ShowClaim(userid string) ([]map[string]interface{}, int) {
	var claims []map[string]interface{}
	//err := DB.Table("claim").Preload("User").Where("claim_user = ?", userid).Find(&claims).Error
	err := DB.Model(&Claim{}).
		Where("claim_user=?", userid).
		Select("claim.claim_id",
			"user.role",
			" user.user_name",
			"claim.callnumber",
			"claim.claim_insurance",
			"claim.car_id",
			"claim.region",
			"claim.createtime",
			"claim.status",
			"claim.address",
			"claim.compensation").
		Joins("left join user on claim.claim_user=user.user_id").Scan(&claims).Error

	// 修改 claims 切片中的键名
	for _, claim := range claims {
		claim["claimid"] = claim["claim_id"]
		delete(claim, "claim_id")
		claim["username"] = claim["user_name"]
		delete(claim, "user_name")
		claim["insuranceid"] = claim["claim_insurance"]
		delete(claim, "claim_insurance")
		claim["carid"] = claim["car_id"]
		delete(claim, "car_id")
		claim["date"] = claim["createtime"]
		delete(claim, "createtime")
	}
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return []map[string]interface{}{}, 200
		} else {
			fmt.Println("申报查询出错：", err.Error())
			return []map[string]interface{}{}, 3003
		}
	} else {
		return claims, 200
	}
}

// 展示所有代办申报
func ShowAllClaim() ([]map[string]interface{}, int) {
	var claims []map[string]interface{}
	err := DB.Model(&Claim{}).
		Where("status=?", 1).
		Where("isvisual=?", 1).
		Select("claim.claim_id",
			"user.role",
			" user.user_name",
			"claim.callnumber",
			"claim.claim_insurance",
			"claim.car_id",
			"claim.region",
			"claim.createtime",
			"claim.status",
			"claim.compensation").
		Joins("left join user on claim.claim_user=user.user_id").Scan(&claims).Error

	// 修改 claims 切片中的键名
	for _, claim := range claims {
		claim["claimid"] = claim["claim_id"]
		delete(claim, "claim_id")
		claim["username"] = claim["user_name"]
		delete(claim, "user_name")
		claim["insuranceid"] = claim["claim_insurance"]
		delete(claim, "claim_insurance")
		claim["carid"] = claim["car_id"]
		delete(claim, "car_id")
		claim["date"] = claim["createtime"]
		delete(claim, "createtime")
	}
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return []map[string]interface{}{}, 200
		} else {
			fmt.Println("申报查询出错：", err.Error())
			return []map[string]interface{}{}, 3003
		}
	} else {
		fmt.Println("claims is", claims)
		return claims, 200
	}
}

// 展示申报细节
func ClaimDetail(claimid string) (map[string]interface{}, int) {
	var claim map[string]interface{}
	err := DB.Model(&Claim{}).
		Where("claim_id=?", claimid).
		Select("claim.claim_id",
			"user.role",
			" user.user_name",
			" user.publicKey",
			"claim.callnumber",
			"claim.claim_insurance",
			"claim.car_id",
			"claim.region",
			"claim.createtime",
			"claim.status",
			"claim.compensation").
		Joins("left join user on claim.claim_user=user.user_id").Scan(&claim).Error
	// 修改 claims 切片中的键名
	claim["claimid"] = claim["claim_id"]
	delete(claim, "claim_id")
	claim["username"] = claim["user_name"]
	delete(claim, "user_name")
	claim["insuranceid"] = claim["claim_insurance"]
	delete(claim, "claim_insurance")
	claim["carid"] = claim["car_id"]
	delete(claim, "car_id")
	claim["date"] = claim["createtime"]
	delete(claim, "createtime")
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return map[string]interface{}{}, 200
		} else {
			fmt.Println("申报查询出错：", err.Error())
			return map[string]interface{}{}, 3003
		}
	} else {
		fmt.Println("claim is", claim)
		return claim, 200
	}
}

func AddressUpdate(claimid string, address string) int {
	var claim Claim
	err := DB.Where("claim_id=?", claimid).First(&claim).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("claim记录未查询到:", err.Error())
			return 2003
		} else {
			fmt.Println("claim记录查询出错:", err.Error())
			return 2999
		}
	} else {
		claim.Status = 4
		claim.Address = address
		fmt.Println(claim)
		err := DB.Model(&Claim{}).Where("claim_id = ?", claimid).Updates(&claim).Error
		if err != nil {
			fmt.Println("数据库修改：", err.Error)
			return 2004
		} else {
			return 200
		}
	}
}
