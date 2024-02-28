package main

import (
	"code-train-3/selfcontainer"
	"fmt"
)

func ExampleArrayQueue() {
	arrayQueue := new(selfcontainer.ArrayQueue)
	arrayQueue.Add(10)
	arrayQueue.Add(9)
	arrayQueue.Add(8)
	fmt.Println("size:", arrayQueue.Len())      // 3
	fmt.Println("remove:", arrayQueue.Remove()) // 10
	fmt.Println("remove:", arrayQueue.Remove()) // 9
	fmt.Println("size:", arrayQueue.Len())      // 1
	arrayQueue.Add(7)
	fmt.Println("pop:", arrayQueue.Remove()) // 8
	fmt.Println("pop:", arrayQueue.Remove()) // 7
	fmt.Println("size:", arrayQueue.Len())   // 0
}
