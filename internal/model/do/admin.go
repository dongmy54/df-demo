// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Admin is the golang structure of table admin for DAO operations like Where/Data.
type Admin struct {
	g.Meta    `orm:"table:admin, do:true"`
	Id        interface{} //
	Name      interface{} // 姓名
	Password  interface{} // 密码
	IsAdmin   interface{} // 是否为超级管理员
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time // 是否删除
}
