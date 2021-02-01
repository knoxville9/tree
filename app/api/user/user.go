package user

import "github.com/gogf/gf/net/ghttp"

var User = new(userApi)

type userApi struct {
}

func (u *userApi) HELLO(request *ghttp.Request) {
	request.Response.Writeln(request.Get("id"))

}
