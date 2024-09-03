package User

import (
	"context"

	"gf-demo/internal/service"

	"github.com/gogf/gf/v2/encoding/ghtml"

	"gf-demo/internal/dao"
	"gf-demo/internal/model"
)

type sUser struct{}

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

// Create 创建内容
func (s *sUser) Create(ctx context.Context, in model.UserCreateInput) (out model.UserCreateOutput, err error) {
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	lastInsertID, err := dao.User.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.UserCreateOutput{UserId: uint(lastInsertID)}, err
}
