package parse

import (
	"fmt"
	"os"

	"github.com/emicklei/proto"
	"github.com/xiazemin/proto2docSeprateByFunc/model"
)

type SandBox struct {
	*model.Template
	meaasgeTable map[string]*proto.Message
}

func NewSandBox() *SandBox {
	return &SandBox{
		Template:     &model.Template{},
		meaasgeTable: make(map[string]*proto.Message),
	}
}

func (s *SandBox) ProtoSplit(src string) *model.Template {
	reader, err := os.Open(src)

	if err != nil {
		fmt.Println(err)
	}
	defer reader.Close()

	parser := proto.NewParser(reader)
	definition, _ := parser.Parse()

	proto.Walk(definition,
		proto.WithPackage(s.handlePackage),
		proto.WithOption(s.handleOption),
	)

	// proto.Walk(definition,
	// 	proto.WithRPC(s.handleRpc))
	// proto.Walk(definition,
	// 	proto.WithOption(s.handleOption))
	//解决嵌套问题
	proto.Walk(definition,
		proto.WithMessage(s.preHandleMessage),
	)
	//解决声明顺序的问题
	proto.Walk(definition,
		proto.WithMessage(s.handleMessage),
	)
	proto.Walk(definition,
		proto.WithService(s.handleService),
	)
	return s.Template
}

func (s *SandBox) handlePackage(p *proto.Package) {
	s.Package = p.Name //package
	//s.Accept(root)
}

func (s *SandBox) handleOption(opt *proto.Option) {
	if opt == nil {
		return
	}
	// for _, e := range opt.Elements {
	// 	e.Accept(l)
	// 	//fmt.Println(i)
	// }
	//fmt.Println(opt, opt.Name, opt.Constant.SourceRepresentation())
	s.Options = append(s.Options, &model.Option{
		Name:  opt.Name,
		Value: opt.Constant.SourceRepresentation(),
	})
}

func (s *SandBox) handleService(srv *proto.Service) {
	// s.serviceName = srv.Name //service
	root := &serviceLister{
		SandBox:        s,
		currentService: &model.Service{Name: srv.Name},
	}
	if srv.Comment != nil {
		root.currentService.Comment = srv.Comment.Message()
	}
	s.Services = append(s.Services, root.currentService)
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
	if message := s.meaasgeTable[m.Name]; message != nil {
		// num++
		// lister.parentId = strconv.FormatInt(num, 10)
		// s.sheet = append(s.sheet, &xmlXmind.XmindNode{
		// 	NodeID:       strconv.FormatInt(num, 10),
		// 	TopicContent: m.Name, //response
		// 	ParentID:     refer,  //函数参数或者返回值
		// })
		lister := &messageLister{
			SandBox: s,
			currentMessage: &model.Message{
				Name: message.Name,
			},
		}
		if message.Comment != nil {
			lister.currentMessage.Comment = message.Comment.Message()
		}
		s.Messages = append(s.Messages, lister.currentMessage)
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
	// s.funcTable[s.packageName+":"+s.serviceName+":"+rpc.Name] = rpc
	for _, e := range rpc.Elements {
		e.Accept(l)
		//fmt.Println(i)
	}
}
