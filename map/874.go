package _map

import "strconv"

// -2 ：向左转 90 度
// -1 ：向右转 90 度
// 1 <= x <= 9 ：向前移动 x 个单位长度
func robotSim(commands []int, obstacles [][]int) int {
	var directions = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	var curDir = 0
	// 障碍物map
	var obsMap = make(map[string]struct{})
	for _, obs := range obstacles {
		obsMap[strconv.Itoa(obs[0])+"_"+strconv.Itoa(obs[1])] = struct{}{}
	}
	x := 0
	y := 0
	maxLen := 0
	for _, c := range commands {
		if c == -1 {
			curDir = (curDir + 1) % 4
		} else if c == -2 {
			curDir = (curDir - 1 + 4) % 4
		} else {
			for i := 0; i < c; i += 1 {
				tmpX := x + directions[curDir][0]
				tmpY := y + directions[curDir][1]
				if _, exist := obsMap[strconv.Itoa(tmpX)+"_"+strconv.Itoa(tmpY)]; exist {
					break
				}
				x = tmpX
				y = tmpY
				tmpMax := x*x + y*y
				if tmpMax > maxLen {
					maxLen = tmpMax
				}
			}
		}
	}
	return maxLen
}
