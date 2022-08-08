package sol

func reverseBits(num uint32) uint32 {
	var ans uint32 = 0
	for pos := 0; pos < 32; pos++ {
		ans = ans << 1
		if (num & 1) != 0 {
			ans += 1
		}
		num = num >> 1
	}
	return ans
}
