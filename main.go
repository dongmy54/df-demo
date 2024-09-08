package main

import (
	_ "gf-demo/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"

	_ "gf-demo/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"gf-demo/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
