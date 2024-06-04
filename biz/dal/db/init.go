package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"hertz_ucenter/pkg/config"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(config.GetDsn()),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		})
	if err != nil {
		panic(err)
	}
}
