package service

import (
	"context"
	"lqc.com/tree/app/dao"
	"lqc.com/tree/app/model"
)

var PostVote = new(postVoteService)

type postVoteService struct {
}

func (s *postVoteService) Vote(ctx context.Context, req *model.PostvoteDoVote) error {
	user := ctx.Value(ContextKey).(*model.Context).User
	//1.同一个post,同一个userid不能点赞多次
	//2.
	if err := IfPostExist(*req.Pid); err != nil {
		return err
	}

	if one, err := dao.Postvote.Where("pid=? and userid = ?", req.Pid, user.Id).FindOne(); err != nil {
		return err
	} else {
		//判断是否为空
		if one == nil {
			req.Userid = int(user.Id)
			//将上下文信息 用户id放入请求中
			dao.Postvote.Save(req)
			//如果预期操作和已存在数据库中的操作不一致则更新
		} else if one.Vote != *req.Vote {
			dao.Postvote.Data("vote = ?", req.Vote).Where("pid=? and userid = ?", req.Pid, user.Id).Update()
			//如果操作一样则提示
		} else if one.Vote == *req.Vote {
			dao.Postvote.Data("vote = ?", req.Vote).Where("pid=? and userid = ?", req.Pid, user.Id).Delete()
		}

	}

	return nil

}
