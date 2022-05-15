package piscine

import "fmt"

func Solve(core *Core) {
	fn := firstSize(core)
	fmt.Println(ScreenClear)
	nMino := len(core.GivenMinos)
	for n := fn; true; n++ {
		nVacant := n*n - nMino*4
		fmt.Printf("try for size %d, minos = %d(%c - %c), vacants = %d\n",
			n, nMino,
			'A', 'A'+rune(nMino)-1,
			nVacant,
		)
		if solveForSize(core, n) {
			fmt.Println("solved")
			return
		}
		fmt.Println("failed")
	}
}

type SolverState struct {
	// サイズ
	Size int
	// 計算中の盤面(rune)
	Board [][]rune
	// 計算中の盤面(bit)
	BitVBoard, BitHBoard []uint64
	// 位置をkey, ミノ種別をvalueとするマップ
	Placement map[int]int
	// ミノ種別ごとの位置のスタック
	PositionStack map[int]*[]int
	// ミノ種別をkey, 残りの数をvalueとするマップ
	RestMinoCount map[int]int
}

func solveForSize(core *Core, size int) bool {
	ss := &SolverState{
		Size: size,
		Board: Map(Seq(0, size), func(_, _ int) []rune {
			row := make([]rune, size, size)
			for j := 0; j < size; j++ {
				row[j] = '.'
			}
			return row
		}),
		BitVBoard:     makeBitBoard(size),
		BitHBoard:     makeBitBoard(size),
		Placement:     map[int]int{},
		PositionStack: map[int]*[]int{},
		RestMinoCount: map[int]int{},
	}
	ForEach(core.Minos, func(k, _ int) {
		s := make([]int, 0, size)
		ss.PositionStack[k] = &s
	})
	ForEach(core.GivenMinos, func(k, _ int) {
		ss.RestMinoCount[k] = ss.RestMinoCount[k] + 1
	})
	return dfs(core, ss, 0)
}

func dfs(core *Core, state *SolverState, k int) bool {
	if k == len(core.GivenMinos) {
		state.bitPrintBoard(core)
		return true
	}
	restVacant := state.getRestVacantPlacable()
	restToPut := 4 * (len(core.GivenMinos) - k)
	if restVacant < restToPut {
		Visualize(core, state, -1, -1, k,
			fmt.Sprintf("no enough vacant: needed %d, actual %d", restToPut, restVacant),
		)
		return false
	}

	mt := core.GivenMinos[k]
	mino := core.MinoReverseMap[mt]
	// fmt.Println(mino)

	i, j := 0, 0
	if len(*state.PositionStack[mt]) > 0 {
		pos := Top(state.PositionStack[mt]) + 1
		i, j = pos/state.Size, pos%state.Size
	}

	for ; i < state.Size-mino.Height+1; i++ {
		for ; j < state.Size-mino.Width+1; j++ {
			restVacant := state.getRestVacantOver(i, j)
			restToPut := state.RestMinoCount[k] * 4
			if restVacant < restToPut {
				Visualize(core, state, i, j, k, fmt.Sprintf("no enough vacant for %d", k))
				return false
			}

			if !state.isBitPlacableAt(mino, i, j) {
				// Visualize(core, state, i, j, k, fmt.Sprintf("not placable %d at (%d, %d)", k, i, j))
				continue
			}

			state.bitPlaceAt(mino, i, j, k)
			// fmt.Println(state.BitVBoard)
			Visualize(core, state, i, j, k, fmt.Sprintf("place %d at (%d, %d)", k, i, j))
			if dfs(core, state, k+1) {
				return true
			}
			state.bitRemoveFrom(mino, i, j, k)
		}
		j = 0
	}

	return false

}
