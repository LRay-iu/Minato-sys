package controller

import (
	"encoding/base64"
	"fmt"
	"gin-Minato/config"
	"gin-Minato/model"
	"gin-Minato/pkg"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"path/filepath"
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
	for _, value := range newfile.Filename {
		fmt.Println("文件名：", value)
		newName := pkg.GenerateRandomString() + ".jpg"
		oldPath := filepath.Join("uploads", value)
		newPath := filepath.Join("uploads", newName)
		if _, err := os.Stat(oldPath); err == nil {
			err = os.Rename(oldPath, newPath)
			if err == nil {
				fmt.Println("旧", value, "改为新", newName)
				code := model.AddClaimfile(newName, newPath, newfile.FileAccording, newfile.Createtime)
				if code == 200 {
					fmt.Println("成功改名并保存至数据库！")
				} else {
					fmt.Println("保存至数据库失败！")
				}
			} else {
				fmt.Println("重命名文件失败:", err)
			}
		} else {
			fmt.Println("文件不存在:", err)
		}
	}
}

func (f FileController) Getfile(ctx *gin.Context) {
	claimid := ctx.Param("claimid")
	result, code := model.ClaimFileGet(claimid)
	if code != 200 {
		config.ReturnFalse(ctx, code, "查询失败")
	} else {
		for key, value := range result {
			//fmt.Println("key", key)
			//fmt.Println("value", value["file_path"])
			imageData, err := ioutil.ReadFile(value["file_path"].(string))
			if err != nil {
				fmt.Println("文件读取失败", err)
			}
			base64Image := base64.StdEncoding.EncodeToString(imageData)
			value["file"] = base64Image
			result[key] = value
		}
		config.ReturnSuccess(ctx, code, "查询成功", result, 1)
	}
}
