package service

import (
	"context"
	"lqc.com/tree/app/model"
)

var Post = new(postService)

type postService struct {
}

func (s *postService) Create(ctx context.Context, req *model.PostDoCreate) {

}
