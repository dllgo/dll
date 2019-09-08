package template

import "html/template"

var ServiceTmpl = template.Must(template.New("service").Parse(serviceTmplCode))

const serviceTmplCode = `
package {{.PkgName}}
import (
	"github.com/dllgo/dlldb"
)

type {{.Name}}Service struct {
	iSrv dlldb.IServices
}
func New{{.Name}}Service() *{{.Name}}Service {
	return &{{.Name}}Service{
		iSrv:dlldb.NewServices(),
	}
}
`
