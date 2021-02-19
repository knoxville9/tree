package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
	"lqc.com/tree/app/model"
	"lqc.com/tree/library/response"
	"net/http"
)

var Reply = new(replyApi)

type replyApi struct {
}

func (a *replyApi) Create(r *ghttp.Request) {
	var req *model.ReplyDoCreate

	if err := r.Parse(&req); err != nil {
		if a, ok := err.(*gvalid.Error); ok {
			response.JsonExit(r, http.StatusInternalServerError, a.FirstString(), nil)
		}
	} else {

	}
	response.Json(r, http.StatusOK, "ok!~")

}
func (a *replyApi) Delete() {}
func (a *replyApi) List()   {}
