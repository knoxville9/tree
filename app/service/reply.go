package service

import (
	"context"
	"errors"
	"lqc.com/tree/app/dao"
	"lqc.com/tree/app/model"
)

var Reply = new(replyService)

type replyService struct {
}

//回帖
func (s *replyService) Create(ctx context.Context, req *model.ReplyDoCreate) error {
	user := ctx.Value(ContextKey).(*model.Context).User

	if err := IfPostExist(*req.Pid); err != nil {
		return err
	}

	if err := IfPostDeleted(*req.Pid); err != nil {
		return err
	}

	if err := IfReplySame(uint(*req.Pid), user.Id, req.Content); err != nil {
		return err
	}

	req.Userid = int(user.Id)
	if _, err := dao.Reply.Save(req); err != nil {
		return err
	}

	return nil
}

//删除回复
func (s *replyService) Delete(ctx context.Context, req *model.ReplyDoDelete) error {
	user := ctx.Value(ContextKey).(*model.Context).User

	if err := IfPostDeleted(*req.Pid); err != nil {
		return err
	}
	if err := IfPostExist(*req.Pid); err != nil {
		return err
	}
	if err := IfReplyExist(*req.Rid); err != nil {
		return err
	}
	if err := IfReplyDeleted(*req.Rid); err != nil {
		return err
	}
	if err := IfReplyOwned(uint(*req.Rid), user.Id); err != nil {
		return err
	}
	if _, err := dao.Reply.Data("deleted", 1).Where("id = ? and userid =? and pid = ?", req.Rid, user.Id, req.Pid).Update(); err != nil {
		return err
	}

	return nil
}

//reply是否存在
func IfReplyExist(id int) error {
	if count, err := dao.Reply.Where("id = ?", id).FindOne(); err != nil {
		return err
	} else {
		if count == nil {
			return errors.New("回复不存在~")
		}
	}
	return nil
}

//reply是否已经被删除
func IfReplyDeleted(id int) error {
	one, err := dao.Reply.Where("id = ?", id).FindOne()
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

//reply是否为自己创建的
func IfReplyOwned(id, userid uint) error {
	one, err := dao.Reply.Where("id = ?", id).FindOne()
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

//reply是否信息一致
func IfReplySame(pid, userid uint, content string) error {
	one, err := dao.Reply.Where("pid = ? and userid = ?", pid, userid).Order("CreateAt desc").FindOne()
	if err != nil {
		return err
	}
	if one != nil {
		if one.Content == content {
			return errors.New("回复内容一致~请勿重复提交~")
		}

	}

	return nil
}
