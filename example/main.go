package main

import (
	"fmt"

	gen "github.com/xiazemin/proto2docSeprateByFunc/gen"
)

func main() {
	all := gen.GenFiles("./example/gen/doc", true, "./example/sub.proto", "./example/request.proto")
	fmt.Println(all)
}
