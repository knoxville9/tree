package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/swagger"
	"lqc.com/tree/app/api/user"
	"lqc.com/tree/app/service"
)

func init() {

	server := g.Server()
	server.Plugin(&swagger.Swagger{})
	server.Use(service.Middlewave.Ctx)
	group := server.Group("/user")
	group.ALL("/login", user.User.Login)
	group.ALL("/signup", user.User.Signup)
	group.ALL("/logout", user.User.Logout)
	group.Middleware(service.Middlewave.Auth)
	group.ALL("/profile", user.User.Profile)
}
