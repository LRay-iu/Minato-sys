package model

import "time"

type Purchase struct {
	PurchaseId        int       `gorm:"column:purchase_id;"`
	PurchaseUser      string    `gorm:"column:purchase_user"`
	PurchaseInsurance string    `gorm:"column:purchase_insurance"`
	Address           string    `gorm:"column:address"`
	Createtime        time.Time `gorm:"column:createtime"`
}

func (Purchase) TableName() string {
	return "purchase"
}
