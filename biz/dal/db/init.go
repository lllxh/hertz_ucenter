package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"hertz_ucenter/pkg/config"
)

var DB *gorm.DB

var dsn string

func Init() {
	var err error
	// 打印一下dsn
	dsn = config.GetDsn()
	fmt.Println(dsn)
	DB, err = gorm.Open(mysql.Open(config.GetDsn()),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		})
	if err != nil {
		panic(err)
	}
}
