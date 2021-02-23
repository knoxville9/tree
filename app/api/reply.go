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

var Reply = new(replyApi)

type replyApi struct {
}

//发帖
// @Summary 回复post
// @Param content formData string true "内容"
// @Param token formData string true "token"
// @Router /reply/create [post]
func (a *replyApi) Create(r *ghttp.Request) {
	g.Log().SetConfigWithMap(g.Map{
		"path":     "log",
		"level":    "all",
		"stdout":   false,
		"StStatus": 0,
		"file":     "replyCreate-{y-m-d}.log",
	})
	var req *model.ReplyDoCreate

	if err := r.Parse(&req); err != nil {
		if a, ok := err.(*gvalid.Error); ok {
			response.JsonExit(r, http.StatusInternalServerError, a.FirstString(), nil)
		}

	} else {
		if err := service.Reply.Create(r.Context(), req); err != nil {
			response.JsonExit(r, http.StatusInternalServerError, err.Error(), nil)
		}

	}
	response.Json(r, http.StatusOK, "回复成功")
	g.Log().Println(r.Request, req)

}

//删帖
// @Summary 删除回复,只能删除自己的回复
// @Param pid formData string true "pid"
// @Param rid formData string true "rid"
// @Param token formData string true "token"
// @Router /reply/delete [post]
func (a *replyApi) Delete(r *ghttp.Request) {
	g.Log().SetConfigWithMap(g.Map{
		"path":     "log",
		"level":    "all",
		"stdout":   false,
		"StStatus": 0,
		"file":     "replyDelete-{y-m-d}.log",
	})
	var req *model.ReplyDoDelete

	if err := r.Parse(&req); err != nil {
		if a, ok := err.(*gvalid.Error); ok {
			response.JsonExit(r, http.StatusInternalServerError, a.FirstString(), nil)
		}
	} else {
		if err := service.Reply.Delete(r.Context(), req); err != nil {
			response.JsonExit(r, http.StatusInternalServerError, err.Error(), nil)
		}

	}
	response.Json(r, http.StatusOK, "删除回复成功~")
	g.Log().Println(r.Request, req)
}
