package parse

import (
	"fmt"
	"strings"

	"github.com/emicklei/proto"
)

type messageLister struct {
	proto.NoopVisitor
	*SandBox
}

//VisitProto(p *Proto)
func (l messageLister) VisitMessage(m *proto.Message) {

}
func (l messageLister) VisitService(v *proto.Service) {

}
func (l messageLister) VisitSyntax(s *proto.Syntax) {

}
func (l messageLister) VisitPackage(p *proto.Package) {

}
func (l messageLister) VisitOption(o *proto.Option) {
	fmt.Println(o.Name)

}
func (l messageLister) VisitImport(i *proto.Import) {

}
func (l messageLister) VisitNormalField(i *proto.NormalField) {
	//fmt.Println(i.Name, l.depth)

	fieldType := i.Type
	if i.Repeated {
		fieldType = "[]" + fieldType
	}

	comment := " "
	if i != nil && i.Comment != nil {
		comment = comment + strings.Join(i.Comment.Lines, " ")
	}

	if msgType := l.meaasgeTable[i.Type]; msgType != nil {
		//msgType.Accept(l)
		for _, each := range msgType.Elements {
			each.Accept(l)
		}
	}
	//i.Comment.Accept(l)
}
func (l messageLister) VisitEnumField(i *proto.EnumField) {

}
func (l messageLister) VisitEnum(e *proto.Enum) {

}
func (l messageLister) VisitComment(e *proto.Comment) {
	if e != nil {
		//fmt.Println(e.Lines)
	}
}
func (l messageLister) VisitOneof(o *proto.Oneof) {
	//fmt.Println("oneof:", o.Name)
	for _, e := range o.Elements {
		e.Accept(l)
	}
}

func (l messageLister) VisitOneofField(o *proto.OneOfField) {
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
}

func (l messageLister) VisitReserved(r *proto.Reserved) {

}
func (l messageLister) VisitRPC(r *proto.RPC) {

}
func (l messageLister) VisitMapField(f *proto.MapField) {

}

// proto2
func (l messageLister) VisitGroup(g *proto.Group) {

}
func (l messageLister) VisitExtensions(e *proto.Extensions) {

}
