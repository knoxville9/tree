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

	if err := IfPostOften(req.Title, user.Id); err != nil {
		return err
	}
	if err := IfPostRename(int(user.Id), req.Title); err != nil {
		return err
	}
	req.Userid = int(user.Id)
	if _, err := dao.Post.Save(req); err != nil {
		return err
	}
	return nil
}

func (s *postService) List(ctx context.Context, req *model.PostDoList) ([]*model.Post, error) {
	//返回没有被删除的post
	if all, err := dao.Post.Page(req.Page, req.Size).Where("deleted = 0").FindAll(); err != nil {
		return nil, err
	} else {
		return all, nil

	}

}

func (s *postService) Delete(ctx context.Context, req *model.PostDoDelete) error {

	user := ctx.Value(ContextKey).(*model.Context).User

	if err := IfPostExist(*req.Pid); err != nil {
		return err
	}
	if err := IfPostOwned(uint(*req.Pid), user.Id); err != nil {
		return err
	}
	if err := IfPostDeleted(*req.Pid); err != nil {
		return err
	}

	_, err := dao.Post.Data("deleted", 1).Where("id = ? and userid = ?", req.Pid, user.Id).Update()
	if err != nil {
		return err
	}
	return nil
}

//查看post详情
func (s *postService) Detail(ctx context.Context, req *model.PostDoDetail) (*model.PostDetail, error) {

	if err := IfPostExist(*req.Pid); err != nil {
		return nil, err
	}

	if err := IfPostDeleted(*req.Pid); err != nil {
		return nil, err
	}

	post, err1 := dao.Post.Where("deleted", 0).FindOne(req.Pid)
	postvotes, err2 := dao.Postvote.Where("pid = ?", req.Pid).FindAll()
	replies, err3 := dao.Reply.Where("pid = ? and deleted = ?", req.Pid, 0).FindAll()
	repliesvote, err4 := dao.Replyvote.As("rv").InnerJoin("reply r", "rv.replyid=r.id").Where("r.pid = ?", req.Pid).FindAll()

	Detail := &model.PostDetail{}

	Detail.Post = post
	Detail.PostVote = postvotes
	Detail.Reply = replies
	Detail.Replyvote = repliesvote

	if err1 != nil {
		return nil, err1
	}
	if err2 != nil {
		return nil, err2
	}
	if err3 != nil {
		return nil, err3
	}
	if err4 != nil {
		return nil, err4
	}

	return Detail, nil
}

//post是否重名
func IfPostRename(userid int, title string) error {
	one, err := dao.Post.Where("userid = ? and deleted = ?", userid, 0).Order("id desc").FindOne("title = ?", title)
	if err != nil {
		return err
	}

	if one != nil {
		return errors.New("标题重复,请重新命名~")
	}

	return nil
}

//postid是否存在
func IfPostExist(req int) error {
	if count, err := dao.Post.Where("id = ?", req).FindOne(); err != nil {
		return err
	} else {
		if count == nil {
			return errors.New("post不存在~")
		}
	}
	return nil
}

//post是否已经被删除
func IfPostDeleted(pid int) error {
	one, err := dao.Post.Where("id = ?", pid).FindOne()
	if err != nil {
		return err
	}
	if one != nil {
		if one.Deleted == 1 {
			return errors.New("post不存在~")
		}
	}

	return nil
}

//post是否为自己创建的
func IfPostOwned(pid, userid uint) error {
	one, err := dao.Post.Where("id = ?", pid).FindOne()
	if err != nil {
		return err
	}
	if one != nil {
		if one.Userid != userid {
			return errors.New("无权删除~")
		}
	}

	return nil
}

//是否频繁发帖
func IfPostOften(title string, userid uint) error {
	if _, err := dao.Post.Where("userid = ? and deleted = ?", userid, 0).Order("id desc").FindOne("title = ?", title); err != nil {
		return err
	} else {
		findOne, _ := dao.Post.Where("userid = ? and deleted = ?", userid, 0).Order("CreateAt desc").FindOne()
		if findOne != nil {
			if time.Now().Unix()-findOne.CreateAt.Timestamp() < 10 {
				fmt.Println(time.Now().Unix() - findOne.CreateAt.Timestamp())
				return errors.New("10秒内已经发过贴,请等待10秒~")
			}
		}
	}
	return nil
}
