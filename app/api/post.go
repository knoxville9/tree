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
// @Summary 获得post列表
// @Param page query string true "页数"
// @Param size query string true "行数"
// @Router /post/list [get]
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
// @Summary 发post
// @Param title formData string true "标题"
// @Param content formData string true "内容"
// @Router /post/create [post]
func (a *postApi) Create(r *ghttp.Request) {
	var req *model.PostDoCreate

	if err := r.Parse(&req); err != nil {
		if a, ok := err.(*gvalid.Error); ok {
			response.JsonExit(r, http.StatusInternalServerError, a.FirstString(), nil)
		}

	} else {
		b := []byte(req.Title)
		if len(b) != 0 {
			if b[0] == '\n' || b[0] == ' ' || b[0] == '\t' {
				response.JsonExit(r, http.StatusInternalServerError, "标题请勿以空白符开始~", nil)
			}
		}

		if err := service.Post.Create(r.Context(), req); err != nil {
			response.JsonExit(r, http.StatusInternalServerError, err.Error(), nil)
		}
	}
	response.Json(r, http.StatusOK, "发帖成功")

}

//删帖
// @Summary 删除post
// @Param pid formData string true "postid"
// @Router /post/delete [post]
func (a *postApi) Delete(r *ghttp.Request) {
	var req *model.PostDoDelete

	if err := r.Parse(&req); err != nil {
		if a, ok := err.(*gvalid.Error); ok {
			response.JsonExit(r, http.StatusInternalServerError, a.FirstString(), nil)
		}
	} else {
		if err := service.Post.Delete(r.Context(), req); err != nil {
			response.JsonExit(r, http.StatusInternalServerError, err.Error(), nil)
		}

	}
	response.Json(r, http.StatusOK, "删除post成功~")
}

//post详细
func (a *postApi) Detail(r *ghttp.Request) {
	var req *model.PostDoDetail
	if err := r.Parse(&req); err != nil {
		if a, ok := err.(*gvalid.Error); ok {
			response.JsonExit(r, http.StatusInternalServerError, a.FirstString(), nil)
		}
	} else {

	}
	response.Json(r, http.StatusOK, "删除post成功~")

}
