package config

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
)

//为了解决循环依赖的问题，我将控制层中的common.go移到了此处,未来或许会对这个进行优化，也有可能不会
/*
模型->中间层->控制层->模型
*/
type JsonStruct struct {
	Code  int         `json:"code"`  //返回状态码
	Msg   interface{} `json:"msg"`   //返回的提示语
	Data  interface{} `json:"data"`  //返回数据
	Count interface{} `json:"count"` //返回条数
}

type JsonErrStruct struct {
	Code int         `json:"code"` //返回状态码
	Msg  interface{} `json:"msg"`  //返回的提示语
}

func ReturnSuccess(c *gin.Context, code int, msg interface{}, data interface{}, count int64) {
	json := &JsonStruct{Code: code, Msg: msg, Data: data, Count: count}
	c.JSON(200, json)
}
func ReturnFalse(c *gin.Context, code int, msg interface{}) {
	json := &JsonErrStruct{Code: code, Msg: msg}
	c.JSON(200, json)
}
func EncryMd5(s string) string {
	ctx := md5.New()
	ctx.Write([]byte(s))
	return hex.EncodeToString(ctx.Sum(nil))
}
