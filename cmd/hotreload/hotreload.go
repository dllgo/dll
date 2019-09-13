package hotreload

import (
	"flag"
	"fmt"
	"os"

	"github.com/dllgo/dll/cmd/hotreload/runner"
	"github.com/urfave/cli"
)

var ReloadCMD = cli.Command{
	//命令全称
	Name: "Run",
	//命令简写
	Aliases: []string{"run"},
	//命令详细描述
	Usage: "dll run -h",
	//命令处理函数
	Action: func(c *cli.Context) {
		reloadCMD()
	},
}

func reloadCMD() {
	configPath := flag.String("c", "", "config file path")
	flag.Parse()

	if *configPath != "" {
		if _, err := os.Stat(*configPath); err != nil {
			fmt.Printf("Can't find config file `%s`\n", *configPath)
			os.Exit(1)
		} else {
			os.Setenv("RUNNER_CONFIG_PATH", *configPath)
		}
	}

	runner.Start()
}
