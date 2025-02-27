// Code generated by hertz generator. DO NOT EDIT.

package share

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	share "github.com/dgdts/CloudNoteServer/biz/handler/share"
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
				_share := _v1.Group("/share", _shareMw()...)
				_share.GET("/note", append(_getsharenoteMw(), share.GetShareNote)...)
				_note := _share.Group("/note", _noteMw()...)
				_note.POST("/comment", append(_createsharenotecommentMw(), share.CreateShareNoteComment)...)
				_note.GET("/comments", append(_listsharenotecommentsMw(), share.ListShareNoteComments)...)
			}
		}
	}
}
