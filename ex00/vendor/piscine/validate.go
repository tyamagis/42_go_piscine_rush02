package piscine

func Validate(blocks [][]string) bool {
	// blocksの各要素([]string)が正しいテトリミノ定義であることを確認する。

	return Every(blocks, func(lines []string, _ int) bool {
		// 1. []stringの長さが4であること
		h := SliceLen(lines)
		if h != MinoSize {
			return false
		}
		// 2. 各要素の長さが4であること
		ok := Every(lines, func(line string, _ int) bool {
			w := SliceLen([]rune(line))
			return w == MinoSize
		})
		if !ok {
			return false
		}
		// 3. `#`の配置がテトリミノとして正しいこと
		return true
	})
}
