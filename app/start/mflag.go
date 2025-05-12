package start

import (
	"github.com/lwmacct/250300-go-app-template/app"

	"github.com/lwmacct/250300-go-mod-mflag/pkg/mflag"
	"github.com/lwmacct/250300-go-mod-mlog/pkg/mlog"
	"github.com/spf13/cobra"
)

type cmd struct{}

var t *cmd

func init() {
	t = &cmd{}
}

func (t *cmd) run(args []string) {
	defer mlog.Close()

	// 定义可用任务映射
	tasks := map[string]func(){
		"job0": func() { mlog.Info(mlog.H{"msg": "job0"}) },
		"job1": func() { mlog.Info(mlog.H{"msg": "job1"}) },
		"job2": func() { mlog.Info(mlog.H{"msg": "job2"}) },
	}

	// 构建提示信息（任务列表）
	var availableTasks []string
	for task := range tasks {
		availableTasks = append(availableTasks, task)
	}

	if len(args) == 0 {
		mlog.Error(mlog.H{"msg": "需要指定任务", "可选任务": availableTasks})
		return
	}

	task := args[0]
	if fn, exists := tasks[task]; exists {
		fn()
	} else {
		mlog.Error(mlog.H{"msg": "需要指定任务", "可选任务": availableTasks})
	}
}

func Cmd() *mflag.Ts {
	mc := mflag.New(app.Flag).UsePackageName("")
	mc.AddCmd(func(cmd *cobra.Command, args []string) {
		t.run(args)
	}, "run", "", "app", "mlog")
	return mc
}
