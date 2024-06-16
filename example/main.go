package main

import (
	"fmt"
	"github.com/cemayan/errorg"
)

func test() error {
	return errorg.New("http", "401")
}

func init() {
	errorg.Init(errorg.JsonLog())
}

func main() {
	fmt.Println(test())
}
