package binary_search

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	if len(matrix[0]) == 0 {
		return false
	}

	m := len(matrix)
	n := len(matrix[0])
	total := m * n
	left := 0
	right := total - 1
	for left <= right {
		mid := left + (right-left)/2
		row := mid / n
		col := mid % n
		if matrix[row][col] == target {
			return true
		}
		if matrix[row][col] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
