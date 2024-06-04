package pack

import (
	"hertz_ucenter/biz/dal/db"
	"hertz_ucenter/biz/model/hertz/user"
)

// Users 转换数据库模型为业务模型
func Users(models []*db.User) []*user.User {
	users := make([]*user.User, 0, len(models))
	for _, m := range models {
		if u := User(m); u != nil {
			users = append(users, u)
		}
	}
	return users
}

func User(model *db.User) *user.User {
	if model == nil {
		return nil
	}
	return &user.User{
		Id:           int64(model.ID),
		UserName:     model.UserName,
		UserAccount:  model.UserAccount,
		AvatarUrl:    model.AvatarUrl,
		Gender:       user.Gender(model.Gender),
		UserPassword: model.UserPassword,
		Phone:        model.Phone,
		Email:        model.Email,
		UserStatus:   int64(model.UserStatus),
		UserRole:     int64(model.UserRole),
	}
}
