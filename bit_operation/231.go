package bit_operation

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	if n-(n&(-n)) == 0 { // lowbit
		return true
	} else {
		return false
	}
}
