package model

import (
	"github.com/gogf/gf/os/gtime"
)

// UserCreateUpdateBase 创建/修改内容基类
type UserCreateUpdateBase struct {
	Name     string // 用户姓名
	Gender   int    // 性别
	Mobile   string // 手机号
	Password string // 密码
}

// UserCreateInput 创建内容
type UserCreateInput struct {
	UserCreateUpdateBase
	UserId uint
}

// UserCreateOutput 创建内容返回结果
type UserCreateOutput struct {
	UserId uint `json:"user_id"`
}

type UserUpdateInput struct {
	UserCreateUpdateBase
	UserId uint
}

type UserGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
}

type UserGetListOutput struct {
	List  []UserListItem `json:"list" description:"用户列表"`
	Page  int            `json:"page" description:"分页码"`
	Size  int            `json:"size" description:"分页数量"`
	Total int            `json:"total" description:"数据总数"`
}

type UserListItem struct {
	Id        uint        `json:"id"`         // 自增ID
	Name      string      `json:"name"`       // 姓名
	Mobile    string      `json:"mobile"`     // 手机号
	Gender    string      `json:"gender"`     // 性别
	Password  string      `json:"password"`   // 密码
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}
