package service

import (
	"github.com/gogf/gf/net/ghttp"
	"lqc.com/tree/app/model"
	"lqc.com/tree/library/response"
	"net/http"
)

var Middlewave = new(middlewaveService)

type middlewaveService struct {
}

func (s *middlewaveService) Ctx(r *ghttp.Request) {
	context := &model.Context{
		Session: r.Session,
	}
	context.User = &model.User{}
	Context.Init(r, context)

	if a := Session.GetUser(r.Context()); a != nil {
		context.User = a
	}

	r.Middleware.Next()

}

func (s *middlewaveService) Auth(r *ghttp.Request) {
	if Session.GetUser(r.Context()) != nil {
		r.Middleware.Next()
	} else {
		response.JsonExit(r, http.StatusForbidden, "认证失败,请重新登录")
	}

}
