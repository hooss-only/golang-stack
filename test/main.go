package main

import (
	"fmt"

	"github.com/hooss-only/stack"
)

func main() {
	s := stack.Stack{}
	s.Push("Hello")
	fmt.Println(s.GetStack())
}
