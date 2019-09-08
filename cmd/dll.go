package cmd

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func Run() {
	app := cli.NewApp()
	app.Name = "dll"
	app.Usage = "code generate!"
	app.Version = "1.0.0"
	app.Action = func(c *cli.Context) error {
		if c.NArg() == 0 {
			fmt.Println(PathArg(c.Args()))
			fmt.Printf("未找到命令: %s\n运行命令 %s help 获取帮助\n", c.Args().First(), app.Name)
			return nil
		}
		return nil
	}
	//设置多个命令，不同的命令执行不同的操作
	app.Commands = []cli.Command{
		NewCMD,     //new
		ModuleCMD,  //module
		ConfigCMD,  //config
		HandlerCMD, //handler
		ServiceCMD, //service
		ModelCMD,   //model
		ProtoCMD,   //proto
	}
	app.Run(os.Args)
}
