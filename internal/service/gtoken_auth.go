// 这个是自己定义的 不是自动生成的
// gtoken 授权
package service

import (
	"fmt"
	"gf-demo/internal/consts"
	"net/http"

	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// 登录认证函数
func GtokenLoginFuc(r *ghttp.Request) (string, interface{}) {
	username := r.Get("username").String()
	passwd := r.Get("passwd").String()

	if username == "" || passwd == "" {
		r.Response.WriteJson(gtoken.Fail("账号或密码错误."))
		r.ExitAll()
	}
	// 唯一标识，扩展参数user data
	return username, "1"
}

// 登录后处理返回函数 非必需
// func LoginAfterFuc(r *ghttp.Request, resData gtoken.Resp) {
// 	//g.Dump("=========:", gfToken.GetTokenData(r))
// 	g.Dump("+=============:resData:", resData)

// 	// 用于自定义返回值 如果没有定义 接口看不到返回哦
// 	if !resData.Success() {
// 		r.Response.WriteJson(resData)
// 	} else {
// 		r.Response.WriteJson((g.Map{
// 			"Custom": "我自己定义的",
// 			"result": resData,
// 		}))
// 	}
// }

// 授权之后做一些操作 比如存下这个人的信息到ctx中
func GtokenAuthAfterFuc(r *ghttp.Request, respData gtoken.Resp) {
	if respData.Success() {
		r.SetCtxVar(consts.CtxAdminId, 2) // 这里先随便写一个数2
		// 授权完成后
		g.Dump("=======respData:", respData)
		r.Middleware.Next()
	} else {
		var params map[string]interface{}
		if r.Method == http.MethodGet {
			params = r.GetMap()
		} else if r.Method == http.MethodPost {
			params = r.GetMap()
		} else {
			r.Response.Writeln("request method error")
			return
		}

		no := gconv.String(gtime.TimestampMilli())

		g.Log().Warning(r.Context(), fmt.Sprintf("[AUTH_%s][url:%s][params:%s][data:%s]",
			no, r.URL.Path, params, respData.Json()))
		respData.Msg = "授权失败"
		r.Response.WriteJson(respData)
		r.ExitAll()
	}
}
