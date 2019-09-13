package micro

import "html/template"

var HandlerTmpl = template.Must(template.New("handler").Parse(handlerTmplCode))

const handlerTmplCode = `
package handler
import (
	"context"

	"github.com/dllgo/dlldb"
	client "github.com/micro/go-micro/client"
)

type {{.Name}}Handler struct {
	Client client.Client,
	iSrv dlldb.IServices
}
func New{{.Name}}Handler(cli client.Client) *{{.Name}}Handler {
	return &{{.Name}}Handler{
		Client:cli,
		iSrv:dlldb.NewServices(),
	}
}
func (this *{{.Name}}Handler) List(ctx context.Context) error {
	return nil
}
`
