package utils

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/sessions"
	"hertz_ucenter/pkg/consts"
)

func IsLogout(_ context.Context, c *app.RequestContext, userAccount string) bool {
	session := sessions.Default(c)
	sessionUserAccount := session.Get(consts.UserAccount)
	hlog.Infof("sessionUserAccount: %v", sessionUserAccount)
	hlog.Infof("userAccount: %v", userAccount)
	// 判断sessionUserAccount存在且与当前请求的UserAccount一致
	if sessionUserAccount != nil && sessionUserAccount == userAccount {
		return false
	}
	return true
}
