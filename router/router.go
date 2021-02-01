package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	server := g.Server()

	server.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/", func(r *ghttp.Request) {
			r.Response.Writeln(r.Get("list"))

		})

	})
}
