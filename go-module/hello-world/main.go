package main

import (
	"fmt"
	h1 "github.com/corn404/go-hello-world"
	h2 "github.com/corn404/go-hello-world/v2"
)

func main() {
	fmt.Println(h1.HelloWold())
	fmt.Println(h2.HelloWold("Code Corn"))
}
