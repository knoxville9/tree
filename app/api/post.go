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
// @Param page formData string true "页数"
// @Param size formData string true "行数"
// @Param token formData string true "token"
// @Router /post/list [Post]
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
// @Param token formData string true "token"
// @Router /post/create [post]
func (a *postApi) Create(r *ghttp.Request) {
	g.Log().SetConfigWithMap(g.Map{
		"path":     "log",
		"level":    "all",
		"stdout":   false,
		"StStatus": 0,
		"file":     "postCreate-{y-m-d}.log",
	})

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
	g.Log().Println(r.Request, req)

}

//
//删帖
// @Summary 删除post
// @Param pid formData string true "postid"
// @Param token formData string true "token"
// @Router /post/delete [post]
func (a *postApi) Delete(r *ghttp.Request) {
	g.Log().SetConfigWithMap(g.Map{
		"path":     "log",
		"level":    "all",
		"stdout":   false,
		"StStatus": 0,
		"file":     "postDelete-{y-m-d}.log",
	})

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
	g.Log().Println(r.Request, req)
}

//post详细
// @Summary post详情
// @Param pid formData string true "postid"
// @Param token formData string true "token"
// @Router /post/detail [post]
func (a *postApi) Detail(r *ghttp.Request) {
	var req *model.PostDoDetail
	if err := r.Parse(&req); err != nil {
		if a, ok := err.(*gvalid.Error); ok {
			response.JsonExit(r, http.StatusInternalServerError, a.FirstString(), nil)
		}
	} else {

		detail, err := service.Post.Detail(r.Context(), req)

		if err != nil {
			response.JsonExit(r, http.StatusInternalServerError, err.Error(), nil)
		}

		response.Json(r, http.StatusOK, "", detail)

	}

}
