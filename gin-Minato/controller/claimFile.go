package controller

import (
	"fmt"
	"gin-Minato/config"
	"gin-Minato/model"
	"gin-Minato/pkg"
	"github.com/gin-gonic/gin"
	"os"
)

type FileController struct {
}
type NewFile struct {
	FileController
	Filename      []string `json:"filename"`
	FileAccording string   `json:"file_according"`
	Createtime    string   `json:"createtime"`
}

func (f FileController) Updatefile(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		fmt.Sprintf("get form err: %s", err.Error())
		config.ReturnFalse(ctx, 3001, "接收文件失败")
	}

	// 保存文件到本地
	err = ctx.SaveUploadedFile(file, "uploads/"+file.Filename)
	if err != nil {
		fmt.Sprintf("upload file err: %s", err.Error())
		config.ReturnFalse(ctx, 3002, "保存文件失败")
	}
	config.ReturnSuccess(ctx, 200, "成功接受并保存文件！", file, 1) // 返回文件名
}

func (f FileController) Addfile(ctx *gin.Context) {
	newfile := NewFile{}
	err := ctx.BindJSON(&newfile)
	if err != nil {
		fmt.Println("申报模块数据绑定失效：", err)
		config.ReturnFalse(ctx, 400, "数据绑定失败")
		return
	}
	for key, value := range newfile.Filename {
		fmt.Println(key, value)
		newName := pkg.GenerateRandomString() + ".jpg"
		err := os.Rename("uploads/"+value, "uploads/"+newName)
		if err != nil {
			fmt.Println("命名文件出现错误")
		}
		code := model.AddClaimfile(newName, "uploads/"+newName, newfile.FileAccording, newfile.Createtime)
		if code == 200 {
			config.ReturnSuccess(ctx, 200, "成功改名并保存至数据库！", newName, 1)
		} else {
			config.ReturnFalse(ctx, code, "改名失败")
		}
	}
}
