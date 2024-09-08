package controller

import (
	"context"

	"gf-demo/api/frontend"
	"gf-demo/internal/dao"
	"gf-demo/internal/model"
	"gf-demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

// Collection 内容管理
var Collection = cCollection{}

type cCollection struct{}

func (a *cCollection) List(ctx context.Context, req *frontend.CollectionListReq) (res *frontend.CollectionListRes, err error) {
	out, err := service.Collection().GetList(ctx, model.CollectionGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}

	g.Dump("+============:", out)
	r, _ := dao.CollectionInfo.Ctx(ctx).One()
	g.Dump("===========:", r)
	gconv.Scan(out, &res) // 第一个参数数据源  第二个参数响应数据
	return res, nil
}
