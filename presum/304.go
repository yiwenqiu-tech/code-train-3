package presum

type NumMatrix struct {
	PreSum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	numMatrix := NumMatrix{}
	var preSum = make([][]int, len(matrix)+1)
	// init
	for i := 0; i < len(matrix)+1; i++ {
		preSum[i] = make([]int, len(matrix[0])+1)
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			preSum[i+1][j+1] = preSum[i+1][j] + preSum[i][j+1] - preSum[i][j] + matrix[i][j]
		}
	}

	numMatrix.PreSum = preSum
	return numMatrix
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.PreSum[row2+1][col2+1] - this.PreSum[row2+1][col1] - this.PreSum[row1][col2+1] + this.PreSum[row1][col1]
}
