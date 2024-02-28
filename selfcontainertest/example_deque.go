package main

import (
	"code-train-3/selfcontainer"
	"fmt"
)

func ExampleDeque() {
	deque := new(selfcontainer.Deque)
	deque.PushBack("aaa")
	deque.PushBack("bbb")
	deque.PushBack("ccc")
	deque.PushBack("ddd")
	deque.PushBack("eee") // 队尾插入

	fmt.Println("size:", deque.Len()) // 5
	head := deque.Front()             // 队头访问
	fmt.Println(head.Value)           // "aaa"
	deque.Remove(head)                // 队头删除
	head = deque.Front()
	fmt.Println(head.Value) // "bbb"
	deque.Remove(head)
	fmt.Println(deque.Back().Value) // "eee" 队尾访问

	deque.PushFront("fff")           // 队头插入
	fmt.Println(deque.Front().Value) // "fff" 队头访问

}
