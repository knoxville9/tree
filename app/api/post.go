package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
	"lqc.com/tree/app/model"
	"lqc.com/tree/app/service"
	"lqc.com/tree/library/response"
	"net/http"
)

var Post = new(postApi)

type postApi struct {
}

//发帖
func (a *postApi) Create(r *ghttp.Request) {
	var req *model.PostDoCreate

	if err := r.Parse(&req); err != nil {
		if a, ok := err.(*gvalid.Error); ok {
			response.JsonExit(r, http.StatusInternalServerError, a.FirstString(), nil)
		}
	} else {
		req.UserId = int(r.Context().Value(service.ContextKey).(*model.Context).User.Id)

		if err := service.Post.Create(r.Context(), req); err != nil {
			response.JsonExit(r, http.StatusInternalServerError, err.Error(), nil)
		}
	}
	response.Json(r, http.StatusOK, "发帖成功")

}

//删帖
func (a *postApi) Delete() {
}

//赞
func (a *postApi) UpVote() {
}

//踩
func (a *postApi) DownVote() {
}

//列表
func (a *postApi) List() {
}
