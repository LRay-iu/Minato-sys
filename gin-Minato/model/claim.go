package model

import (
	"fmt"
	"gorm.io/gorm"
)

type Claim struct {
	ClaimId        int     `gorm:"column:claim_id;"`
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

func AddClaim(claimuser string, claiminsurance int, Callnumber string, carid string, region string, createtime string) int {
	//_, err := FindClaimById(claimid)
	//if err != nil && err.Error() != "record not found" {
	//	fmt.Println("注册查询出错：", err.Error())
	//}
	//if err.Error() == "record not found" {
	var newClaim = Claim{
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
func UpdateStatus(claimid int, compensation float64) int {
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

func ShowClaim() ([]Claim, int) {
	var claims []Claim
	err := DB.Find(&claims).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return []Claim{}, 200
		} else {
			fmt.Println("保险查询出错：", err.Error())
			return []Claim{}, 3003
		}
	} else {
		return claims, 200
	}
}
