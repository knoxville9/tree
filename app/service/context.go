package service

import (
	"context"
	"github.com/gogf/gf/net/ghttp"
	"lqc.com/tree/app/model"
)

var Context = new(ContextService)

type ContextService struct{}

func (c *ContextService) Init(r *ghttp.Request, ctx *model.Context) {
	r.SetCtxVar(model.ContextKey, ctx)

}

func (c *ContextService) Get(ctx context.Context) *model.Context {
	value := ctx.Value(model.ContextKey)
	if ctx.Value(model.ContextKey) == nil {
		return nil

	}

	if a, b := value.(*model.Context); b {
		return a
	}

	return nil
}

func (c *ContextService) Set(ctx context.Context, user *model.User) {

	c.Get(ctx).User = user

}
