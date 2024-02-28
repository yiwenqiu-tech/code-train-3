package main

import (
	"code-train-3/selfcontainer"
	"fmt"
)

func ExampleListQueue() {
	listQueue := new(selfcontainer.ListQueue)
	listQueue.Add(10)
	listQueue.Add(9)
	listQueue.Add(8)
	fmt.Println("size:", listQueue.Len())      // 3
	fmt.Println("remove:", listQueue.Remove()) // 10
	fmt.Println("remove:", listQueue.Remove()) // 9
	fmt.Println("size:", listQueue.Len())      // 1
	listQueue.Add(7)
	fmt.Println("pop:", listQueue.Remove()) // 8
	fmt.Println("pop:", listQueue.Remove()) // 7
	fmt.Println("size:", listQueue.Len())   // 0
}
