package backend

import (
	"context"
	"gf-demo/api/backend"
	"gf-demo/internal/logic/auth"

	"github.com/gogf/gf/v2/frame/g"
)

type authController struct{}

var Auth = authController{}

func (c *authController) Login(ctx context.Context, req *backend.AuthLoginReq) (res *backend.AuthLoginRes, err error) {
	res = &backend.AuthLoginRes{}
	res.Token, res.Expire = auth.Auth().LoginHandler(ctx)
	g.Log().Debug(ctx, res.Token)
	return
}

// func (c *authController) RefreshToken(ctx context.Context, req *api.AuthRefreshTokenReq) (res *api.AuthRefreshTokenRes, err error) {
// 	res = &api.AuthRefreshTokenRes{}
// 	res.Token, res.Expire = service.Auth().RefreshHandler(ctx)
// 	return
// }

// func (c *authController) Logout(ctx context.Context, req *api.AuthLogoutReq) (res *api.AuthLogoutRes, err error) {
// 	service.Auth().LogoutHandler(ctx)
// 	return
// }
