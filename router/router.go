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
	User := s.Group("/user")
	User.ALL("/login", api.User.Login)
	User.ALL("/signup", api.User.Signup)
	User.ALL("/logout", api.User.Logout)
	User.Middleware(service.Middlewave.Auth)
	User.ALL("/profile", api.User.Profile)
	Post := s.Group("/post")
	Post.Middleware(service.Middlewave.Auth)
	Post.ALL("/create", api.Post.Create)

}
