package str

import "math"

func repeatedStringMatch(a string, b string) int {
	an := len(a)
	bn := len(b)
	max := int(math.Ceil(float64(bn)/float64(an))) + 1 // 最多b/a + 1
	aH := make([]int64, max*an+1)
	var c int64 = 131
	var p int64 = 1e9 + 7
	for i := 1; i <= max*an; i++ {
		aH[i] = (aH[i-1]*c + (int64(a[(i-1)%an])) - 'a' + 1) % p
	}
	var base int64 = 1
	var bHash int64
	for i := 0; i < bn; i++ {
		bHash = (bHash*c + int64(b[i]) - 'a' + 1) % p
		base = base * c % p
	}

	for l := 0; l <= max*an-bn; l++ {
		r := l + bn - 1
		curHash := ((aH[r+1]-aH[l]*base)%p + p) % p
		if curHash == bHash {
			return r/an + 1
		}
	}
	return -1
}
