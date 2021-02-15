package service

import (
	"context"
	"errors"
	"fmt"
	"lqc.com/tree/app/dao"
	"lqc.com/tree/app/model"
	"time"
)

var Post = new(postService)

type postService struct {
}

func (s *postService) Create(ctx context.Context, req *model.PostDoCreate) error {
	if one, err := dao.Post.Where("userid = ?", req.UserId).Order("id desc").FindOne("title = ?", req.Title); err != nil {
		return err
	} else {
		if one != nil {
			return errors.New("标题重复,请重新命名")
		}

		findOne, _ := dao.Post.Where("userid = ?", req.UserId).Order("CreateAt desc").FindOne()
		if time.Now().Unix()-findOne.CreateAt.Timestamp() < 300 {
			fmt.Println(time.Now().Unix() - findOne.CreateAt.Timestamp())
			return errors.New("5分钟内已经发过贴,请等待5分钟")
		}

	}

	if _, err := dao.Post.Save(req); err != nil {
		return err
	}
	return nil
}
