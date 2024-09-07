package Admin

import (
	"context"

	"gf-demo/internal/model"
	"gf-demo/internal/service"

	// 这里要用第二版

	"gf-demo/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
)

type sAdmin struct{}

func init() {
	service.RegisterAdmin(New())
}

func New() *sAdmin {
	return &sAdmin{}
}

// Update Admin
func (s *sAdmin) Get(ctx context.Context, admin_id uint) (output model.AdminShowOutput, err error) {
	g.Dump("=============admin_id:", admin_id)

	err = dao.Admin.
		Ctx(ctx).
		Where(dao.Admin.Columns().Id, admin_id).Scan(&output)

	if err != nil {
		g.Dump("======error:", err)
		return output, err
	}
	return
}
