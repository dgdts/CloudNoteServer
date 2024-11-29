// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	binary "github.com/dgdts/UniversalServer/biz/router/binary"
	note "github.com/dgdts/UniversalServer/biz/router/note"
	user "github.com/dgdts/UniversalServer/biz/router/user"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	binary.Register(r)

	user.Register(r)

	note.Register(r)
}
