package controller

import (
	"context"

	v1 "gf-demo/api/hello/v1"
	"gf-demo/internal/model"
	"gf-demo/internal/service"
)

// Content 内容管理
var User = cUser{}

type cUser struct{}

func (a *cUser) Create(ctx context.Context, req *v1.UserCreateReq) (res *v1.UserCreateRes, err error) {
	out, err := service.User().Create(ctx, model.UserCreateInput{
		UserCreateUpdateBase: model.UserCreateUpdateBase{
			Name:     req.Name,
			Mobile:   req.Mobile,
			Gender:   req.Gender,
			Password: req.Password,
		},
	})
	if err != nil {
		return nil, err
	}
	return &v1.UserCreateRes{UserId: out.UserId}, nil
}

func (a *cUser) Update(ctx context.Context, req *v1.UserUpdateReq) (res *v1.UserUpdateRes, err error) {
	err = service.User().Update(ctx, model.UserUpdateInput{
		UserCreateUpdateBase: model.UserCreateUpdateBase{
			Name:     req.Name,
			Mobile:   req.Mobile,
			Gender:   req.Gender,
			Password: req.Password,
		},
		UserId: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UserUpdateRes{UserId: req.Id}, nil
}
