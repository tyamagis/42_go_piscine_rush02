package piscine

import "fmt"

func PrintBoard(ss *SolverState) {
	ForEach(ss.Board, func(rs []rune, _ int) {
		fmt.Println(string(rs))
	})
}

func isPlacableAt(ss *SolverState, mino *MinoMaster, i, j int) bool {
	for ii, row := range mino.shape {
		for jj, c := range row {
			if c != '#' {
				continue
			}
			if ss.Board[i+ii][j+jj] != '.' {
				return false
			}
		}
	}
	return true
}

func placeAt(ss *SolverState, mino *MinoMaster, i, j, k int) {
	r := 'A' + rune(k)
	for ii, row := range mino.shape {
		for jj, c := range row {
			if c != '#' {
				continue
			}
			ss.Board[i+ii][j+jj] = r
		}
	}
}

func removeFrom(ss *SolverState, mino *MinoMaster, i, j int) {
	for ii, row := range mino.shape {
		for jj, c := range row {
			if c != '#' {
				continue
			}
			ss.Board[i+ii][j+jj] = '.'
		}
	}
}
