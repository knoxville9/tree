package service

import "errors"

var User = new(UserService)

type UserService struct{}

//注册
func (s *UserService) Signup() error {
	return errors.New("")
}

//登录
func (s *UserService) Login() error {
	return errors.New("")

}

//退出
func (s *UserService) Logout() error {
	return errors.New("")
}
