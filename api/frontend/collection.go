package frontend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CollectionListReq struct {
	g.Meta `path:"/collection/list" method:"get" tags:"收藏" summary:"创建订单"`
	Page   int `json:"page"`
	Size   int `json:"size"`
}

type CollectionListRes struct {
	List  interface{} `json:"list"`
	Page  int         `json:"page" description:"分页 页码"`
	Size  int         `json:"size" description:"分页 每页数量"`
	Total int         `json:"total" description:"总数量"`
}
