package main

import (
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	_ "lqc.com/tree/router"
)

func main() {
	s := g.Server()

	//gfToken := &gtoken.GfToken{
	//	LoginPath:       "/login",
	//	LoginBeforeFunc: loginFunc,
	//	LogoutPath:      "/logout",
	//}
	//s.Group("/", func(group *ghttp.RouterGroup) {
	//	gfToken.Middleware(group)
	//
	//	group.ALL("/index", func(r *ghttp.Request) {
	//		response.Json(r,http.StatusOK,"a","a")
	//	})
	//
	//})

	s.Run()
}

func loginFunc(r *ghttp.Request) (string, interface{}) {
	username := r.GetVar("passport").String()
	passwd := r.Get("passwd")

	// TODO 进行登录校验
	if username == "" || passwd == "" {
		r.Response.WriteJson(gtoken.Fail("账号或密码错误."))
		r.ExitAll()
	}

	return username, ""

}
