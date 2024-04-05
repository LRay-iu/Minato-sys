package model

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Purchase struct {
	PurchaseId        int       `gorm:"column:purchase_id;"`
	PurchaseUser      string    `gorm:"column:purchase_user"`
	PurchaseInsurance int       `gorm:"column:purchase_insurance"`
	Address           string    `gorm:"column:address"`
	Createtime        time.Time `gorm:"column:createtime"`
}

func (Purchase) TableName() string {
	return "purchase"
}

// 通过用户id和保单名来判断是否买过
func FindPurchaseByUser(userId string, insuranceId int) (Purchase, bool) {
	var purchase Purchase
	err := DB.Where("purchase_user=? AND purchase_insurance=?", userId, insuranceId).First(&purchase).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return Purchase{}, false
		} else {
			fmt.Println("查询保单购买情况出错：", err.Error())
			return Purchase{}, false
		}
	} else {
		return purchase, true
	}
}

// 创建订单
func AddPurchase(purchaseUser string, purchaseInsurance int, address string) int {
	_, result := FindPurchaseByUser(purchaseUser, purchaseInsurance)
	if result == false {
		newPurchase := Purchase{
			PurchaseUser:      purchaseUser,
			PurchaseInsurance: purchaseInsurance,
			Address:           address,
			Createtime:        time.Now().Truncate(24 * time.Hour),
		}
		err := DB.Create(&newPurchase).Error
		if err != nil {
			fmt.Println("创建数据出错：", err.Error())
			return 3001
		} else {
			return 200
		}
	} else {
		return 3002
	}
}
