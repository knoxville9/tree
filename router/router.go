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
		CacheMode:        2,
		LoginPath:        "/login",
		LoginBeforeFunc:  api.User.Login,
		LogoutPath:       "/logout",
		AuthExcludePaths: g.SliceStr{"/signup"}, // 不拦截路径 /user/info,/system/user/info,/system/user,
		GlobalMiddleware: true,                  // 开启全局拦截，默认关闭
	}
	//

	s.Use(service.Middlewave.MiddlewareErrorHandler)
	s.Plugin(&swagger.Swagger{})
	s.Use(service.Middlewave.Ctx)

	root := s.Group("/")
	gtoken.Middleware(root)

	root.ALLMap(g.Map{
		"/list":             api.Post.List,
		"/signup":           api.User.Signup,
		"/post/create":      api.Post.Create,
		"/post/delete/:pid": api.Post.Delete,
		"/post/detail/:pid": api.Post.Detail,
		"/profile":          api.User.Profile,
		"/reply/delete":     api.Reply.Delete,
		"/reply/create":     api.Reply.Create,
		"/postvote":         api.PostVote.Vote,
		"/replyvote":        api.ReplyVote.Vote,
	})
}
