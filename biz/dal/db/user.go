package db

import (
	"gorm.io/gorm"
	"hertz_ucenter/pkg/errno"
	"hertz_ucenter/pkg/utils"
	"strconv"
)

type User struct {
	UserName     string `gorm:"column:username" json:"user_name"`
	UserAccount  string `gorm:"column:userAccount" json:"user_account"`
	AvatarUrl    string `gorm:"column:avatarUrl" json:"avatar_url"`
	Gender       int8   `gorm:"column:gender" json:"gender"`
	UserPassword string `gorm:"column:userPassword" json:"user_password"`
	Phone        string `gorm:"column:phone" json:"phone"`
	Email        string `gorm:"column:email" json:"email"`
	UserStatus   int    `gorm:"column:userStatus" json:"user_status"`
	UserRole     int    `gorm:"column:userRole" json:"user_role"`
	gorm.Model
}

func (User) TableName() string {
	return "user"
}

func CreateUser(user *User) (userId uint, err error) {
	if err := DB.Create(user).Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}

func QueryUserLogin(userAccount, userPassword string) (*User, error) {
	var user User
	if err := DB.Where("userAccount = ? and deleted_at IS NULL", userAccount).Find(&user).Error; err != nil {
		return nil, errno.UserIsNotExistErr
	}
	if !utils.ComparePassword(userPassword, user.UserPassword) {
		return nil, errno.AuthorizationFailedErr
	}
	return &user, nil
}

func QueryUser(userAccount string) (*User, error) {
	var user User
	if err := DB.Where("userAccount = ?", userAccount).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// QueryUserByPage 按照分页查询用户
func QueryUserByPage(currentPage, pageSize string) ([]*User, error) {
	var users []*User
	// 转int
	currentPageInt, err := strconv.Atoi(currentPage)
	if err != nil {
		return nil, err
	}
	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		return nil, err
	}
	// 分页逻辑
	p := &utils.Pagination{
		PageSize:    pageSizeInt,
		CurrentPage: currentPageInt,
	}
	if err := DB.Limit(p.PageSize).Offset(p.Offset()).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
