package controller

import (
	"context"
	"gf-demo/api/frontend"
	"gf-demo/internal/model"
	"gf-demo/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

// Order 订单管理
var Order = cOrder{}

type cOrder struct{}

// 下单
func (c *cOrder) Add(ctx context.Context, req *frontend.AddOrderReq) (res *frontend.AddOrderRes, err error) {
	orderAddInput := model.OrderAddInput{}
	//注意：这里要用scan 而不是struct
	if err = gconv.Scan(req, &orderAddInput); err != nil {
		return nil, err
	}

	addRes, err := service.Order().Add(ctx, orderAddInput)
	if err != nil {
		return nil, err
	}

	return &frontend.AddOrderRes{
		Id: addRes.Id,
	}, err
}
