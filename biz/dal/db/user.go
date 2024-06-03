package db

import (
	"gorm.io/gorm"
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
