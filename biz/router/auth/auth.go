// Code generated by hertz generator. DO NOT EDIT.

package auth

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	auth "github.com/dgdts/CloudNoteServer/biz/handler/auth"
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
		_api := root.Group("/api", _apiMw()...)
		{
			_v1 := _api.Group("/v1", _v1Mw()...)
			{
				_auth := _v1.Group("/auth", _authMw()...)
				_auth.POST("/login", append(_loginMw(), auth.Login)...)
				_auth.POST("/logout", append(_logoutMw(), auth.Logout)...)
				_auth.POST("/refresh", append(_refreshtokenMw(), auth.RefreshToken)...)
				_auth.POST("/register", append(_registerMw(), auth.Register)...)
			}
			{
				_users := _v1.Group("/users", _usersMw()...)
				_users.GET("/me", append(_getcurrentuserMw(), auth.GetCurrentUser)...)
			}
		}
	}
}
