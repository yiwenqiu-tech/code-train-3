package main

import (
	"code-train-3/selfcontainer"
	"container/heap"
	"fmt"
)

func ExamplePriorityQueue() {
	var pq selfcontainer.PriorityQueue
	pq.Push(&selfcontainer.Item{
		Value:    "aaa",
		Priority: 3,
	})
	pq.Push(&selfcontainer.Item{
		Value:    "bbb",
		Priority: 2,
	})
	pq.Push(&selfcontainer.Item{
		Value:    "ccc",
		Priority: 10,
	})
	heap.Init(&pq)

	value := heap.Pop(&pq)
	fmt.Println(value) // value: ccc, priority: 10

	//pq.Push(&selfcontainer.Item{ // 错误，由于需要在堆里添加，这里需要用heap.Push
	//	ElemMap:    "ddd",
	//	Priority: 120,
	//})
	heap.Push(&pq, &selfcontainer.Item{
		Value:    "ddd",
		Priority: 120,
	})
	eValue := &selfcontainer.Item{
		Value:    "eee",
		Priority: 20,
	}
	heap.Push(&pq, eValue)
	value = heap.Pop(&pq)
	fmt.Println(value) // value: ddd, priority: 120

	pq.Update(eValue, "fff", 200)
	value = heap.Pop(&pq)
	fmt.Println(value) // value: fff, priority: 200
}
