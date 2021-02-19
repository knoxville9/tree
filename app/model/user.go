// ==========================================================================
// This is auto-generated by gf cli tool. Fill this file as you wish.
// ==========================================================================

package model

import (
	"lqc.com/tree/app/model/internal"
)

// User is the golang structure for table user.
type User internal.User

// Fill with you ideas below.

type UserDoSignUpReq struct {
	Passport string `orm:"passport"   json:"passport" v:"required|passport #|账号字母开头，只能包含字母、数字和下划线，长度在6~18之间"` // 用户账号
	Password string `orm:"password"   json:"password" v:"required|password #|通用密码(任意可见字符，长度在6~18之间)"`         // 用户密码
	Nickname string `orm:"nickname"   json:"nickname" `                                                       // 用户昵称

}

type UserDoLogInReq struct {
	Passport string `orm:"passport"   json:"passport" v:"required|length:6,18"` // 用户账号
	Password string `orm:"password"   json:"-" v:"required|length:6,18"`        // 用户密码

}
