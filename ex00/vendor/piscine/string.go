package piscine

func StringHasRune(s string, r rune) bool {
	for _, c := range s {
		if c == r {
			return true
		}
	}
	return false
}

// LU -> 2x3

// 縦 4 - 2 + 1 = 3パターン(0, 1, 2)
// 横 4 - 3 + 1 = 2パターン(0, 1)
