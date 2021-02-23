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

var ReplyVote = new(replyVoteApi)

type replyVoteApi struct {
}

//赞踩
// @Summary 踩赞
// @Param replyid formData  string true "回复id"
// @Param vote formData string true "1赞,0踩"
// @Param token formData string true "token"
// @Router /replyvote/vote [post]
func (a *replyVoteApi) Vote(r *ghttp.Request) {
	g.Log().SetConfigWithMap(g.Map{
		"path":     "log",
		"level":    "all",
		"stdout":   false,
		"StStatus": 0,
		"file":     "replyVote-{y-m-d}.log",
	})
	var req *model.ReplyvoteDoVote

	if err := r.Parse(&req); err != nil {
		if a, ok := err.(*gvalid.Error); ok {
			response.JsonExit(r, http.StatusInternalServerError, a.FirstString(), nil)
		}
	} else {
		if err := service.ReplyVote.Vote(r.Context(), req); err != nil {
			response.JsonExit(r, http.StatusInternalServerError, err.Error(), nil)
		}

	}
	response.Json(r, http.StatusOK, "ok!~")
	g.Log().Println(r.Request, req)

}
