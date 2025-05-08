package app

import "github.com/lwmacct/250300-go-mod-mlog/pkg/mlog"

type TsFlag struct {
	Label []string `group:"app" note:"标签, 可用于区分不同的服务" default:""`
	Log   mlog.Opts
	Start struct{} `group:"start" note:"默认配置"`

	App struct {
		ListenAddr string   `group:"app" note:"Http 监听地址" default:"0.0.0.0:8888"`
		Plugin     []string `group:"app" note:"插件, 启用的插件列表" default:""`
	}

	Server struct {
		ListenAddr string `group:"server" note:"监听地址" default:"0.0.0.0:8888"`
	}

	Client struct {
		ListenAddr string `group:"client" note:"监听地址" default:"0.0.0.0:8888"`
	}
}
