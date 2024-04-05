package model

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Insurance struct {
	InsuranceId    int       `gorm:"column:insurance_id;"`
	InsuranceName  string    `gorm:"column:insurance_name"`
	Thumbnail      string    `gorm:"column:thumbnail"`
	Content        string    `gorm:"column:content"`
	InsurancePrice float64   `gorm:"column:insurance_price"`
	PurchaseCount  int       `gorm:"column:purchase_count"`
	Createtime     time.Time `gorm:"type:datetime;column:createtime"`
}

func (Insurance) TableName() string {
	return "insurance"
}

func Showinsurance() ([]Insurance, int) {
	var insurances []Insurance
	err := DB.Find(&insurances).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return []Insurance{}, 200
		} else {
			fmt.Println("保险查询出错：", err.Error())
			return []Insurance{}, 3003
		}
	} else {
		return insurances, 200
	}

}
