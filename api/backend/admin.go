package backend

import (
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/v2/frame/g"
)

type AdminShowReq struct {
	g.Meta `path:"/admin/{id}" method:"get" tags:"用户" summary:"创建内容接口"`
	Id     uint `json:"id"   in:"path"    v:"required#请输入id"      dc:"管理员Id"`
}
type AdminShowRes struct {
	Id        uint        `json:"id"`         // 自增ID
	Name      string      `json:"name"`       // 姓名
	IsAdmin   bool        `json:"is_admin"`   //是否为管理员
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}
