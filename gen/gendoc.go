package gen

import (
	_ "embed" // for including embedded resources
)

var (
	//go:embed resources/proto.tmpl
	proto []byte
)
