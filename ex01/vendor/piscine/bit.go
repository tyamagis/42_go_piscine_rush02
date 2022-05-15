package piscine

func PopCount(u uint64) int {
	u = (u & 0b0101010101010101010101010101010101010101010101010101010101010101) +
		((u & 0b1010101010101010101010101010101010101010101010101010101010101010) >> 1)
	u = (u & 0b0011001100110011001100110011001100110011001100110011001100110011) +
		((u & 0b1100110011001100110011001100110011001100110011001100110011001100) >> 2)
	u = (u & 0x0f0f0f0f0f0f0f0f) +
		((u & 0xf0f0f0f0f0f0f0f0) >> 4)
	u = (u & 0x00ff00ff00ff00ff) +
		((u & 0xff00ff00ff00ff00) >> 8)
	u = (u & 0x0000ffff0000ffff) +
		((u & 0xffff0000ffff0000) >> 16)
	u = (u & 0x00000000ffffffff) +
		((u & 0xffffffff00000000) >> 32)
	return int(u)
}
