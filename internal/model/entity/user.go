// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id        int64       `json:"id"        orm:"id"         description:""`
	Name      string      `json:"name"      orm:"name"       description:"用户姓名"`
	Gender    int         `json:"gender"    orm:"gender"     description:"用户性别"`
	Mobile    string      `json:"mobile"    orm:"mobile"     description:"用户电话"`
	Password  string      `json:"password"  orm:"password"   description:"用户密码"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:""`
}
