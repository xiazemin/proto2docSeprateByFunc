package gen

import (
	"bytes"
	text_template "text/template"

	"github.com/xiazemin/proto2docSeprateByFunc/model"
)

func NewProtoRender() *protoRenderer {
	return &protoRenderer{string(proto)}
}

type protoRenderer struct {
	inputTemplate string
}

func (mr *protoRenderer) Apply(template *model.Template) ([]byte, error) {
	tmpl, err := text_template.New("Text Template").Parse(mr.inputTemplate)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err = tmpl.Execute(&buf, template); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
