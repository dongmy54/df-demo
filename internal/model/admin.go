package model

import "github.com/gogf/gf/os/gtime"

type AdminListItem struct {
	Id        uint        `json:"id"`         // 自增ID
	Name      string      `json:"name"`       // 姓名
	IsAdmin   bool        `json:"is_admin"`   //是否为管理员
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}

type AdminShowOutput struct {
	AdminListItem
}
