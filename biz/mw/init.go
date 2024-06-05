package mw

import "github.com/cloudwego/hertz/pkg/app/server"

func InitMiddleware(h *server.Hertz) {
	InitSession(h)
	InitAccessLog(h)
}
