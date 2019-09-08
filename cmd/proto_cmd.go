package cmd

import (
	"bytes"
	"fmt"

	tp "github.com/dllgo/dll/template"
	"github.com/urfave/cli"
)

var ProtoCMD = cli.Command{
	//命令全称
	Name: "proto",
	//命令简写
	Aliases: []string{"pb"},
	//命令详细描述
	Usage: "dll proto[pb] name pkgname path",
	//命令处理函数
	Action: func(c *cli.Context) {
		protoCMD(c)
	},
}

//proto
func protoCMD(c *cli.Context) {
	if c.NArg() == 0 {
		fmt.Println("请输入 proto 名称")
		return
	}
	filepath, _ := GetFilePath(PathArg(c.Args()), c.Args().First()+".proto")
	WriteFile(filepath, protoTemplate(c.Args()))
	fmt.Printf("proto=%v \n", c.Args())
}

func protoTemplate(args []string) string {
	var b bytes.Buffer
	tp.ProtoTmpl.Execute(&b, &InputData{
		PkgName: PkgName(args),
		Name:    StructName(args),
	})
	c := b.String()
	return c
}
