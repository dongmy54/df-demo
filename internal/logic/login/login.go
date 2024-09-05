package login

import (
	"context"
	"gf-demo/internal/dao"
	"gf-demo/internal/model/entity"

	"gf-demo/internal/service"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/v2/errors/gerror"

	"gf-demo/internal/model"
)

type sLogin struct {
}

func init() {
	user := New()
	service.RegisterLogin(user)
}

func New() *sLogin {
	return &sLogin{}
}

// 执行登录
func (s *sLogin) Login(ctx context.Context, in model.LoginInput) error {
	userEntity, err := s.GetLoginByPassportAndPassword(
		ctx,
		in.Mobile,
		in.Password,
	)
	if err != nil {
		return err
	}
	if userEntity == nil {
		return gerror.New(`账号或密码错误`)
	}
	if err := service.Session().SetUser(ctx, userEntity); err != nil {
		return err
	}
	// 自动更新上线
	service.BizCtx().SetUser(ctx, &model.ContextUser{
		Id:     uint(userEntity.Id),
		Mobile: userEntity.Mobile,
		Name:   userEntity.Name,
	})
	return nil
}

// 根据账号和密码查询用户信息，一般用于账号密码登录。
// 注意password参数传入的是按照相同加密算法加密过后的密码字符串。
func (s *sLogin) GetLoginByPassportAndPassword(ctx context.Context, mobile, password string) (user *entity.User, err error) {
	// 先根据用户的姓名+密码查找
	err = dao.User.Ctx(ctx).Where(g.Map{
		dao.User.Columns().Mobile:   mobile,
		dao.User.Columns().Password: password,
	}).Scan(&user)
	return
}
