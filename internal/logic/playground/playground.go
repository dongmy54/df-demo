package playground

import (
	"gf-demo/internal/model/entity"
	"gf-demo/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type sPlayground struct{}

func init() {
	service.RegisterPlayground(New())
}

func New() *sPlayground {
	return &sPlayground{}
}

// 组合信息
type UserInfo struct {
	User       *entity.User
	UserDetail *entity.UserDetail   // 一对一
	UserScores []*entity.UserScores // 一对多
}

// 这是我自己定义的一个练写方法
func (s *sPlayground) Do() {
	// 定义用户列表
	var users []UserInfo
	// 查询用户基础数据
	// SELECT * FROM `user`
	err := g.Model("user").ScanList(&users, "User")
	if err != nil {
		panic(err)
	}
	// 查询用户详情数据
	// SELECT * FROM `user_detail` WHERE `uid` IN(1,2)
	err = g.Model("user_detail").
		Where("user_id", gdb.ListItemValuesUnique(users, "User", "Id")).
		ScanList(&users, "UserDetail", "User", "user_id:id")
	if err != nil {
		panic(err)
	}
	// 查询用户学分数据
	// SELECT * FROM `user_scores` WHERE `uid` IN(1,2)
	err = g.Model("user_scores").
		Where("user_id", gdb.ListItemValuesUnique(users, "User", "Id")).
		ScanList(&users, "UserScores", "User", "user_id:id")

	g.Dump("=======当前users:", users)
	if err != nil {
		panic(err)
	}
}
