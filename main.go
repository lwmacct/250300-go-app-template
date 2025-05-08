package main

import (
	"fmt"
	"os"

	"github.com/lwmacct/250300-go-app-template/app"
	"github.com/lwmacct/250300-go-app-template/app/client"
	"github.com/lwmacct/250300-go-app-template/app/server"
	"github.com/lwmacct/250300-go-app-template/app/start"
	"github.com/lwmacct/250300-go-app-template/app/version"

	"github.com/lwmacct/250300-go-mod-mflag/pkg/mflag"
	"github.com/lwmacct/250300-go-mod-mlog/pkg/mlog"
)

var mc *mflag.Ts

func main() {
	mc = mflag.New(nil)

	{
		// 命令行参数
		mc.AddCobra(version.Cmd().Cobra())

		// 如果程序只有一个命令, 建议使用 start 入口
		mc.AddCobra(start.Cmd().Cobra())

		// 客户端, 当指定的环境变量正确时, 会自动添加此命令, 可以设置自己的 salt
		if os.Getenv("ACF_CLIENT_FLAG") == "1" {
			mc.AddCobra(client.Cmd().Cobra())
		}

		// 服务端, 当指定的环境变量正确时, 会自动添加此命令, 可以设置自己的 salt
		if os.Getenv("ACF_SERVER_FLAG") == "1" {
			mc.AddCobra(server.Cmd().Cobra())
		}

	}

	{
		mlog.SetNew(
			mlog.WithFile(app.Flag.Log.File),
			mlog.WithLevel(app.Flag.Log.Level),
			mlog.WithCallerClip(version.Workspace),
		)
	}

	if err := mc.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
