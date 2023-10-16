package parse

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/emicklei/proto"
	"github.com/xiazemin/proto2docSeprateByFunc/model"
)

type messageLister struct {
	proto.NoopVisitor
	*SandBox
	currentMessage *model.Message
}

// VisitProto(p *Proto)
func (l *messageLister) VisitMessage(m *proto.Message) {

}
func (l *messageLister) VisitService(v *proto.Service) {

}
func (l *messageLister) VisitSyntax(s *proto.Syntax) {

}
func (l *messageLister) VisitPackage(p *proto.Package) {

}
func (l *messageLister) VisitOption(o *proto.Option) {
	fmt.Println(o.Name)

}
func (l *messageLister) VisitImport(i *proto.Import) {

}
func (l *messageLister) VisitNormalField(i *proto.NormalField) {
	//fmt.Println(i.Name, l.depth)

	fieldType := i.Type
	if i.Repeated {
		fieldType = "[]" + fieldType
	}

	comment := " "
	if i != nil && i.Comment != nil {
		comment = comment + strings.Join(i.Comment.Lines, " ")
	}

	// if msgType := l.meaasgeTable[i.Type]; msgType != nil {
	// 	//msgType.Accept(l)
	// 	for _, each := range msgType.Elements {
	// 		each.Accept(l)
	// 	}
	// }
	field := &model.Field{
		Type:   i.Type,
		Name:   i.Name,
		Number: strconv.FormatInt(int64(i.Sequence), 10),
	}
	l.currentMessage.Fields = append(l.currentMessage.Fields, field)
	if i.Comment != nil {
		field.Comment = i.Comment.Message()
	}

	data, _ := ioutil.ReadFile(l.fileName)
	lines := strings.Split(string(data), "\n")
	line := lines[i.Position.Line-1]
	fmt.Println("libe:", line)
	field.Options = line
	// fmt.Println("=====>", i.Sequence, i, l.currentMessage.Fields,
	// 	i.Type,
	// 	i.Name,
	// 	i.Position.String())
	//i.Comment.Accept(l)
}
func (l *messageLister) VisitEnumField(i *proto.EnumField) {

}
func (l *messageLister) VisitEnum(e *proto.Enum) {

}
func (l *messageLister) VisitComment(e *proto.Comment) {
	if e != nil {
		//fmt.Println(e.Lines)
	}
}
func (l *messageLister) VisitOneof(o *proto.Oneof) {
	//fmt.Println("oneof:", o.Name)
	for _, e := range o.Elements {
		e.Accept(l)
	}
	l.currentMessage.Oneof.Name = o.Name
}

func (l *messageLister) VisitOneofField(o *proto.OneOfField) {
	//fmt.Println("oneofField:", o.Field.Name)
	comment := " "
	if o != nil && o.Comment != nil {
		comment = comment + strings.Join(o.Comment.Lines, " ")
	}

	if msgType := l.meaasgeTable[o.Field.Type]; msgType != nil {
		//msgType.Accept(l)
		for _, each := range msgType.Elements {
			each.Accept(l)
		}
	}
	field := &model.Field{
		Type:   o.Field.Type,
		Name:   o.Field.Name,
		Number: strconv.FormatInt(int64(o.Field.Sequence), 10),
	}
	if l.currentMessage.Oneof == nil {
		l.currentMessage.Oneof = &model.Oneof{}
	}
	l.currentMessage.Oneof.Fields = append(l.currentMessage.Oneof.Fields, field)
	if o.Field.Comment != nil {
		field.Comment = o.Field.Comment.Message()
	}
}

func (l *messageLister) VisitReserved(r *proto.Reserved) {

}
func (l *messageLister) VisitRPC(r *proto.RPC) {

}
func (l *messageLister) VisitMapField(f *proto.MapField) {

}

// proto2
func (l *messageLister) VisitGroup(g *proto.Group) {

}
func (l *messageLister) VisitExtensions(e *proto.Extensions) {

}
