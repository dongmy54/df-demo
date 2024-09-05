package controller

import (
	"context"
	v1 "gf-demo/api/hello/v1"
	"gf-demo/internal/model"
	"gf-demo/internal/service"
)

// 登录管理
var Login = cLogin{}

type cLogin struct{}

func (a *cLogin) Login(ctx context.Context, req *v1.LoginDoReq) (res *v1.LoginDoRes, err error) {
	res = &v1.LoginDoRes{}

	err = service.Login().Login(ctx, model.LoginInput{
		Mobile:   req.Mobile,
		Password: req.Password,
	})
	if err != nil {
		return
	}

	res.Info = service.Session().GetUser(ctx)
	return
}
