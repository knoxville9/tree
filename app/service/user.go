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

	user, err := dao.User.FindOne("password=? and passport=?", req.Password, req.Passport)

	if err != nil {
		return err
	} else {
		if user == nil {
			return errors.New("账号或密码错误")
		}
	}

	Context.SetUser(ctx, user)
	if err := Context.Get(ctx).Session.Set(SessionKey, user); err != nil {
		return err
	}

	return nil

}

//退出
func (s *UserService) Logout(ctx context.Context) error {
	if err := Context.Get(ctx).Session.Remove(SessionKey); err != nil {
		return err
	}
	return nil
}

//返回用户信息
func (s *UserService) Profile(ctx context.Context) (*model.User, error) {
	get := Context.Get(ctx)
	user := get.Session.GetVar(SessionKey)
	if !user.IsNil() {
		if a, b := user.Interface().(*model.User); b {
			return a, nil
		}
	}
	return nil, errors.New("信息已清空")
}
