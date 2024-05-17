package bit_operation

func myPow(x float64, n int) float64 {
	if n < 0 {
		return myPow(x, -n)
	}
	var temp = x
	var ans float64 = 1
	for n > 0 {
		if n&1 == 1 {
			ans *= temp
		}
		temp *= temp
		n = n >> 1
	}
	return ans
}
