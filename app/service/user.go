package service

import (
	"context"
	"errors"
)

var User = new(UserService)

type UserService struct{}

//注册
func (s *UserService) Signup(ctx context.Context) error {

	return errors.New("")
}

//登录
func (s *UserService) Login(ctx context.Context) error {
	return errors.New("")

}

//退出
func (s *UserService) Logout(ctx context.Context) error {

	return errors.New("")
}
