package service

import (
	"github.com/gogf/gf/net/ghttp"
	"lqc.com/tree/app/model"
)

var Middlewave = new(middlewaveService)

type middlewaveService struct {
}

func (s *middlewaveService) Ctx(r *ghttp.Request) {
	context := &model.Context{}

	context.User = &model.User{}
	Context.Init(r, context)
	if user := Session.GetUser(r.GetCtx()); user != nil {
		context.User = user
	}

	r.Middleware.Next()

}
