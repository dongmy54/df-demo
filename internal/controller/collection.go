package controller

import (
	"context"

	"gf-demo/api/frontend"
	"gf-demo/internal/model"
	"gf-demo/internal/service"

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

	// 会报错的 因为这里res还是空指针 还没有初始化呢
	// 所以这里调用&res找出这个指针的地址（PS：虽然res是一个空指针，指向的地址是空；但是它本身是有地址的，这个地址也在内存中所以&res能取到哦）
	gconv.Scan(out, &res) // 第一个参数数据源  第二个参数响应数据

	return res, nil
}
