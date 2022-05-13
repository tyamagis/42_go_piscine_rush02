package piscine

import "fmt"

// blocksの各要素([]string)が正しいテトリミノ定義であることを確認する。
// check that every elements in block is valid as tetrimino definition.
func Validate(core *Core, blocks [][]string) bool {
	return Every(blocks, func(lines []string, _ int) bool {
		// []stringの長さが4であるか？
		// size of lines is 4?
		h := SliceLen(lines)
		if h != MinoSize {
			return false
		}
		// 各要素の長さが4であるか？
		// size of every rows are 4?
		ok := Every(lines, func(line string, _ int) bool {
			w := SliceLen([]rune(line))
			return w == MinoSize
		})
		if !ok {
			return false
		}

		// 3. すべての文字が`#`か`.`であるか
		// every characters are '#' or '.'?
		runeCount := map[rune]int{}
		for _, r := range Join(lines, "") {
			runeCount[r] = runeCount[r] + 1
		}
		hasInvalidChar := runeCount['#']+runeCount['.'] != MinoSize*MinoSize
		if hasInvalidChar {
			return false
		}
		// 3. `#`がちょうど4つあるか？
		// `lines`` contains just 4 '#'?
		sharps := runeCount['#']
		if sharps != 4 {
			return false
		}

		// 4. `#`の配置がテトリミノとして正しいこと
		// '#' arrangement is valid as a tetrimino?
		trimmed := TrimBlankRowCol(lines)
		joined := Join(trimmed, "\n")
		// joinedは標準形になっている
		minoType, exists := core.MinoMap[joined]
		fmt.Println(minoType, exists)
		fmt.Println(joined)
		if !exists {
			return false
		}
		return true
	})
}
