package controller

import (
	"context"

	v1 "gf-demo/api/hello/v1"
	"gf-demo/internal/model"
	"gf-demo/internal/service"
	redisCache "gf-demo/utility/redis_cache"

	"github.com/gogf/gf/frame/g"
)

// User 内容管理
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

	g.Log().Debug(ctx, "hello word") // 日志
	return &v1.UserUpdateRes{UserId: req.Id}, nil
}

func (a *cUser) GetList(ctx context.Context, req *v1.UserGetListReq) (res *v1.UserGetListRes, err error) {
	// 缓存测试
	g.Dump("===========redisCache test:", redisCache.Cache.MustGet(ctx, "test:cache:key").String())
	out, err := service.User().GetList(ctx, model.UserGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}

	// 这里能获取到它们的信息
	g.Log().Debug("=====测试当前用户Session信息的返回：======", service.Session().GetUser(ctx))
	g.Log().Debug("=====测试当前用户Session信息的返回：======", service.BizCtx().Get(ctx).User)

	return &v1.UserGetListRes{
		List:  out.List,
		Page:  out.Page,
		Size:  out.Size,
		Total: out.Total,
	}, nil
}

func (a *cUser) Delete(ctx context.Context, req *v1.UserDeleteReq) (res *v1.UserDeleteRes, err error) {
	err = service.User().Delete(ctx, req.Id)
	return
}
