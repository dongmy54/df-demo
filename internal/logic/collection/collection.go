package collection

import (
	"context"
	"gf-demo/internal/dao"
	"gf-demo/internal/model"
	"gf-demo/internal/service"
)

type sCollection struct{}

func init() {
	service.RegisterCollection(New())
}

func New() *sCollection {
	return &sCollection{}
}

// Create 创建内容
func (s *sCollection) GetList(ctx context.Context, in model.CollectionGetListInput) (out model.CollectionGetListOut, err error) {
	// dao.CollectionInfo.Ctx(ctx).With().Page(in.Page, in.Size).Scan(out.List)
	// With用于一对一/一对多 With可限定这里返回的数据范围
	//dao.CollectionInfo.Ctx(ctx).With(model.GoodsItem{}).Scan(&out.List)
	dao.CollectionInfo.Ctx(ctx).Where(dao.CollectionInfo.Columns().Type, 2).With(model.ArticleItem{}).Scan(&out.List)
	return out, err
}
