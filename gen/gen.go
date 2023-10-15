package gen

import (
	"fmt"

	gendoc "github.com/pseudomuto/protoc-gen-doc"
	"github.com/xiazemin/proto2docSeprateByFunc/model"
	parse "github.com/xiazemin/proto2docSeprateByFunc/parse"
)

func GenFiles(dst string, files ...string) {
	if len(files) < 1 {
		fmt.Println("files empty")
		return
	}

	sandbox := parse.NewSandBox()
	for i := 1; i < len(files); i++ {
		msgs := sandbox.ProtoSplit(files[i])
		for name, m := range msgs {
			fmt.Println(dst+name+".proto", m)
		}
	}

	t := &model.Template{
		Package: "hello",
		Options: "option go_package = \"hello/go\"",
		Services: []*model.Service{{Name: "hello1",
			Rpcs: []*model.Rpc{{Name: "Hi",
				Request:  "req1",
				Response: "resp1"}, {Name: "Hi2",
				Request:  "req2",
				Response: "resp2"}}}, {Name: "hello2",
			Rpcs: []*model.Rpc{{Name: "Hi",
				Request:  "req1",
				Response: "resp1"}, {Name: "Hi2",
				Request:  "req2",
				Response: "resp2"}}}},
		Messages: []*model.Message{{Name: "req1",
			Fields: []*model.Field{{Type: "string",
				Name:   "field1",
				Number: "1"}}}, {Name: "resp1",
			Fields: []*model.Field{{Type: "int64",
				Name:   "field2",
				Number: "1"}}},
			{Name: "req2",
				Fields: []*model.Field{{Type: "string",
					Name:   "field1",
					Number: "1"}}}, {Name: "resp2",
				Fields: []*model.Field{{Type: "int64",
					Name:   "field2",
					Number: "1"}}}},
	}
	r := NewProtoRender()
	b, e := r.Apply(t)
	fmt.Println(string(b), e)

	// if err := protokit.RunPlugin(new(gendoc.Plugin)); err != nil {
	// 	log.Fatal(err)
	// }

	p := new(gendoc.Plugin)
	req := &protokit.CodeGeneratorRequest{}
	resp, err := p.Generate(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}
