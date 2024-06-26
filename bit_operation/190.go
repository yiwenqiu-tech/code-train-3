package bit_operation

func reverseBits(num uint32) uint32 {
	var ans uint32
	for i := 0; i < 32; i++ { // 注意需要遍历32位
		ans = ans << 1
		ans += num & 1
		num = num >> 1
	}
	return ans
}
