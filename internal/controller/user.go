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
