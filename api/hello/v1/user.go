package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UserCreateReq struct {
	g.Meta   `path:"/user/create" method:"post" tags:"用户" summary:"创建内容接口"`
	Name     string `json:"name"       v:"required#请输入姓名"      dc:"姓名"`
	Gender   int    `json:"gender"     dc:"性别"`
	Mobile   string `json:"mobile"     v:"required#请输入手机号"    dc:"手机号"`
	Password string `json:"password"   v:"required#请输入密码"      dc:"标题"`
}
type UserCreateRes struct {
	UserId uint `json:"userId"`
}

type UserUpdateReq struct {
	g.Meta   `path:"/user/{ID}" method:"put" tags:"用户" summary:"更新用户信息"`
	Id       uint   `json:"id"         v:"required#请输用户Id"      dc:"用户id"`
	Name     string `json:"name"            dc:"姓名"`
	Gender   int    `json:"gender"     dc:"性别"`
	Mobile   string `json:"mobile"         dc:"手机号"`
	Password string `json:"password"       dc:"标题"`
}

type UserUpdateRes struct {
	UserId uint `json:"userId"`
}
