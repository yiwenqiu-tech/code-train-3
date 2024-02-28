package main

import (
	"code-train-3/selfcontainer"
	"fmt"
)

func ExampleArrayStack() {
	arrayStack := new(selfcontainer.ArrayStack)
	arrayStack.Push(10)
	arrayStack.Push(9)
	arrayStack.Push(8)
	fmt.Println("size:", arrayStack.Len())
	fmt.Println("pop:", arrayStack.Pop())
	fmt.Println("pop:", arrayStack.Pop())
	fmt.Println("size:", arrayStack.Len())
	arrayStack.Push(7)
	fmt.Println("pop:", arrayStack.Pop())
}
