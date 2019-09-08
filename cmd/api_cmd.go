package cmd

import (
	"fmt"

	"github.com/urfave/cli"
)

var ApiCMD = cli.Command{
	//命令全称
	Name: "api",
	//命令简写
	Aliases: []string{"api"},
	//命令详细描述
	Usage: "dll new api name pkgname path",
	//命令处理函数
	Action: func(c *cli.Context) {
		apiCMD(c)
	},
}

//api
func apiCMD(c *cli.Context) {
	if c.NArg() == 0 {
		fmt.Println("请输入 api 名称")
		return
	}
	fmt.Printf("api=%v \n", c.Args().First())
}
