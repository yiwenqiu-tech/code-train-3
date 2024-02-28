package divide_conquer

func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1 / recurForMyPow(x, -n)
	}
	return recurForMyPow(x, n)
}

func recurForMyPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	odd := n % 2
	if odd == 0 {
		res := recurForMyPow(x, n/2)
		return res * res
	} else {
		res := recurForMyPow(x, n/2)
		return res * res * x
	}
}
