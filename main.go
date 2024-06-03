// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"hertz_ucenter/biz/dal"
	"hertz_ucenter/biz/dal/db"
)

func Init() {
	dal.Init()
}

func main() {
	Init()

	// 判断是否需要自动迁移
	if !db.DB.Migrator().HasTable(&db.User{}) {
		hlog.Infof("数据库表不存在，开始自动迁移")
		err := db.DB.AutoMigrate(&db.User{})
		if err != nil {
			panic(err)
		}
	}

	h := server.Default()
	register(h)
	h.Spin()
}
