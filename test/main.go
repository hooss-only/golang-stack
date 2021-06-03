package main

import (
	"fmt"

	"github.com/hooss-only/stack"
)

func main() {
	s := stack.Stack{}
	s.Push("Hello")
	s.Push("World")
	s.Pop()
	fmt.Println(s.GetStack())
	fmt.Println("stack lenght:", len(s.GetStack()))
	s.Init()
	fmt.Println(s.GetStack())
}
