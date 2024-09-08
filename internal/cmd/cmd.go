package cmd

import (
	"context"

	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"gf-demo/internal/controller"
	"gf-demo/internal/controller/backend"
	"gf-demo/internal/controller/hello"
	"gf-demo/internal/service"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/api/v1", func(group *ghttp.RouterGroup) {
				group.Middleware(
					ghttp.MiddlewareHandlerResponse,
					service.Middleware().Ctx,
				)
				group.Bind(
					hello.NewV1(),
					controller.User, // 这里其实是一组多个url组成的路由
					controller.Login,
					controller.Order,
				)
			})
			s.Group("/api/backend", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(backend.Auth)
			})

			// 启动gtoken只对下面接口
			gfToken := &gtoken.GfToken{
				ServerName:      "my gf demo",
				LoginPath:       "/login", // 这样写了后就会自动添加这个路径的请求接口做认证
				LoginBeforeFunc: service.GtokenLoginFuc,
				AuthAfterFunc:   service.GtokenAuthAfterFuc,
				LogoutPath:      "/user/logout",
				//AuthExcludePaths: g.SliceStr{"/user/info", "/system/user/info"}, // 不拦截路径 /user/info,/system/user/info,/system/user,
				MultiLogin: true,
			}

			s.Group("/backend", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					controller.File,
				)

				err := gfToken.Middleware(ctx, group)
				if err != nil {
					panic(err)
				}

				group.Bind(controller.Admin)
			})

			s.Run()
			return nil
		},
	}
)
