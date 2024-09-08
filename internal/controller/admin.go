package controller

import (
	"context"

	"gf-demo/api/backend"
	"gf-demo/internal/consts"
	"gf-demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

// Admin 内容管理
var Admin = cAdmin{}

type cAdmin struct{}

// 方法名区分后端和前端
func (a *cAdmin) BackendShow(ctx context.Context, req *backend.AdminShowReq) (res *backend.AdminShowRes, err error) {
	out, err := service.Admin().Get(ctx, req.Id)

	// 这里是获取到前面r.SetCtxVar中的值
	admin_id := gconv.Int(ctx.Value(consts.CtxAdminId))
	g.Dump("===============:", admin_id)
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
