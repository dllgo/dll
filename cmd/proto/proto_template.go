package proto

import "html/template"

var ProtoTmplCode = template.Must(template.New("proto").Parse(protoTmplCode))

const protoTmplCode = `
syntax = "proto3";//指定版本

option go_package = "{{.PkgName}}";
package {{.PkgName}};//包名，通过protoc生成时go文件时

service {{.Name}}Service {
	rpc Test({{.Name}}Request) returns ({{.Name}}Response) {}
}

message {{.Name}}Request {
	string name = 1;
}

message {{.Name}}Response {
	string msg = 1;
}

message {{.Name}} {
    int32 id = 1;
}
`
