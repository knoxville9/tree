package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"lqc.com/tree/app/service"
)

// @title       tree project
// @version     1.0
// @schemes     http
func main() {

	s := g.Server()
	s.Group("/admin", func(group *ghttp.RouterGroup) {
		group.ALL("/login", service.Middlewave.GenToken)
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.Middleware(service.Middlewave.AuthInRedis)
			group.ALL("/index", func(r *ghttp.Request) {
				r.Response.Writeln("ok")
			})

		})

	})

	g.Server().Run()
}
