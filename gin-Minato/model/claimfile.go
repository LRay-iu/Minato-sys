package model

import (
	"fmt"
	"gorm.io/gorm"
)

type Claimfile struct {
	Fileid        int    `gorm:"column:file_id;"`
	Filename      string `gorm:"column:file_name"`
	Filepath      string `gorm:"column:file_path"`
	FileAccording string `gorm:"column:file_according"`
	Createtime    string `gorm:"column:createtime"`
}

func (Claimfile) TableName() string {
	return "claimfile"
}

func AddClaimfile(filename string, filepath string, fileaccording string, createtime string) int {
	var newFile = Claimfile{
		Filename:      filename,
		Filepath:      filepath,
		FileAccording: fileaccording,
		Createtime:    createtime,
	}
	err := DB.Where(newFile).FirstOrCreate(&newFile).Error
	if err != nil {
		// 如果错误不为空，表示存在相同的内容
		fmt.Println("存在相同的内容：", err.Error)
		return 2002
	} else {
		// 如果错误为空，表示不存在相同的内容，或者成功创建了新的记录
		fmt.Println("不存在相同的内容")
		return 200
	}
}

// 展示申报中的图片
func ClaimFileGet(claimid string) ([]map[string]interface{}, int) {
	var claimfile []map[string]interface{}
	err := DB.Model(&Claimfile{}).
		Where("file_according=?", claimid).
		Select("claimfile.file_path").
		Joins("left join claim on claimfile.file_according=claim.claim_id").Scan(&claimfile).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return []map[string]interface{}{}, 200
		} else {
			fmt.Println("申报查询出错：", err.Error())
			return []map[string]interface{}{}, 3003
		}
	} else {
		fmt.Println("claim is", claimfile)
		return claimfile, 200
	}
}
