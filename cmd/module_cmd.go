package cmd

import (
	"fmt"

	"github.com/urfave/cli"
)

var ModuleCMD = cli.Command{
	//命令全称
	Name: "module",
	//命令简写
	Aliases: []string{"mod"},
	//命令详细描述
	Usage: "dll module[mod]  name pkgname path",
	//命令处理函数
	Action: func(c *cli.Context) {
		moduleCMD(c)
	},
}

//module
func moduleCMD(c *cli.Context) {
	if c.NArg() == 0 {
		fmt.Println("请输入 module 名称")
		return
	}
	dir := "/" + StructName(c.Args()) + "/" + ToCamel(c.Args().First())
	//handler
	handlerpath, _ := GetFilePath(PathArg(c.Args()), dir+"Handler.go")
	WriteFile(handlerpath, handlerTemplate(c.Args()))
	//service
	servicepath, _ := GetFilePath(PathArg(c.Args()), dir+"Service.go")
	WriteFile(servicepath, serviceTemplate(c.Args()))
	//model
	modelpath, _ := GetFilePath(PathArg(c.Args()), dir+"Model.go")
	WriteFile(modelpath, modelTemplate(c.Args()))
	//proto
	protopath, _ := GetFilePath(PathArg(c.Args()), dir+".proto")
	WriteFile(protopath, protoTemplate(c.Args()))
	fmt.Printf("module=%v \n", c.Args().First())
}
