// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"gf-demo/internal/model"
)

type (
	IUser interface {
		// Create 创建内容
		Create(ctx context.Context, in model.UserCreateInput) (out model.UserCreateOutput, err error)
		// Update User
		Update(ctx context.Context, in model.UserUpdateInput) (err error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
