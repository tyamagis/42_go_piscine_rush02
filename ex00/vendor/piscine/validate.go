package piscine

// blocksの各要素([]string)が正しいテトリミノ定義であることを確認する。
// check that every elements in block is valid as tetrimino definition.
func Validate(blocks [][]string) bool {

	return Every(blocks, func(lines []string, _ int) bool {
		// []stringの長さが4であるか？
		// size of lines is 4?
		h := SliceLen(lines)
		if h != MinoSize {
			return false
		}
		// 各要素の長さが4であること？
		// size of every rows are 4?
		ok := Every(lines, func(line string, _ int) bool {
			w := SliceLen([]rune(line))
			return w == MinoSize
		})
		if !ok {
			return false
		}

		// 3. すべての文字が`#`か`.`であること
		// every characters are '#' or '.'?
		hasInvalidChar := Some([]rune(Join(lines, "")), func(r rune, _ int) bool {
			if r == '#' {
				return false
			}
			if r == '.' {
				return false
			}
			return true
		})
		if hasInvalidChar {
			return false
		}

		// 3. `#`がちょうど4つあること
		// `lines`` contains just 4 '#'?
		sharps := Reduce(lines, 0, func(s int, line string, _ int) int {
			return Reduce([]rune(line), 0, func(s int, r rune, _ int) int {
				if r == '#' {
					return s + 1
				}
				return s
			}) + s
		})
		if sharps != 4 {
			return false
		}

		// 4. `#`の配置がテトリミノとして正しいこと
		// '#' arrangement is valid as a tetrimino?
		xs := []int{0, -1, 0, 1}
		ys := []int{1, 0, -1, 0}
		for i, row := range lines {
			for j, c := range row {
				if c == '#' {
					// '#' must be adjacent to another '#'
					validBlack := Some([]int{0, 1, 2, 3}, func(k int, _ int) bool {
						dx, dy := xs[k], ys[k]
						ii, jj := i+dy, j+dx
						return 0 <= ii && ii < MinoSize && 0 <= jj && jj < MinoSize && lines[ii][jj] == '#'
					})
					if !validBlack {
						return false
					}
				}
			}
		}
		return true
	})
}
