package controller

import (
	"context"

	"gf-demo/api/backend"
	"gf-demo/internal/service"
)

// Admin 内容管理
var Admin = cAdmin{}

type cAdmin struct{}

// 方法名区分后端和前端
func (a *cAdmin) BackendShow(ctx context.Context, req *backend.AdminShowReq) (res *backend.AdminShowRes, err error) {
	out, err := service.Admin().Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &backend.AdminShowRes{
		Id:        out.Id,
		Name:      out.Name,
		IsAdmin:   out.IsAdmin,
		CreatedAt: out.CreatedAt,
		UpdatedAt: out.UpdatedAt,
	}, nil
}
