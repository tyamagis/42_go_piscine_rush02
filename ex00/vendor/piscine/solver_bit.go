package piscine

import "fmt"

// 1 <= size <= 64
func makeBitBoard(size int) []uint64 {
	return Map(Seq(0, size), func(_, _ int) uint64 {
		return 0
	})
}

func (ss *SolverState) bitPrintBoard(core *Core) {
	board := ss.makeBitBoardString(core)
	ForEach(board, func(rs []rune, _ int) {
		fmt.Println(string(rs))
	})
}

func (ss *SolverState) makeBitBoardString(core *Core) [][]rune {
	board := Map(Seq(0, ss.Size), func(_, _ int) []rune {
		row := make([]rune, ss.Size)
		for i, _ := range Seq(0, ss.Size) {
			row[i] = '.'
		}
		return row
	})
	for pos, k := range ss.Placement {
		i, j := pos/ss.Size, pos%ss.Size
		t := core.GivenMinos[k]
		mino := core.MinoReverseMap[t]
		r := 'A' + rune(k)
		for ii, row := range mino.shape {
			for jj, c := range row {
				if c != '#' {
					continue
				}
				board[i+ii][j+jj] = r
			}
		}
	}
	return board
}

func (ss *SolverState) isBitPlacableAt(mino *MinoMaster, i, j int) bool {
	if len(mino.bitVShape) <= len(mino.bitHShape) {
		for ii, bitRow := range mino.bitVShape {
			mask := (bitRow << j)
			if (ss.BitVBoard[i+ii] & mask) != 0 {
				return false
			}
		}
	} else {
		for jj, bitCol := range mino.bitHShape {
			mask := (bitCol << i)
			if (ss.BitHBoard[j+jj] & mask) != 0 {
				return false
			}
		}
	}
	return true
}

func (ss *SolverState) bitPlaceAt(mino *MinoMaster, i, j, k int) {
	ss.Placement[i*ss.Size+j] = k
	Push(ss.PositionStack[mino.MinoType], i*ss.Size+j)
	ss.RestMinoCount[k] = ss.RestMinoCount[k] - 1
	for ii, bitRow := range mino.bitVShape {
		ss.BitVBoard[i+ii] |= (bitRow << j)
	}
	for jj, bitCol := range mino.bitHShape {
		ss.BitHBoard[j+jj] |= (bitCol << i)
	}
}

func (ss *SolverState) bitRemoveFrom(mino *MinoMaster, i, j, k int) {
	delete(ss.Placement, i*ss.Size+j)
	Pop(ss.PositionStack[mino.MinoType])
	ss.RestMinoCount[k] = ss.RestMinoCount[k] + 1
	for ii, bitRow := range mino.bitVShape {
		ss.BitVBoard[i+ii] ^= (bitRow << j)
	}
	for jj, bitCol := range mino.bitHShape {
		ss.BitHBoard[j+jj] ^= (bitCol << i)
	}
}

func (ss *SolverState) getRestVacantOver(i, j int) int {
	n := 0
	for ii := i; ii < ss.Size; ii++ {
		var mask uint64 = (1 << ss.Size) - 1
		if ii == i {
			mask = (1 << ss.Size) - (1 << j)
		}
		n += PopCount((^ss.BitVBoard[ii]) & mask)
		// fmt.Println(ii, mask, n)
	}
	return n
}
