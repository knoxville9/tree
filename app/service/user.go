package service

import (
	"context"
	"errors"
	"github.com/gogf/gf/crypto/gmd5"
	"lqc.com/tree/app/dao"
	"lqc.com/tree/app/model"
)

var User = new(UserService)

type UserService struct{}

//注册
func (s *UserService) Signup(ctx context.Context, req *model.UserDoSignUpReq) error {
	//如何昵称没有就用用户名
	if req.Nickname == "" {
		req.Nickname = req.Passport

	}
	//密码加密
	encrypt, _ := gmd5.Encrypt(req.Password)
	req.Password = encrypt
	//查看当前注册用户名是否存在
	one, _ := dao.User.FindCount("passport = ?", req.Passport)
	if one != 0 {
		return errors.New("账号已存在")

	}
	//保存用户
	if _, err := dao.User.Save(req); err != nil {
		return errors.New(err.Error())
	}
	return nil

}

//登录
func (s *UserService) Login(ctx context.Context, req *model.UserDoLogInReq) error {

	encrypt, _ := gmd5.Encrypt(req.Password)
	req.Password = encrypt

	if user, err := dao.User.FindOne("password=? and passport=?", req.Password, req.Passport); err != nil {
		return err
	} else {
		if user == nil {
			return errors.New("账号或密码错误")
		}
	}
	Session.SetUser(ctx, req)
	return nil

}

//退出
func (s *UserService) Logout(ctx context.Context) error {
	return Session.Remove(ctx)

}
