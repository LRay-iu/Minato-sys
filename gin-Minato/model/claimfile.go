package model

import "fmt"

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
