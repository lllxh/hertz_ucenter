package pack

import (
	"hertz_ucenter/biz/dal/db"
	"hertz_ucenter/biz/model/hertz/user"
)

// SafeUsers 转换数据库模型DTO为领域模型DO
func SafeUsers(models []*db.User) []*user.User {
	users := make([]*user.User, 0, len(models))
	for _, m := range models {
		if u := SafeUser(m); u != nil {
			users = append(users, u)
		}
	}
	return users
}

func SafeUser(model *db.User) *user.User {
	if model == nil {
		return nil
	}
	return &user.User{
		Id:          int64(model.ID),
		UserName:    model.UserName,
		UserAccount: model.UserAccount,
		AvatarUrl:   model.AvatarUrl,
		Gender:      user.Gender(model.Gender),
		Phone:       model.Phone,
		Email:       model.Email,
		UserStatus:  int64(model.UserStatus),
		UserRole:    int64(model.UserRole),
	}
}
