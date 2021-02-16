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
	if one, err := dao.Post.Where("userid = ? and deleted = ?", req.UserId, 0).Order("id desc").FindOne("title = ?", req.Title); err != nil {
		return err
	} else {
		findOne, _ := dao.Post.Where("userid = ? and deleted = ?", req.UserId, 0).Order("CreateAt desc").FindOne()
		if time.Now().Unix()-findOne.CreateAt.Timestamp() < 300 {
			fmt.Println(time.Now().Unix() - findOne.CreateAt.Timestamp())
			return errors.New("5分钟内已经发过贴,请等待5分钟")
		}

		if one != nil {
			return errors.New("标题重复,请重新命名")
		}

	}

	if _, err := dao.Post.Save(req); err != nil {
		return err
	}
	return nil
}

type Pagination struct {
	Code int
}

func (s *postService) List(ctx context.Context, req *model.PostDoList) ([]*model.Post, error) {
	if all, err := dao.Post.Page(req.Page, req.Size).Where("deleted = 0").FindAll(); err != nil {
		return nil, err
	} else {
		return all, nil

	}

}
