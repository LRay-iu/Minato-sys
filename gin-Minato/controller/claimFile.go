package controller

import (
	"fmt"
	"gin-Minato/config"
	"github.com/gin-gonic/gin"
)

type FileController struct {
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
