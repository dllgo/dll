package template

import "html/template"

var ModelTmpl = template.Must(template.New("model").Parse(modelTmplCode))

const modelTmplCode = `
package {{.PkgName}}
import (
	"github.com/dllgo/dlldb"
)

type {{.Name}}Model struct {
	dlldb.IModel
}
`
