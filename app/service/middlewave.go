package service

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/guid"
	"lqc.com/tree/app/model"
	"lqc.com/tree/library/response"
	"net/http"
)

var Middlewave = new(middlewaveService)

type middlewaveService struct {
}

func (s *middlewaveService) Ctx(r *ghttp.Request) {
	context := &model.Context{
		Session: r.Session,
	}
	context.User = &model.User{}
	//和下面的语句相同作用
	//r.SetCtxVar(ContextKey,context)
	Context.Init(r, context)

	if a := Session.GetUser(r.Context()); a != nil {
		context.User = a
	}

	r.Middleware.Next()

}

func (s *middlewaveService) Auth(r *ghttp.Request) {
	if Session.GetUser(r.Context()) != nil {
		r.Middleware.Next()
	} else {
		response.JsonExit(r, http.StatusForbidden, "认证失败,请重新登录")
	}

}

func (s *middlewaveService) AuthInRedis(r *ghttp.Request) {
	conn := g.Redis().Conn()
	defer conn.Close()
	doVar, _ := conn.DoVar("get", r.Session.Id())
	if r.Header.Get("token") == gconv.String(doVar) {
		r.Middleware.Next()
	} else {
		r.Response.WriteStatusExit(http.StatusForbidden)

	}
}

func (s *middlewaveService) MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

func (s *middlewaveService) GenToken(r *ghttp.Request) {
	r.Session.Set("time", guid.S())

	var token string
	conn := g.Redis().Conn()
	defer conn.Close()

	if g1, _ := conn.DoVar("EXISTS", r.Session.Id()); gconv.Int(g1) == 0 {
		//if g2, _ := conn.DoVar("ttl", r.Session.Id()); gconv.Int(g2) < 0 {
		token = guid.S()
		conn.DoVar("set", r.Session.Id(), token)
		conn.DoVar("EXPIRE", r.Session.Id(), 600)

		//}
	}
	doVar, _ := conn.DoVar("get", r.Session.Id())
	r.Response.Writeln(gconv.String(doVar))

}

func (s *middlewaveService) MiddlewareErrorHandler(r *ghttp.Request) {
	r.Middleware.Next()
	if err := r.GetError(); err != nil {
		// 记录到自定义错误日志文件
		g.Log("exception").Error(err)
		//返回固定的友好信息
		r.Response.ClearBuffer()
		r.Response.Writeln("服务器异常，请稍后再试")
	}
}
