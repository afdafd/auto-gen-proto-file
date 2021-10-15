package core

import (
	"bytes"
	"customPro/protoGen/model"
	"fmt"
	"os"
	"text/template"
)

const protoFile = `syntax="proto3";
{{ $msg := .ProFile }}
package {{ $msg.BaseSet.PackageName }};

{{ range $s := $msg.ProtoService }}
service {{ $s.SerName }} {
  {{- range $md := $s.SerMethods }}
  rpc {{ $md.MethodName -}}({{ $md.RequestName }}) returns ({{ $md.ResponseName }}) {}
  {{- end }}
}
{{ end }}

{{ range $m := $msg.ProtoRequest }}
message {{ $m.ReqName }} {
  {{- range $k, $data := $m.Fields }} 
  {{ $data.type }}  {{ $data.field }} = {{ incr $k }};
  {{- end }}
}
{{ end }}

{{ range $m := $msg.ProtoResResponse }}
message {{ $m.ResName }} {
  {{- range $k, $data := $m.Fields }} 
  {{ $data.type }}  {{ $data.field }} = {{ incr $k }};
  {{- end }}
}
{{ end }}
`

//proto文件主结构体
type ProtoFile struct {
	BaseSet           *model.BaseSets
	ProtoService      []*model.Service
	ProtoRequest      []*model.Request
	ProtoResResponse  []*model.Response
}

//开始生成proto-file文件
func (p *ProtoFile) GenProtoFile()  {
	data := struct {
		ProFile *ProtoFile
	}{
		ProFile: p,
	}

	t := template.Must(template.New("protoFile").Funcs(template.FuncMap{
		"incr": func(key int) int {
			return key + 1
		},
	}).Parse(protoFile))

	//缓冲
	out := bytes.NewBuffer(nil)

	//执行
	if err := t.Execute(out, data); err != nil {
		panic(err)
	}

	//路径处理
	filePath := p.BaseSet.ClassName + ".proto"
	if _, er := os.Stat(filePath); er == nil {
		if errr := os.Remove(filePath); errr != nil {
			panic(errr)
		}
	}

	//输出
	files, _ := os.Create(filePath)
	in, e := files.Write(out.Bytes())

	if e != nil {
		fmt.Println("写入数据错误")
		panic(e)
	}

	fmt.Println("创建成功", in)
}


func (p *ProtoFile) AddProtoService(protoSer *model.Service) {
	p.ProtoService = append(p.ProtoService, protoSer)
}

func(p *ProtoFile) AddProtoMsgFromRequest(protoMsg *model.Request) {
	p.ProtoRequest = append(p.ProtoRequest, protoMsg)
}

func (p *ProtoFile) AddProtoMsgFromResponse(protoRes *model.Response) {
	p.ProtoResResponse = append(p.ProtoResResponse, protoRes)
}

