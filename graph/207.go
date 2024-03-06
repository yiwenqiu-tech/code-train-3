package graph

import "container/list"

func canFinish(numCourses int, prerequisites [][]int) bool {
	to := make([][]int, numCourses)
	in := make([]int, numCourses)
	for _, p := range prerequisites {
		a := p[0]
		b := p[1]
		to[b] = append(to[b], a) // 构造出边数组 题目：[0, 1]表示完成0之前，必须完成1。所以出边数组是，1->0
		in[a]++                  // a的入度+1
	}
	queue := list.List{}
	for index, i := range in {
		if i == 0 { // 入度为0，则可添加到队列里
			queue.PushBack(index)
		}
	}

	var lesson []int
	for queue.Len() > 0 {
		head := queue.Front()
		queue.Remove(head)
		headVal := head.Value.(int)
		lesson = append(lesson, headVal)

		toData := to[headVal]
		for _, i := range toData {
			in[i]--
			if in[i] == 0 { // 当入度为0时，则可添加到队列里
				queue.PushBack(i)
			}
		}
	}

	return len(lesson) == numCourses
}
