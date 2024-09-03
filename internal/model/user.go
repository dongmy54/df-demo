package model

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
