package mw

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/logger/accesslog"
)

func InitAccessLog(h *server.Hertz) {
	// accessLog middleware
	h.Use(accesslog.New())
}
