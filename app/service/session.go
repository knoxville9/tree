package service

import (
	"context"
	"lqc.com/tree/app/model"
)

var SessionKey = "SessionKey"

var Session = new(sessionService)

type sessionService struct {
}

//退出时删除session
func (s *sessionService) Remove(ctx context.Context) error {
	get := Context.Get(ctx)
	get.Session.Remove(SessionKey)
	return nil

}

//获取用户信息时,可以调用该接口
func (s *sessionService) GetUser(ctx context.Context) *model.User {
	get := Context.Get(ctx)
	user := get.User

	return user

}

//登录时在context中设置用户的session
func (s *sessionService) SetUser(ctx context.Context, req *model.UserDoLogInReq) {
	get := Context.Get(ctx).User
	get.Passport = req.Passport

}
