package train

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}
	total := m * n
	left := 0
	right := total - 1
	for left <= right {
		mid := left + (right-left)/2
		midNum := getNumByIndex(matrix, mid, m, n)
		if midNum == target {
			return true
		} else if midNum > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

func getNumByIndex(matrix [][]int, index int, m int, n int) int {
	row := index / n
	col := index % n
	return matrix[row][col]
}
