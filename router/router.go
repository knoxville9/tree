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
	User := s.Group("/user")
	gtoken.Middleware(User)
	User.ALL("/signup", api.User.Signup)
	User.ALL("/profile", api.User.Profile)
	Post := s.Group("/post")

	gtoken.Middleware(Post)
	Post.ALL("/", api.Post)

	PostVote := s.Group("/postvote")
	gtoken.Middleware(PostVote)
	PostVote.ALL("/", api.PostVote)

	reply := s.Group("/reply")
	gtoken.Middleware(reply)
	reply.ALL("/", api.Reply)

}
