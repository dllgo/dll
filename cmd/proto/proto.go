package proto

import (
	"bytes"
	"fmt"
	"os"
	path "path/filepath"
	"strings"

	"github.com/dllgo/dll/utils"
	"github.com/urfave/cli"
)

var ProtoCMD = cli.Command{
	//命令全称
	Name: "proto",
	//命令简写
	Aliases: []string{"pb"},
	//命令详细描述
	Usage: "dll proto[pb] name pkg",
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
	output := c.Args().First()
	appPath, _, err := utils.CheckEnv(output)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	if utils.IsExist(appPath) {
		fmt.Printf("proto '%s' already exists", appPath)
	}

	fmt.Println("Creating proto...")

	os.MkdirAll(appPath, 0755)
	//
	utils.WriteToFile(path.Join(appPath, output+".proto"), protoTemplate(output, output))
	//
	utils.WriteToFile(path.Join(appPath, output+".go"), modelTemplate(output, output))
	//
	utils.WriteToFile(path.Join(appPath, "Makefile"), strings.Replace(Makefile, "{{.Appname}}", output, -1))
	//
	fmt.Println("New proto successfully created!")
}

func protoTemplate(pkg, name string) string {
	var b bytes.Buffer
	ProtoTmplCode.Execute(&b, &utils.FormatData{
		PkgName: pkg,
		Name:    utils.ToCamel(name),
	})
	c := b.String()
	return c
}
func modelTemplate(pkg, name string) string {
	var b bytes.Buffer
	ModelTmpl.Execute(&b, &utils.FormatData{
		PkgName: pkg,
		Name:    utils.ToCamel(name),
	})
	c := b.String()
	return c
}

var Makefile = `
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. *.proto

.PHONY: build
build: proto
`
