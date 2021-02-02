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

	r.SetCtxVar(ContextKey, context)
	r.Middleware.Next()

}
