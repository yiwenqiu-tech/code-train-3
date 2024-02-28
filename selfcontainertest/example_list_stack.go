package main

import (
	"code-train-3/selfcontainer"
	"fmt"
)

func ExampleListStack() {
	listStack := new(selfcontainer.ListStack)
	listStack.Push(10)
	listStack.Push(9)
	listStack.Push(8)
	fmt.Println("size:", listStack.Len())
	fmt.Println("pop:", listStack.Pop())
	fmt.Println("pop:", listStack.Pop())
	fmt.Println("size:", listStack.Len())
	listStack.Push(7)
	fmt.Println("pop:", listStack.Pop())
}
