package cmd

import (
	"bytes"
	"fmt"

	tp "github.com/dllgo/dll/template"
	"github.com/urfave/cli"
)

var ServiceCMD = cli.Command{
	//命令全称
	Name: "service",
	//命令简写
	Aliases: []string{"srv"},
	//命令详细描述
	Usage: "dll service[srv] name pkgname path",
	//命令处理函数
	Action: func(c *cli.Context) {
		serviceCMD(c)
	},
}

//service
func serviceCMD(c *cli.Context) {
	if c.NArg() == 0 {
		fmt.Println("请输入 service 名称")
		return
	}
	filepath, _ := GetFilePath(PathArg(c.Args()), c.Args().First()+"Service.go")
	WriteFile(filepath, serviceTemplate(c.Args()))
	fmt.Printf("service=%v \n", c.Args())
}

func serviceTemplate(args []string) string {
	var b bytes.Buffer
	tp.ServiceTmpl.Execute(&b, &InputData{
		PkgName: PkgName(args),
		Name:    StructName(args),
	})
	c := b.String()
	return c
}
