package api

import (
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
// @Summary 发post
// @Param content formData string true "内容"
// @Router /reply/create [post]
func (a *replyApi) Create(r *ghttp.Request) {
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

}

//删帖
// @Summary 删除post
// @Param pid formData string true "pid"
// @Param rid formData string true "rid"
// @Router /reply/delete [post]
func (a *replyApi) Delete(r *ghttp.Request) {
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
}
