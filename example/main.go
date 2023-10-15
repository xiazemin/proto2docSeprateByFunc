package main

import (
	gen "github.com/xiazemin/proto2docSeprateByFunc/gen"
)

func main() {
	gen.GenFiles("./doc/", "./example/sub.proto", "./example/request.proto")
}
