package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"hertz_ucenter/biz/dal/db"
	"hertz_ucenter/biz/model/hertz/user"
	"hertz_ucenter/biz/pack"
	"hertz_ucenter/pkg/errno"
	"hertz_ucenter/pkg/utils"
)

type UserService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewUserService(ctx context.Context, c *app.RequestContext) *UserService {
	return &UserService{
		ctx: ctx,
		c:   c,
	}
}

func (s *UserService) UserRegister(req *user.UserRegisterRequest) (uint, error) {
	// 校验是否存在该用户
	queryUser, err := db.QueryUser(req.UserAccount)
	if err != nil {
		return 0, err
	}
	if *queryUser != (db.User{}) {
		return 0, errno.UserAlreadyExistErr
	}
	// 校验userAccount
	// 非空
	if req.UserAccount == "" {
		return 0, errno.ParamErr.WithMessage("param error, userAccount can't be empty")
	}
	// 长度不小于4位
	if len(req.UserAccount) < 4 {
		return 0, errno.ParamErr.WithMessage("param error, userAccount length must be longer than 4")
	}
	hashedPasswd, err := utils.Crypt(req.UserPassword)
	userId, err := db.CreateUser(&db.User{
		UserAccount:  req.UserAccount,
		UserPassword: hashedPasswd,
		UserName:     req.UserAccount,
		Gender:       0, // 默认性别为0, 未知
		UserStatus:   1, // 默认状态为1, 正常
		UserRole:     1, // 默认角色为1, 普通用户, 0为管理员
	})
	return userId, err
}

func (s *UserService) QueryUserByPage(req *user.UserRequest) ([]*user.User, error) {
	dbUsers, err := db.QueryUserByPage(req.CurrentPage, req.PageSize)
	if err != nil {
		return nil, err
	}
	users := pack.Users(dbUsers)

	return users, err
}
