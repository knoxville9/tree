package service

import (
	"context"
	"github.com/gogf/gf/net/ghttp"
	"lqc.com/tree/app/model"
)

var Context = new(ContextService)

type ContextService struct{}

var ContextKey = "ContextKey"

func (c *ContextService) Init(r *ghttp.Request, ctx *model.Context) {
	r.SetCtxVar(ContextKey, ctx)
}

func (c *ContextService) Get(ctx context.Context) *model.Context {
	value := ctx.Value(ContextKey)
	if ctx.Value(ContextKey) == nil {
		return nil

	}

	if a, b := value.(*model.Context); b {
		return a
	}

	return nil
}

func (c *ContextService) SetUser(ctx context.Context, user *model.User) {
	c.Get(ctx).User = user
}

func (c *ContextService) SetPost(ctx context.Context, post *model.Post) {
	c.Get(ctx).Post = post
}
