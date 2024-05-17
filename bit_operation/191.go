package bit_operation

func hammingWeight(n int) int {
	var ans int
	for n != 0 {
		if n&1 == 1 {
			ans++
		}
		n = n >> 1
	}
	return ans
}
