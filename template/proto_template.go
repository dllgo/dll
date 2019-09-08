package template

import "html/template"

var ProtoTmpl = template.Must(template.New("proto").Parse(protoTmplCode))

const protoTmplCode = `//指定版本
syntax = "proto3";

//包名，通过protoc生成时go文件时
package {{.PkgName}};

message {{.Name}} {
    int32 id = 1;
}
`
