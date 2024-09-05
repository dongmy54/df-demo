// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"gf-demo/internal/model"
	"gf-demo/internal/model/entity"
)

type (
	ILogin interface {
		// 执行登录
		Login(ctx context.Context, in model.LoginInput) error
		// 根据账号和密码查询用户信息，一般用于账号密码登录。
		// 注意password参数传入的是按照相同加密算法加密过后的密码字符串。
		GetLoginByPassportAndPassword(ctx context.Context, mobile string, password string) (user *entity.User, err error)
	}
)

var (
	localLogin ILogin
)

func Login() ILogin {
	if localLogin == nil {
		panic("implement not found for interface ILogin, forgot register?")
	}
	return localLogin
}

func RegisterLogin(i ILogin) {
	localLogin = i
}
