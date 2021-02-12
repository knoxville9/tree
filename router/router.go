package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/swagger"
	"lqc.com/tree/app/api"
	"lqc.com/tree/app/service"
)

func init() {

	s := g.Server()
	s.Use(service.Middlewave.MiddlewareErrorHandler)
	s.Plugin(&swagger.Swagger{})
	s.Use(service.Middlewave.Ctx)
	group := s.Group("/user")
	group.ALL("/login", api.User.Login)
	group.ALL("/signup", api.User.Signup)
	group.ALL("/logout", api.User.Logout)
	group.Middleware(service.Middlewave.Auth)
	group.ALL("/profile", api.User.Profile)

}
