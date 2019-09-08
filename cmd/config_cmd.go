package cmd

import (
	"fmt"

	"github.com/urfave/cli"
)

var ConfigCMD = cli.Command{
	//命令全称
	Name: "config",
	//命令简写
	Aliases: []string{"cfg"},
	//命令详细描述
	Usage: "dll config[cfg] name pkgname path",
	//命令处理函数
	Action: func(c *cli.Context) {
		configCMD(c)
	},
}

//config
func configCMD(c *cli.Context) {
	if c.NArg() == 0 {
		fmt.Println("请输入 config 名称")
		return
	}
	fmt.Printf("config=%v \n", c.Args().First())
}
