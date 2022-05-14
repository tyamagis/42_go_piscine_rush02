package piscine

import "fmt"

func Solve(core *Core) {
	fn := firstSize(core)
	fmt.Println(ScreenClear)
	for n := fn; true; n++ {
		fmt.Printf("try for size %d\n", n)
		if solveForSize(core, n) {
			fmt.Println("solved")
			return
		}
		fmt.Println("failed")
	}
}

type SolverState struct {
	Size                 int
	Board                [][]rune
	BitVBoard, BitHBoard []uint64
	Placement            map[int]int
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
		BitVBoard: makeBitBoard(size),
		BitHBoard: makeBitBoard(size),
		Placement: map[int]int{},
	}
	return dfs(core, ss, 0)
}

func dfs(core *Core, state *SolverState, k int) bool {
	if k == len(core.GivenMinos) {
		state.bitPrintBoard(core)
		return true
	}
	mt := core.GivenMinos[k]
	mino := core.MinoReverseMap[mt]
	// fmt.Println(mino)

	for i := 0; i < state.Size-mino.Height+1; i++ {
		for j := 0; j < state.Size-mino.Width+1; j++ {
			if !state.isBitPlacableAt(mino, i, j) {
				continue
			}
			state.bitPlaceAt(mino, i, j, k)
			// fmt.Println(state.BitVBoard)
			Visualize(core, state, i, j, k, "put")
			if dfs(core, state, k+1) {
				return true
			}
			state.bitRemoveFrom(mino, i, j)
		}
	}

	return false

}
