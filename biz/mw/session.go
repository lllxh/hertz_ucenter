package mw

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/cookie"
	"hertz_ucenter/pkg/consts"
)

func InitSession(h *server.Hertz) {
	// session middleware
	store := cookie.NewStore([]byte("secret"))
	h.Use(sessions.New(consts.UcenterSession, store))
}
