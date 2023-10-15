package parse

import (
	"fmt"
	"os"

	"github.com/emicklei/proto"
)

type SandBox struct {
	packageName  string
	serviceName  string
	meaasgeTable map[string]*proto.Message
	funcTable    map[string]*proto.RPC
}

func NewSandBox() *SandBox {
	return &SandBox{
		meaasgeTable: make(map[string]*proto.Message),
		funcTable:    make(map[string]*proto.RPC),
	}
}

func (s *SandBox) ProtoSplit(src string) map[string]*proto.RPC {
	reader, err := os.Open(src)

	if err != nil {
		fmt.Println(err)
	}
	defer reader.Close()

	parser := proto.NewParser(reader)
	definition, _ := parser.Parse()

	proto.Walk(definition,
		proto.WithPackage(s.handlePackage),
		proto.WithService(s.handleService),
	)
	//解决嵌套问题
	proto.Walk(definition,
		proto.WithMessage(s.preHandleMessage),
	)
	//解决声明顺序的问题
	proto.Walk(definition,
		proto.WithMessage(s.handleMessage),
	)

	proto.Walk(definition,
		proto.WithRPC(s.handleRpc))
	return s.funcTable
}

func (s *SandBox) handlePackage(p *proto.Package) {
	s.packageName = p.Name //package
	//s.Accept(root)
}
func (s *SandBox) handleService(srv *proto.Service) {
	s.serviceName = srv.Name //service
	root := serviceLister{
		SandBox: s,
	}
	//s.Accept(root)
	for _, e := range srv.Elements {
		e.Accept(root)
		//fmt.Println(i)
	}
}

func (s *SandBox) preHandleMessage(m *proto.Message) {
	//fmt.Println("preMsg:", m.Name)
	s.meaasgeTable[m.Name] = m
}

func (s *SandBox) handleMessage(m *proto.Message) {
	lister := new(messageLister)
	lister.SandBox = s
	if message := s.meaasgeTable[m.Name]; message != nil {
		// num++
		// lister.parentId = strconv.FormatInt(num, 10)
		// s.sheet = append(s.sheet, &xmlXmind.XmindNode{
		// 	NodeID:       strconv.FormatInt(num, 10),
		// 	TopicContent: m.Name, //response
		// 	ParentID:     refer,  //函数参数或者返回值
		// })
		for _, each := range m.Elements {
			each.Accept(lister)
		}
	} else {
		fmt.Println("message not found:", m.Name)
	}
	//fmt.Println(m.Name)
}

func (s *SandBox) handleRpc(rpc *proto.RPC) {
	if rpc == nil {
		return
	}
	l := new(rpcLister)
	s.funcTable[s.packageName+":"+s.serviceName+":"+rpc.Name] = rpc
	for _, e := range rpc.Elements {
		e.Accept(l)
		//fmt.Println(i)
	}
}
