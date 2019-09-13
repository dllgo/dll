package gateway

import (
	"fmt"
	"os"
	path "path/filepath"
	"strings"

	"github.com/dllgo/dll/utils"
	"github.com/urfave/cli"
)

var GWCMD = cli.Command{
	//命令全称
	Name: "gateway",
	//命令简写
	Aliases: []string{"gw"},
	//命令详细描述
	Usage: "dll gateway[gw] name",
	//命令处理函数
	Action: func(c *cli.Context) {
		gatewayCMD(c)
	},
}

//gateway
func gatewayCMD(c *cli.Context) {
	if c.NArg() == 0 {
		fmt.Println("请输入 gateway 名称")
		return
	}
	output := c.Args().First()
	appPath, _, err := utils.CheckEnv(output)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	if utils.IsExist(appPath) {
		fmt.Printf("gateway '%s' already exists", appPath)
	}

	fmt.Println("Creating gateway...")

	os.MkdirAll(appPath, 0755)
	//
	utils.WriteToFile(path.Join(appPath, output+"_gw.go"), strings.Replace(gateway, "{{.Appname}}", output, -1))
	//
	fmt.Println("New gateway successfully created!")
}

var gateway = `
package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/dllgo/dllkit/gins"
)

func main() {
	mconf := gins.Config{Address: ":8600", ReadTimeout: 30, WriteTimeout: 30}

	httpserver, err := gins.NewServerHttp(mconf)
	if err != nil {
		fmt.Println(err)
		return
	}
	httpserver.Router = InitRouter()
	err = httpserver.Listen()
	defer httpserver.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "找不到该路由",
		})
		return
	})

	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "找不到该方法",
		})
		return
	})
	//注册api
	registerApiRouter(router)

	return router
}
func registerApiRouter(router *gin.Engine) {
	// user.NewUserHandler().Router(router)
}
`
