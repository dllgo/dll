package cmd

import (
	"bytes"
	"fmt"

	tp "github.com/dllgo/dll/template"
	"github.com/urfave/cli"
)

var HandlerCMD = cli.Command{
	//命令全称
	Name: "handler",
	//命令简写
	Aliases: []string{"hdl"},
	//命令详细描述
	Usage: "dll handler[hdl]  name pkgname path",
	//命令处理函数
	Action: func(c *cli.Context) {
		handlerCMD(c)
	},
}

//handler
func handlerCMD(c *cli.Context) {
	if c.NArg() == 0 {
		fmt.Println("请输入 handler 名称")
		return
	}
	filepath, _ := GetFilePath(PathArg(c.Args()), StructName(c.Args())+"Handler.go")
	WriteFile(filepath, handlerTemplate(c.Args()))
	fmt.Printf("handler=%v \n", c.Args().First())
}

func handlerTemplate(args []string) string {
	var b bytes.Buffer
	tp.HandlerTmpl.Execute(&b, &InputData{
		PkgName: PkgName(args),
		Name:    StructName(args),
	})
	c := b.String()
	return c
}
