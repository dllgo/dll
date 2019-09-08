package cmd

import (
	"bytes"
	"fmt"

	tp "github.com/dllgo/dll/template"
	"github.com/urfave/cli"
)

var ModelCMD = cli.Command{
	//命令全称
	Name: "model",
	//命令简写
	Aliases: []string{"m"},
	//命令详细描述
	Usage: "dll model[m] name pkgname path",
	//命令处理函数
	Action: func(c *cli.Context) {
		modelCMD(c)
	},
}

//model
func modelCMD(c *cli.Context) {
	if c.NArg() == 0 {
		fmt.Println("请输入 model 名称")
		return
	}
	filepath, _ := GetFilePath(PathArg(c.Args()), c.Args().First()+"Model.go")
	WriteFile(filepath, modelTemplate(c.Args()))
	fmt.Printf("model=%v \n", c.Args())
}

func modelTemplate(args []string) string {
	var b bytes.Buffer
	tp.ModelTmpl.Execute(&b, &InputData{
		PkgName: PkgName(args),
		Name:    StructName(args),
	})
	c := b.String()
	return c
}
