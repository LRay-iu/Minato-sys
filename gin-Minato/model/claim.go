package model

import (
	"time"
)

type Claim struct {
	ClaimId        int       `gorm:"column:claim_id;"`
	ClaimUser      string    `gorm:"column:claim_user"`
	ClaimInsurance int       `gorm:"column:claim_insurance"`
	Callnumber     string    `gorm:"column:callnumber"`
	CarId          int       `gorm:"column:car_id"`
	Region         string    `gorm:"column:region"`
	Compensation   float64   `gorm:"column:compensation"`
	Isvisual       int       `gorm:"column:isvisual"`
	Status         int       `gorm:"column:status"`
	Address        string    `gorm:"column:address"`
	Createtime     time.Time `gorm:"column:createtime"`
}

func (Claim) TableName() string {
	return "claim"
}
