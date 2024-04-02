package model

import "time"

type Insurance struct {
	InsuranceId    int       `gorm:"column:insurance_id;"`
	InsuranceName  string    `gorm:"column:insurance_name"`
	Thumbnail      string    `gorm:"column:thumbnail"`
	Content        string    `gorm:"column:content"`
	InsurancePrice float64   `gorm:"column:insurance_price"`
	PurchaseCount  int       `gorm:"column:purchase_count"`
	Createtime     time.Time `gorm:"column:createtime"`
}

func (Insurance) TableName() string {
	return "insurance"
}
