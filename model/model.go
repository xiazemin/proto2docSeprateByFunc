package model

type Template struct {
	Package  string
	Options  string
	Services []*Service
	Messages []*Message
}

type Service struct {
	Name string
	Rpcs []*Rpc
}

type Rpc struct {
	Name     string
	Request  string
	Response string
}

type Message struct {
	Name   string
	Fields []*Field
}

type Field struct {
	Type   string
	Name   string
	Number string
}
