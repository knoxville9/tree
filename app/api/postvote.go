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

var PostVote = new(postVoteApi)

type postVoteApi struct {
}

//赞踩
// @Summary 踩赞
// @Param pid formData  string true "post的id"
// @Param vote formData string true "1赞,0踩"
// @Param token formData string true "token"
// @Router /postvote/vote [post]
func (a *postVoteApi) Vote(r *ghttp.Request) {
	g.Log().SetConfigWithMap(g.Map{
		"path":     "log",
		"level":    "all",
		"stdout":   false,
		"StStatus": 0,
		"file":     "postVote-{y-m-d}.log",
	})
	var req *model.PostvoteDoVote

	if err := r.Parse(&req); err != nil {
		if a, ok := err.(*gvalid.Error); ok {
			response.JsonExit(r, http.StatusInternalServerError, a.FirstString(), nil)
		}
	} else {
		if err := service.PostVote.Vote(r.Context(), req); err != nil {
			response.JsonExit(r, http.StatusInternalServerError, err.Error(), nil)
		}

	}
	response.Json(r, http.StatusOK, "ok!~")
	g.Log().Println(r.Request, req)

}
