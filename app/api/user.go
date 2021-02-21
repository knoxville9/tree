package api

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

// @Summary 注册
// @Param passport formData string true "用户名"
// @Param password formData string true "密码"
// @Param nickname formData string false "昵称"
// @Router /user/signup [post]
func (u *userApi) Signup(r *ghttp.Request) {
	ctxVar := r.GetCtx()
	var req *model.UserDoSignUpReq

	if err := r.Parse(&req); err != nil {
		if a, ok := err.(*gvalid.Error); ok {
			response.JsonExit(r, http.StatusInternalServerError, a.FirstString(), nil)
		}
	}

	if err := service.User.Signup(ctxVar, req); err != nil {
		response.JsonExit(r, http.StatusInternalServerError, err.Error(), nil)
	}

	if err := r.Response.WriteJson(
		g.Map{
			"code":  http.StatusOK,
			"error": nil,
			"data":  "注册成功",
		}); err != nil {
		panic(err)

	}

}

// @Summary 登录
// @Param passport formData string true "用户名"
// @Param password formData string true "密码"
// @Router /login [post]
func (u *userApi) Login(r *ghttp.Request) (string, interface{}) {

	var req *model.UserDoLogInReq
	//校验登录参数
	if err := r.Parse(&req); err != nil {

		if a, ok := err.(*gvalid.Error); ok {
			response.JsonExit(r, http.StatusInternalServerError, a.FirstString(), nil)
		}
	}

	if err := service.User.Login(r.Context(), req); err != nil {
		response.JsonExit(r, http.StatusInternalServerError, err.Error(), nil)
	}

	if !r.GetVar("passport").IsNil() {
		return r.GetVar("passport").String(), ""

	}
	return "", nil

}

// @Summary 获得用户信息
// @Param token formData string true "token"
// @Router /user/profile [Post]
func (u *userApi) Profile(r *ghttp.Request) {
	if user, err := service.User.Profile(r.Context()); err != nil {
		response.JsonExit(r, http.StatusInternalServerError, err.Error(), nil)
	} else {
		response.Json(r, http.StatusOK, "nil", user)
	}

}

// @Summary 退出
// @Param token formData string true "token"
// @Router /logout [Post]
func (u *userApi) Logout(r *ghttp.Request) {

}
