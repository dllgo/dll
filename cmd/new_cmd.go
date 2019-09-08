package cmd

import (
	"github.com/urfave/cli"
)

var NewCMD = cli.Command{
	//命令全称
	Name: "new",
	//命令简写
	Aliases: []string{"new"},
	//命令详细描述
	Usage: "dll new -h",
	//命令处理函数
	Subcommands: []cli.Command{
		ApiCMD,
	},
}
