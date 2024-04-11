package controller

import (
	"fmt"
	"gin-Minato/config"
	"gin-Minato/model"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

// 用来绑定json
type UserRegister struct {
	UserController
	UserId     string `json:"user_id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Callnumber string `json:"callnumber"`
	PublicKey  string `json:"publicKey"`
	Role       string `json:"role"`
}
type UserLogin struct {
	UserController
	UserId   string `json:"user_id"`
	Password string `json:"password"`
}

// 返回给前端的，总不能密码也返回去吧
type UserReturn struct {
	UserId      string `json:"user_id"`
	Username    string `json:"username"`
	PublicKey   string `json:"publicKey"`
	KonohaToken string `json:"konohaToken"`
	Role        string `json:"role"`
}

func (u UserController) Register(ctx *gin.Context) {
	newuser := UserRegister{}
	err := ctx.BindJSON(&newuser)
	if err != nil {
		fmt.Println("注册模块数据绑定失效：", err)
		config.ReturnFalse(ctx, 400, "数据绑定失败")
		return
	}
	code := model.Register(config.EncryMd5(newuser.UserId),
		newuser.Username,
		config.EncryMd5(newuser.Password),
		newuser.Callnumber,
		newuser.PublicKey,
		newuser.Role)
	if code == 200 {
		config.ReturnSuccess(ctx, code, "注册成功", newuser, 1)
		return
	} else {
		config.ReturnFalse(ctx, code, "创建数据库失败")
		return
	}
}

// 登录模块功能实现
func (u UserController) Login(ctx *gin.Context) {
	loginer := UserLogin{}
	err := ctx.BindJSON(&loginer)
	if err != nil {
		fmt.Println("登录模块数据绑定失效：", err)
		config.ReturnFalse(ctx, 400, "数据绑定失败")
		return
	}
	if loginer.UserId == "" || loginer.Password == "" {
		config.ReturnFalse(ctx, 1009, "账号和密码缺失")
		return
	}
	code, user, konohaToken := model.Login(config.EncryMd5(loginer.UserId), config.EncryMd5(loginer.Password))
	switch code {
	case 1005:
		config.ReturnFalse(ctx, 1005, "用户不存在")
	case 1006:
		config.ReturnFalse(ctx, 1006, "id或密码不正确")
	//case 1999:
	//	config.ReturnFalse(ctx, 1999, "其他错误")
	case 200:
		userReturn := UserReturn{
			UserId:      user.UserId,
			Username:    user.Username,
			PublicKey:   user.PublicKey,
			KonohaToken: konohaToken,
			Role:        user.Role,
		}
		fmt.Println("key:", user.PublicKey)
		config.ReturnSuccess(ctx, 200, "登录成功", userReturn, 1)
		return
	}
}
