package parse

import (
	"fmt"

	"github.com/emicklei/proto"
	"github.com/xiazemin/proto2docSeprateByFunc/model"
)

type serviceLister struct {
	proto.NoopVisitor
	*SandBox
	currentService *model.Service
}

func (l *serviceLister) VisitService(v *proto.Service) {
	//fmt.Println(v.Name)
	l.currentService.Name = v.Name
	l.Services = append(l.Services, l.currentService)
	// fmt.Println("VisitService:", l, v)
	for i, e := range v.Elements {
		e.Accept(l)
		fmt.Println(i)
	}
	v.Accept(l)
}
func (l *serviceLister) VisitNormalField(i *proto.NormalField) {
	//fmt.Println(i.Name)
}
func (l *serviceLister) VisitRPC(r *proto.RPC) {
	rpc := &model.Rpc{
		Name:     r.Name,
		Request:  r.RequestType,
		Response: r.ReturnsType,
	}
	l.currentService.Rpcs = append(l.currentService.Rpcs, rpc)
	if r.Comment != nil {
		rpc.Comment = r.Comment.Message()
	}
	// fmt.Println(l, r)
	//fmt.Println(r.Name, r.RequestType, r.ReturnsType)
	// l.num++
	// l.sheet = append(l.sheet, &xmlXmind.XmindNode{
	// 	NodeID:       strconv.FormatInt(l.num, 10),
	// 	TopicContent: r.Name,     //func (s*SandBox)tion
	// 	ParentID:     l.parentId, //service
	// })

	// l.sheet = append(l.sheet, &xmlXmind.XmindNode{
	// 	NodeID:       strconv.FormatInt(l.num+1, 10),
	// 	TopicContent: r.RequestType,                //request
	// 	ParentID:     strconv.FormatInt(l.num, 10), //func (s*SandBox)tion
	// })
	// l.messageTypeToParentId[r.RequestType] = append(l.messageTypeToParentId[r.RequestType], strconv.FormatInt(l.num+1, 10))
	// l.sheet = append(l.sheet, &xmlXmind.XmindNode{
	// 	NodeID:       strconv.FormatInt(l.num+2, 10),
	// 	TopicContent: r.ReturnsType,                //response
	// 	ParentID:     strconv.FormatInt(l.num, 10), //func (s*SandBox)tion
	// })
	// l.messageTypeToParentId[r.ReturnsType] = append(l.messageTypeToParentId[r.ReturnsType], strconv.FormatInt(l.num+2, 10))
	// l.num += 2
}
