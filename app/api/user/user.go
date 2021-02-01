package user

import (
	"github.com/gogf/gf/net/ghttp"
	"lqc.com/tree/app/model"
)

var User = new(userApi)

type userApi struct {
}

func (u *userApi) A(r *ghttp.Request) {
	ctxVar := r.GetCtxVar("a").Interface().(model.Context)
	ctxVar.User.Id = 100

}

func (u *userApi) B(r *ghttp.Request) {
	r.Response.Writeln(r.Context().Value("a"))

}
