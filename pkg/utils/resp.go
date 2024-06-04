package utils

import (
	"errors"
	"hertz_ucenter/pkg/errno"
)

type BaseResp struct {
	StatusCode int32
	StatusMsg  string
}

func BuildBaseResp(err error) *BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *BaseResp {
	return &BaseResp{
		StatusMsg:  err.ErrMsg,
		StatusCode: err.ErrCode,
	}
}
