package template

import "html/template"

var HandlerTmpl = template.Must(template.New("handler").Parse(handlerTmplCode))

const handlerTmplCode = `
package {{.PkgName}}

type {{.Name}}Handler struct {

}
func New{{.Name}}Handler() *{{.Name}}Handler {
	return &{{.Name}}Handler{}
}
`
