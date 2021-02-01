package router

import (
	"github.com/gogf/gf/frame/g"
	"lqc.com/tree/app/api/user"
	"lqc.com/tree/app/service"
)

func init() {
	server := g.Server()
	server.Use(service.Middlewave.Ctx)
	server.BindHandler("/a", user.User.A)
	server.BindHandler("/b", user.User.B)

}
