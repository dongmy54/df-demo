package model

import (
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/v2/frame/g"
)

type CollectionGetListInput struct {
	Page int `json:"page" description:"分页 页码"`
	Size int `json:"size" description:"分页 每页数量"`
}

type CollectionGetListOut struct {
	List  []CollectionListOutputItem `json:"list"`
	Page  int                        `json:"page" description:"分页 页码"`
	Size  int                        `json:"size" description:"分页 每页数量"`
	Total int                        `json:"total" description:"总数量"`
}

type CollectionListOutputItem struct {
	Id       int `json:"id"        description:""`
	UserId   int `json:"user_id"    description:"用户id"`
	ObjectId int `json:"object_id"  description:"对象id"`
	Type     int `json:"type"      description:"收藏类型：1商品 2文章"`
	//Goods     GoodsItem   `json:"goods" orm:"with:id=object_id"`
	Article   ArticleItem `json:"article" orm:"with:id=object_id"`
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}

type GoodsItem struct {
	g.Meta `orm:"table:goods_info"` // 注意with使用这里要多加一层
	Id     uint                     `json:"id"`
	Name   string                   `json:"name"`
	PicUrl string                   `json:"pic_url"`
	Price  int                      `json:"price"`
}

type ArticleItem struct {
	g.Meta `orm:"table:article_info"`
	Id     uint   `json:"id"`
	Title  string `json:"title"`
	Desc   string `json:"desc"`
	PicUrl string `json:"pic_url"`
}
