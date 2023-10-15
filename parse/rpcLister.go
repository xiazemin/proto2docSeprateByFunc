package parse

import (
	"github.com/emicklei/proto"
)

type rpcLister struct {
	proto.NoopVisitor
	*SandBox
}

func (l rpcLister) VisitService(v *proto.Service) {
	//fmt.Println(v.Name)
	// for i, e := range v.Elements {
	// 	e.Accept(root)
	// 	fmt.Println(i)
	// }
	// v.Accept(root)
}
func (l rpcLister) VisitNormalField(i *proto.NormalField) {
	//fmt.Println(i.Name)
}
func (l rpcLister) VisitRPC(r *proto.RPC) {
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
