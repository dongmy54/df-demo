// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"gf-demo/internal/model/entity"
)

type (
	ISession interface {
		// 设置用户Session.
		SetUser(ctx context.Context, user *entity.User) error
		// 获取当前登录的用户信息对象，如果用户未登录返回nil。
		GetUser(ctx context.Context) *entity.User
		// 删除用户Session。
		RemoveUser(ctx context.Context) error
		// 设置LoginReferer.
		SetLoginReferer(ctx context.Context, referer string) error
		// 获取LoginReferer.
		GetLoginReferer(ctx context.Context) string
		// 删除LoginReferer.
		RemoveLoginReferer(ctx context.Context) error
	}
)

var (
	localSession ISession
)

func Session() ISession {
	if localSession == nil {
		panic("implement not found for interface ISession, forgot register?")
	}
	return localSession
}

func RegisterSession(i ISession) {
	localSession = i
}
