package user

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
	"lqc.com/tree/app/model"
	"lqc.com/tree/app/service"
	"lqc.com/tree/library/response"
	"net/http"
)

var User = new(userApi)

type userApi struct {
}

//注册
func (u *userApi) Signup(r *ghttp.Request) {
	ctxVar := r.GetCtx()
	var req *model.UserDoSignUpReq

	if err := r.Parse(&req); err != nil {
		if a, ok := err.(*gvalid.Error); ok {
			response.JsonExit(r, http.StatusInternalServerError, a.FirstString(), "")

		}

	}

	if err := service.User.Signup(ctxVar, req); err != nil {
		response.JsonExit(r, http.StatusInternalServerError, err.Error(), "")
	}

	if err := r.Response.WriteJsonExit(
		g.Map{
			"code":  http.StatusOK,
			"error": "",
			"data":  "ok",
		}); err != nil {
		panic(err)

	}

}

//登录
func (u *userApi) Login(r *ghttp.Request) {

	var req *model.UserDoLogInReq
	ctxVar := r.GetCtx()
	if err := r.Parse(&req); err != nil {
		if a, ok := err.(*gvalid.Error); ok {
			response.JsonExit(r, http.StatusInternalServerError, a.FirstString(), "")
		}
	}
	if err := service.User.Login(ctxVar, req); err != nil {
		response.JsonExit(r, http.StatusInternalServerError, err.Error(), "")

	} else {
		response.JsonExit(r, http.StatusOK, "ok", "")
	}

}

//获取用户信息
func (u *userApi) Profile(r *ghttp.Request) {

	service.Session.GetUser(r.GetCtx())
	response.JsonExit(r, http.StatusOK, "ok", service.Session.GetUser(r.GetCtx()))

}

//退出
func (u *userApi) Logout(r *ghttp.Request) {
	ctx := r.GetCtx()
	if err := service.User.Logout(ctx); err != nil {
		panic(err)
	}

}
