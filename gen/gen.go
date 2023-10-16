package gen

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	// plugin_go "github.com/golang/protobuf/protoc-gen-go/plugin"
	// gendoc "github.com/pseudomuto/protoc-gen-doc"
	"github.com/xiazemin/proto2docSeprateByFunc/model"
	parse "github.com/xiazemin/proto2docSeprateByFunc/parse"
)

func GenFiles(dst string, debug bool, files ...string) {
	if len(files) < 1 {
		fmt.Println("files empty")
		return
	}

	for i := 0; i < len(files); i++ {
		sandbox := parse.NewSandBox(files[i])
		template := sandbox.ProtoSplit(files[i])
		// for name, m := range msgs {
		// 	fmt.Println(dst+name+".proto", m)
		// }
		r := NewProtoRender()
		b, e := r.Apply(template)
		if debug {
			fmt.Println("position:"+files[i]+".gen",
				string(b),
				e,
				marshal(template),
				ioutil.WriteFile(files[i]+".gen", b, 0777))
		}
		messages := make(map[string]*model.Message)
		for _, msg := range template.Messages {
			messages[msg.Name] = msg
		}

		for _, srv := range template.Services {
			for _, rpc := range srv.Rpcs {
				splitTemplate := &model.Template{
					Package: template.Package,
					Imports: template.Imports,
					Options: template.Options,
					Services: []*model.Service{{
						Name:    srv.Name,
						Comment: srv.Comment,
						Rpcs: []*model.Rpc{{
							Name:     rpc.Name,
							Comment:  rpc.Comment,
							Request:  rpc.Request,
							Response: rpc.Response,
							Options:  rpc.Options,
						}},
					}},
					Messages: append(findRelativeMessages(messages, rpc.Request), findRelativeMessages(messages, rpc.Response)...),
				}
				r := NewProtoRender()
				b, e := r.Apply(splitTemplate)
				if e != nil {
					fmt.Println(e)
				}
				ioutil.WriteFile(dst+strconv.FormatInt(int64(i), 10)+"."+srv.Name+"."+rpc.Name+".proto", b, 0777)
			}
		}
	}

	fmt.Println("========================")
	t := &model.Template{
		Package: "hello",
		Options: []*model.Option{{Name: "go_package", Value: "\"hello/go\""}},
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
	if debug {
		ioutil.WriteFile("./example/gen.proto", b, 0777)
	}

	// if err := protokit.RunPlugin(new(gendoc.Plugin)); err != nil {
	// 	log.Fatal(err)
	// }

	// p := new(gendoc.Plugin)
	// req := &plugin_go.CodeGeneratorRequest{}
	// resp, err := p.Generate(req)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(resp)

}

func marshal(v interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(data)
}

func findRelativeMessages(all map[string]*model.Message, name string) []*model.Message {
	if all[name] == nil {
		return nil
	}
	var r []*model.Message
	r = append(r, all[name])
	for _, field := range all[name].Fields {
		fm := findRelativeMessages(all, field.Type)
		r = append(r, fm...)
	}
	return r
}
