/*
初始化数据库链接
*/
package config

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"time"
)

var (
	Db      *gorm.DB
	err     error
	Mysqldb string
)

func init() {
	//初始化数据库的钥匙
	name := os.Getenv("name")
	password := os.Getenv("password")
	if name == "" || password == "" {
		panic("缺少必要的环境变量 name 或 password")
	}
	Mysqldb = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/minato_sys?charset=utf8", name, password)
	//链接数据库
	Db, err = gorm.Open(mysql.Open(Mysqldb), &gorm.Config{})
	if err != nil {
		fmt.Println("mysql connect error:", err.Error())
	}
	if Db.Error != nil {
		fmt.Println("mysql connect error:", Db.Error)
	}
	//设置连接池
	sqlDB, _ := Db.DB()
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)
}
