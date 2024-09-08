package position

import (
	"context"
	"gf-demo/internal/dao"
	"gf-demo/internal/model"
	"gf-demo/internal/service"

	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type sOrder struct{}

func init() {
	service.RegisterOrder(New())
}

func New() *sOrder {
	return &sOrder{}
}

// 下单
func (s *sOrder) Add(ctx context.Context, in model.OrderAddInput) (out *model.OrderAddOutput, err error) {
	in.Number = gtime.Datetime()
	out = &model.OrderAddOutput{}
	//官方建议的事务闭包处理 注意这里版本不同会导致写法不一样
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		//生成主订单
		lastInsertId, err := dao.OrderInfo.Ctx(ctx).InsertAndGetId(in)
		if err != nil {
			return err
		}
		//生成商品订单
		for _, info := range in.OrderAddGoodsInfos {
			info.OrderId = gconv.Int(lastInsertId)
			_, err := dao.OrderGoodsInfo.Ctx(ctx).Insert(info)
			if err != nil {
				return err
			}
		}
		//更新商品销量和库存，todo 后期接入消息
		for _, info := range in.OrderAddGoodsInfos {
			//商品增加销量
			_, err := dao.GoodsInfo.Ctx(ctx).WherePri(info.GoodsId).Increment(dao.GoodsInfo.Columns().Sale, info.Count)
			if err != nil {
				return err
			}
			//商品减少库存
			_, err2 := dao.GoodsInfo.Ctx(ctx).WherePri(info.GoodsId).Decrement(dao.GoodsInfo.Columns().Stock, info.Count)
			if err2 != nil {
				return err
			}
		}
		out.Id = uint(lastInsertId)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return
}
