package model

import "time"

type Claimfile struct {
	Fileid        int       `gorm:"column:file_id;"`
	Filename      string    `gorm:"column:file_name"`
	Filepath      string    `gorm:"column:file_path"`
	Fileaccording int       `gorm:"column:file_according"`
	Createtime    time.Time `gorm:"column:createtime"`
}

func (Claimfile) TableName() string {
	return "claimfile"
}
