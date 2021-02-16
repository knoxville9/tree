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
	user := ctx.Value(ContextKey).(*model.Context).User

	if one, err := dao.Post.Where("userid = ? and deleted = ?", user.Id, 0).Order("id desc").FindOne("title = ?", req.Title); err != nil {
		return err
	} else {
		findOne, _ := dao.Post.Where("userid = ? and deleted = ?", user.Id, 0).Order("CreateAt desc").FindOne()
		if time.Now().Unix()-findOne.CreateAt.Timestamp() < 300 {
			fmt.Println(time.Now().Unix() - findOne.CreateAt.Timestamp())
			return errors.New("5分钟内已经发过贴,请等待5分钟~")
		}

		if one != nil {
			return errors.New("标题重复,请重新命名~")
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
	//返回没有被删除的post
	if all, err := dao.Post.Page(req.Page, req.Size).Where("deleted = 0").FindAll(); err != nil {
		return nil, err
	} else {
		return all, nil

	}

}

func (s *postService) Vote(ctx context.Context, req *model.PostvoteDoVote) error {
	user := ctx.Value(ContextKey).(*model.Context).User
	//1.同一个post,同一个userid不能点赞多次
	//2.
	if count, err := dao.Post.Where("id = ?", req.Pid).FindCount(); err != nil {
		return err
	} else {
		if count == 0 {
			return errors.New("post不存在~")
		}

	}
	if one, err := dao.Postvote.Where("pid=? and userid = ?", req.Pid, user.Id).FindOne(); err != nil {
		return err
	} else {
		//判断是否为空
		if one == nil {
			//将上下文信息 用户id放入请求中
			req.Userid = user.Id
			dao.Postvote.Save(req)
			//如果预期操作和已存在数据库中的操作不一致则更新
		} else if one.Vote != *req.Vote {
			dao.Postvote.Data("vote = ?", req.Vote).Where("pid=? and userid = ?", req.Pid, user.Id).Update()
			//如果操作一样则提示
		} else if one.Vote == *req.Vote {
			return errors.New("你已经操作过了")
		}

	}

	return nil

}

func (s *postService) Delete() {}
