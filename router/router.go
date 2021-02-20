package router

import (
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/swagger"
	"lqc.com/tree/app/api"
	"lqc.com/tree/app/service"
)

func init() {

	s := g.Server()
	gtoken := &gtoken.GfToken{
		LoginPath:        "/login",
		LoginBeforeFunc:  api.User.Login,
		LogoutPath:       "/logout",
		AuthExcludePaths: g.SliceStr{"/user/signup"}, // 不拦截路径 /user/info,/system/user/info,/system/user,
		GlobalMiddleware: true,                       // 开启全局拦截，默认关闭
	}

	s.Use(service.Middlewave.MiddlewareErrorHandler)
	s.Plugin(&swagger.Swagger{})
	s.Use(service.Middlewave.Ctx)
	s.Use(service.Middlewave.MiddlewareTea)

	root := s.Group("/")
	gtoken.Middleware(root)

	root.ALL("/user", api.User)
	root.ALL("/postvote", api.PostVote)
	root.ALL("/replyvote", api.ReplyVote)
	root.ALL("/post", api.Post)
	root.ALL("/reply", api.Reply)

}
