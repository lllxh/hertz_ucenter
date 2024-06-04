// Code generated by hertz generator.

package user

import (
	"context"
	"hertz_ucenter/biz/service"
	"hertz_ucenter/pkg/errno"
	"hertz_ucenter/pkg/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	user "hertz_ucenter/biz/model/hertz/user"
)

// UserRegister .
// @router /v1/user/register [POST]
func UserRegister(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserRegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, user.UserRegisterResponse{
			Code: resp.StatusCode,
			Msg:  resp.StatusMsg,
		})
		return
	}
	userId, err := service.NewUserService(ctx, c).UserRegister(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, user.UserRegisterResponse{
			Code: resp.StatusCode,
			Msg:  resp.StatusMsg,
		})
		return
	}
	c.JSON(consts.StatusOK, user.UserRegisterResponse{
		Code:   errno.SuccessCode,
		Msg:    errno.SuccessMsg,
		UserId: uint64(userId),
	})

}

// UserLogin .
// @router /v1/user/login [POST]
func UserLogin(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserLoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.UserLoginResponse)

	c.JSON(consts.StatusOK, resp)
}

// QueryUserResponse .
// @router /v1/user/query [GET]
func QueryUserResponse(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, user.UserResponse{
			Code: resp.StatusCode,
			Msg:  resp.StatusMsg,
		})
		return
	}

	//u, err := service.NewUserService(ctx, c).UserInfo(&req)
	//
	//resp := utils.BuildBaseResp(err)
	//c.JSON(consts.StatusOK, user.UserResponse{
	//	Code: resp.StatusCode,
	//	Msg:  resp.StatusMsg,
	//	User: u,
	//})
}

// UpdateUserResponse .
// @router /v1/user/update/:user_id [POST]
func UpdateUserResponse(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserUpdateRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.UserUpdateResponse)

	c.JSON(consts.StatusOK, resp)
}

// DeleteUserResponse .
// @router /v1/user/delete/:user_id [POST]
func DeleteUserResponse(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserDeleteRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.UserDeleteResponse)

	c.JSON(consts.StatusOK, resp)
}

// UserLogout .
// @router /v1/user/logout [POST]
func UserLogout(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserLogoutRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.UserLogoutResponse)

	c.JSON(consts.StatusOK, resp)
}