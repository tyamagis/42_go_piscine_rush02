package piscine

import "fmt"

func Solve(core *Core) {
	for n := 4; true; n++ {
		fmt.Printf("try for size %d\n", n)
		if solveForSize(core, n) {
			fmt.Println("solved")
			return
		}
		fmt.Println("failed")
	}
}

type SolverState struct {
	Size  int
	Board [][]rune
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
	}
	return dfs(core, ss, 0)
}

func dfs(core *Core, state *SolverState, k int) bool {
	if k == len(core.GivenMinos) {
		PrintBoard(state)
		return true
	}
	mt := core.GivenMinos[k]
	mino := core.MinoReverseMap[mt]
	// fmt.Println(mino)

	for i := 0; i < state.Size-mino.Height+1; i++ {
		for j := 0; j < state.Size-mino.Width+1; j++ {
			if !isPlacableAt(state, mino, i, j) {
				continue
			}
			placeAt(state, mino, i, j, k)
			if dfs(core, state, k+1) {
				return true
			}
			removeFrom(state, mino, i, j)
		}
	}

	return false

}
