package cmd

import (
	"context"

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
				)
			})
			s.Group("/api/backend", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(backend.Auth)
			})

			s.Group("/backend", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(controller.Admin)
			})

			s.Run()
			return nil
		},
	}
)
