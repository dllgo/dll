package micro

import (
	"bytes"
	"fmt"
	"os"
	path "path/filepath"
	"strings"

	"github.com/dllgo/dll/utils"
	"github.com/urfave/cli"
)

var MicroCMD = cli.Command{
	//命令全称
	Name: "micro",
	//命令简写
	Aliases: []string{"micro"},
	//命令详细描述
	Usage: "dll micro  appname",
	//命令处理函数
	Action: func(c *cli.Context) {
		microCMD(c)
	},
}

func microCMD(c *cli.Context) {
	if c.NArg() == 0 {
		fmt.Println("请输入 micro 名称")
		return
	}
	output := c.Args().First()
	appPath, packPath, err := utils.CheckEnv(output)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	if utils.IsExist(appPath) {
		fmt.Printf("micro '%s' already exists", appPath)
		return
	}

	fmt.Println("Creating micro...")

	os.MkdirAll(appPath, 0755)
	//
	utils.WriteToFile(path.Join(appPath, "main.go"), strings.Replace(Maingo, "{{.Appname}}", packPath, -1))
	utils.WriteToFile(path.Join(appPath, "Makefile"), strings.Replace(Makefile, "{{.Appname}}", output, -1))
	utils.WriteToFile(path.Join(appPath, "Dockerfile"), strings.Replace(Dockerfile, "{{.Appname}}", output, -1))
	//
	os.Mkdir(path.Join(appPath, "conf"), 0755)
	utils.WriteToFile(path.Join(appPath, "conf", "config.json"), strings.Replace(Config, "{{.Appname}}", output, -1))
	//
	os.Mkdir(path.Join(appPath, "cmd"), 0755)
	utils.WriteToFile(path.Join(appPath, "cmd", "server.go"), strings.Replace(Servergo, "{{.Appname}}", packPath, -1))
	//
	os.Mkdir(path.Join(appPath, "handler"), 0755)
	utils.WriteToFile(path.Join(appPath, "handler", output+".go"), handlerTemplate(packPath, output))
	//
	os.Mkdir(path.Join(appPath, "api"), 0755)
	utils.WriteToFile(path.Join(appPath, "api", output+".go"), apiTemplate(packPath, output))

	fmt.Println("New micro successfully created!")

}
func handlerTemplate(pkg, name string) string {
	var b bytes.Buffer
	HandlerTmpl.Execute(&b, &utils.FormatData{
		PkgName: pkg,
		Name:    utils.ToCamel(name),
	})
	c := b.String()
	return c
}
func apiTemplate(pkg, name string) string {
	var b bytes.Buffer
	ApiTmpl.Execute(&b, &utils.FormatData{
		PkgName: pkg,
		Name:    utils.ToCamel(name),
	})
	c := b.String()
	return c
}
