// Code generated by hertz generator. DO NOT EDIT.

package user

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	user "hertz_ucenter/biz/handler/user"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_v1 := root.Group("/v1", _v1Mw()...)
		{
			_user := _v1.Group("/user", _userMw()...)
			_user.POST("/login", append(_userloginMw(), user.UserLogin)...)
			_user.POST("/logout", append(_userlogoutMw(), user.UserLogout)...)
			_user.POST("/query", append(_queryuserMw(), user.QueryUser)...)
			_user.POST("/register", append(_userregisterMw(), user.UserRegister)...)
			{
				_delete := _user.Group("/delete", _deleteMw()...)
				_delete.POST("/:user_id", append(_deleteuserMw(), user.DeleteUser)...)
			}
			{
				_update := _user.Group("/update", _updateMw()...)
				_update.POST("/:user_id", append(_updateuserMw(), user.UpdateUser)...)
			}
		}
	}
}
