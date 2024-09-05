package model

// UserLoginInput 用户登录
// 执行登录后 信息都存到session中 因此无需返回
type LoginInput struct {
	Mobile   string // 账号
	Password string // 密码(明文)
}
