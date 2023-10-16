package model

type Template struct {
	Package  string
	Imports  []string
	Options  []*Option
	Services []*Service
	Messages []*Message
}

type Option struct {
	Name  string
	Value string
}

type Service struct {
	Name    string
	Comment []string
	Rpcs    []*Rpc
}

type Oneof struct {
	Name   string
	Fields []*Field
}

type Rpc struct {
	Name     string
	Comment  []string
	Request  string
	Response string
	Options  []*Option
}

type Message struct {
	Name    string
	Comment []string
	Fields  []*Field
	Oneof   *Oneof
}

type Field struct {
	Type    string
	Name    string
	Number  string
	Comment []string
	Options string
}
