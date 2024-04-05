package model

import (
	"fmt"
	"gin-Minato/config"
	"gin-Minato/middleware"
	"gorm.io/gorm"
	"time"
)

type User struct {
	UserId     string    `gorm:"column:user_id;"`
	Username   string    `gorm:"column:user_name"`
	Password   string    `gorm:"column:password"`
	Callnumber string    `gorm:"column:callnumber"`
	PublicKey  string    `gorm:"column:publicKey"`
	Role       string    `gorm:"column:role"`
	Createtime time.Time `gorm:"type:datetime;column:createtime"`
}

// 定义一个数据库指针，这个包仅此声明一次，其他地方直接调用即可
var DB = config.Db

func (User) TableName() string {
	return "user"
}

// 根据userid找人
func FindUserById(userid string) (User, error) {
	var user User
	err := DB.Where("user_id=?", userid).First(&user).Error
	return user, err
}

// 注册
func Register(userid string, username string, password string, callnumber string, publickey string, role string) int {
	_, err := FindUserById(userid)
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Println("注册查询出错：", err.Error())
	}
	if err != nil && err == gorm.ErrRecordNotFound {
		var newUser = User{
			UserId:     userid,
			Username:   username,
			Password:   password,
			Callnumber: callnumber,
			PublicKey:  publickey,
			Role:       role,
			Createtime: time.Now(),
		}
		err := DB.Create(&newUser).Error
		if err != nil {
			fmt.Println("创建数据出错：", err.Error())
			return 1008
		}
		return 200
	} else {
		return 1007
	}

}

// 登录接口
func Login(userid string, password string) (int, User, string) {
	user, err := FindUserById(userid)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 没有找到符合条件的记录
			fmt.Println("用户未注册")
			return 1005, user, "nil"
		} else {
			// 处理错误，但是不要返回
			fmt.Println("登录查询出现异常：", err.Error())
		}
	}
	// 找到了符合条件的记录
	if user.Password == password {
		konohaToken := middleware.Tokencreate(user.Username)
		fmt.Println("User:", user)
		return 200, user, konohaToken
	} else {
		return 1006, user, "nil"
	}

}
