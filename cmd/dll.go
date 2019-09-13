package cmd

import (
	"fmt"
	"os"

	"github.com/dllgo/dll/cmd/gateway"
	"github.com/dllgo/dll/cmd/hotreload"
	"github.com/dllgo/dll/cmd/micro"
	"github.com/dllgo/dll/cmd/proto"
	"github.com/urfave/cli"
)

func Run() {
	app := cli.NewApp()
	app.Name = "dll"
	app.Usage = "code generate!"
	app.Version = "1.0.0"
	app.Action = func(c *cli.Context) error {
		if c.NArg() == 0 {
			fmt.Printf("未找到命令: %s\n运行命令 %s help 获取帮助\n", c.Args().First(), app.Name)
			return nil
		}
		return nil
	}
	//设置多个命令，不同的命令执行不同的操作
	app.Commands = []cli.Command{
		micro.MicroCMD,      //micro
		proto.ProtoCMD,      //proto
		gateway.GWCMD,       //gateway
		hotreload.ReloadCMD, //hotreload
	}
	app.Run(os.Args)
}
