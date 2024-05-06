package middleware

import (
	"fmt"
	"gin-Minato/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"os"
	"time"
)

// 弄个简单一点的token
type Konoha struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// 从env文件中读取密钥
var salt = os.Getenv("salt")
var mySignkey = []byte(salt)

func Tokencreate(username string) string {
	c := Konoha{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,
			//设置过期时间在一天后
			ExpiresAt: time.Now().Unix() + 60*60*24,
			Issuer:    "Minato",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	konohaToken, err := token.SignedString(mySignkey)
	if err != nil {
		fmt.Println("token加密出错", err.Error())
	}
	//fmt.Println(konohaToken)
	return konohaToken
}

func ParseToken(konohaToken string) int {
	token, err := jwt.ParseWithClaims(konohaToken, &Konoha{}, func(token *jwt.Token) (interface{}, error) {
		return mySignkey, nil
	})
	if err != nil {
		fmt.Println(err.Error())
		return 1001
	}
	if !token.Valid {
		fmt.Println("token无效")
		return 1002 // 返回无效token的错误码
	}
	claims, ok := token.Claims.(*Konoha)
	if !ok {
		fmt.Println("token claims类型错误")
		return 1003 // 返回token claims类型错误的错误码
	}
	if claims.Username != "Minato" {
		return 1004
	}
	return 200
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenStr := c.Request.Header.Get("konohaToken")
		if tokenStr == "" {
			config.ReturnFalse(c, 1000, "token缺失")
		}
		status := ParseToken(tokenStr)
		switch status {
		case 1000:
			config.ReturnFalse(c, 1001, "token过期")
		case 1001:
			config.ReturnFalse(c, 1002, "token签发人不正确")
		case 1002:
			config.ReturnFalse(c, 1003, "token claims类型错误")
		case 1003:
			config.ReturnFalse(c, 1004, "token无效")
		case 200:
			// token验证通过
			fmt.Println(200)
			return
		default:
			// 处理其他未知错误
			config.ReturnFalse(c, 1999, "未知错误")
		}
	}
}
