package main

import (
	gen "github.com/xiazemin/proto2docSeprateByFunc/gen"
)

func main() {
	gen.GenFiles("./example/gen/doc", true, "./example/sub.proto", "./example/request.proto")
}
