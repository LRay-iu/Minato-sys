package controller

import (
	"fmt"
	"gin-Minato/config"
	"gin-Minato/model"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserId     string `json:"user_id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Callnumber string `json:"callnumber"`
	PublicKey  string `json:"publicKey"`
	Role       string `json:"role"`
}

// 返回给前端的，总不能密码也返回去吧
type UserReturn struct {
	UserId      string `json:"user_id"`
	Username    string `json:"username"`
	PublicKey   string `json:"publicKey"`
	KonohaToken string `json:"konohaToken"`
}

func (u UserController) Register(c *gin.Context) {
	user := UserController{}
	err := c.BindJSON(&user)
	if err != nil {
		fmt.Println("注册模块数据绑定失效：", err)
		config.ReturnFalse(c, 400, "数据绑定失败")
	}
	msg := model.Register(config.EncryMd5(user.UserId), user.Username, config.EncryMd5(user.Password), user.Callnumber, user.PublicKey, user.Role)
	if msg == 200 {
		config.ReturnSuccess(c, msg, "注册成功", user, 1)
	} else {
		config.ReturnFalse(c, msg, "创建数据库失败")
	}
}

func (u UserController) Login(c *gin.Context) {
	loginer := UserController{}
	err := c.BindJSON(&loginer)
	if err != nil {
		fmt.Println("登录模块数据绑定失效：", err)
		config.ReturnFalse(c, 400, "数据绑定失败")
	}
	if loginer.UserId == "" || loginer.Password == "" {
		config.ReturnFalse(c, 1009, "账号和密码缺失")
	}
	code, user, konohaToken := model.Login(config.EncryMd5(loginer.UserId), config.EncryMd5(loginer.Password))
	switch code {
	case 1005:
		config.ReturnFalse(c, 1005, "用户不存在")
	case 1006:
		config.ReturnFalse(c, 1006, "id或密码不正确")
	//case 1999:
	//	config.ReturnFalse(c, 1999, "其他错误")
	case 200:
		userReturn := UserReturn{
			UserId:      user.UserId,
			Username:    user.Username,
			PublicKey:   user.PublicKey,
			KonohaToken: konohaToken,
		}
		config.ReturnSuccess(c, 200, "登陆成功", userReturn, 1)
	}
}
