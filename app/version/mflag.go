package version

import (
	"fmt"
	"runtime"

	"github.com/lwmacct/250300-go-mod-mflag/pkg/mflag"
	"github.com/spf13/cobra"
)

var (
	AppRawName string = "Unknown" // 应用名称
	AppProject string = "Unknown" // 项目名称
	AppVersion string = "Unknown" // 应用版本
	GitCommit  string = "Unknown" // Git提交
	BuildTime  string = "Unknown" // 编译时间
	Developer  string = "Unknown" // 开发者
	Workspace  string = "Unknown" // 工作目录, 用于去除callstack中的绝对路径
)

func Cmd() *mflag.Ts {
	mc := mflag.New(nil).UsePackageName("")
	mc.AddCmd(func(cmd *cobra.Command, args []string) {
		run(cmd, args)
	}, "run", "", "")
	return mc
}

func run(cmd *cobra.Command, args []string) {
	_ = map[string]any{"cmd": cmd, "args": args}
	fmt.Printf("AppRawName:   %s\n", AppRawName)
	fmt.Printf("AppVersion:   %s\n", AppVersion)
	fmt.Printf("Go Version:   %s\n", runtime.Version())
	fmt.Printf("Git Commit:   %s\n", GitCommit)
	fmt.Printf("Build Time:   %s\n", BuildTime)
	fmt.Printf("AppProject:   %s\n", AppProject)
	fmt.Printf("Developer :   %s\n", Developer)

}
