package greedy

import (
	"fmt"
	"math"
	"sort"
)

var minForMinimumEffort = math.MaxInt32

func minimumEffort(tasks [][]int) int {
	minForMinimumEffort = math.MaxInt32
	dfsForMinimumEffort(tasks, []int{}, 0)
	return minForMinimumEffort
}

func dfsForMinimumEffort(tasks [][]int, chosen []int, cost int) {
	if len(chosen) == len(tasks) {
		fmt.Printf("chosen: %v, cost: %v\n", chosen, cost)
		if cost < minForMinimumEffort {
			minForMinimumEffort = cost
		}
		return
	}
	tmpCost := cost
	for index, t := range tasks {
		if inChosen(chosen, index) {
			continue
		}
		// 修改值
		cost = max(cost+t[0], t[1])
		chosen = append(chosen, index)
		// 递归
		dfsForMinimumEffort(tasks, chosen, cost)
		// 还原值
		chosen = chosen[:len(chosen)-1]
		cost = tmpCost
	}
}

func inChosen(chosen []int, index int) bool {
	for _, c := range chosen {
		if c == index {
			return true
		}
	}
	return false
}

// 贪心做法
func minimumEffort2(tasks [][]int) (ans int) {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i][0]-tasks[i][1] < tasks[j][0]-tasks[j][1]
	})
	for i := len(tasks) - 1; i >= 0; i-- {
		ans = max(ans+tasks[i][0], tasks[i][1])
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
