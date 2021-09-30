package core

import (
	"customPro/protoGen/model"
)

const protoFileTmp = `syntax="proto3";
{{ $msg := .ProFile }}
package {{ $msg.PackageName }};
{{ $s := $msg.ProtoService }}
service {{ $s.Name }} {
  {{- range $method := $s.Methods }} 
  rpc {{ $method.method }}
  {{- end }}
}
{{ range $m := $msg.ProtoRequest }}
message {{ $m.Name }} {
  {{- range $k, $data := $m.Fields }} 
  {{ $data.type }}  {{ $data.field }} = {{ incr $k }};
  {{- end }}
}
{{ end }}
{{ range $m := $msg.ProtoResResponse }}
message {{ $m.Name }} {
  {{- range $k, $data := $m.Fields }} 
  {{ $data.type }}  {{ $data.field }} = {{ incr $k }};
  {{- end }}
}
{{ end }}
`

//proto文件主结构体
type ProtoFile struct {
	PackageName       string
	FilePath          string
	ClassPath         string
	ProtoService      *model.Service
	ProtoRequest      []*model.Request
	ProtoResResponse  []*model.Response
}

func (p *ProtoFile) AddProtoService(protoSer *model.Service) {
	p.ProtoService = protoSer
}

func(p *ProtoFile) AddProtoMsgFromRequest(protoMsg *model.Request) {
	p.ProtoRequest = append(p.ProtoRequest, protoMsg)
}

func (p *ProtoFile) AddProtoMsgFromResponse(protoRes *model.Response) {
	p.ProtoResResponse = append(p.ProtoResResponse, protoRes)
}

