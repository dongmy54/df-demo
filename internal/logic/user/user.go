package User

import (
	"context"

	"gf-demo/internal/model/entity"
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

// Update User
func (s *sUser) Update(ctx context.Context, in model.UserUpdateInput) (err error) {
	_, err = dao.User.
		Ctx(ctx).
		Data(in).
		FieldsEx(dao.User.Columns().Id).
		Where(dao.User.Columns().Id, in.UserId).
		OmitEmpty(). // 忽略空的字段
		Update()

	if err != nil {
		return err
	}
	return
}

// GetList 查询内容列表
func (s *sUser) GetList(ctx context.Context, in model.UserGetListInput) (out *model.UserGetListOutput, err error) {
	var (
		m = dao.User.Ctx(ctx)
	)
	out = &model.UserGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	// 分配查询
	listModel := m.Page(in.Page, in.Size)

	// 执行查询
	var list []*entity.User
	if err := listModel.Scan(&list); err != nil {
		return out, err
	}
	// 没有数据
	if len(list) == 0 {
		return out, nil
	}
	out.Total, err = m.Count()
	if err != nil {
		return out, err
	}
	// User
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}
