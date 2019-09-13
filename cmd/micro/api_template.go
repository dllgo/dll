package micro

import "html/template"

var ApiTmpl = template.Must(template.New("api").Parse(apiTmplCode))

const apiTmplCode = `
package api
import (
	"github.com/gin-gonic/gin"
)

type {{.Name}}Api struct {
}
func New{{.Name}}Api() *{{.Name}}Api {
	return &{{.Name}}Api{}
}
func (this *{{.Name}}Api) Router(router *gin.Engine) {
	r := router.Group("{{.Name}}")
	r.POST("test", this.test)
	r.GET("test", this.test)
}
func (this *{{.Name}}Api) test(ctx *gin.Context) {
	return ctx.JSON(http.StatusOK, {"code":200,"msg":"成功"})
}
`
