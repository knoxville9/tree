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

var Post = new(postApi)

type postApi struct {
}

//列表
func (a *postApi) List(r *ghttp.Request) {
	var req *model.PostDoList

	if err := r.Parse(&req); err != nil {
		if a, ok := err.(*gvalid.Error); ok {
			response.JsonExit(r, http.StatusInternalServerError, a.FirstString(), nil)
		}
	} else {
		if list, err := service.Post.List(r.Context(), req); err != nil {
			response.JsonExit(r, http.StatusInternalServerError, err.Error(), nil)
		} else {
			pagination := model.Pagination{
				Page:  req.Page,
				Limit: req.Size,
				Size:  len(list),
			}
			err := r.Response.WriteJson(
				g.Map{
					"code":       http.StatusOK,
					"pagination": pagination,
					"data":       list,
				})
			if err != nil {
				response.JsonExit(r, http.StatusInternalServerError, "", nil)
			}

		}

	}
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
