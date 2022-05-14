package piscine

import (
	"fmt"
	"time"
)

const TargetPoint = "\x1b[42m\x1b[30m"
const TargetLine = "\x1b[47m\x1b[30m"
const TextBold = "\x1b[1m"
const Reset = "\x1b[m"
const CursorTop = "\x1b[H"
const ScreenClear = "\x1b[2J"

func Visualize(core *Core, state *SolverState, i, j, k int, comment string) {
	fmt.Print(CursorTop)

	head := " │"
	for x, _ := range Seq(0, state.Size) {
		head += string(rune(x%10) + '0')
	}
	fmt.Println(head)
	head = "─┼"
	for range Seq(0, state.Size) {
		head += "─"
	}
	fmt.Println(head)

	mark := 'A' + rune(k)
	for y, row := range state.makeBitBoardString(core) {
		s := ""
		s += string(rune(y%10) + '0')
		s += "│"
		for _, c := range row {
			if c == mark {
				s += TargetPoint + string(c) + Reset
			} else {
				s += string(c)
			}
		}
		fmt.Println(s)
	}
	st := time.Duration(500) * time.Millisecond
	time.Sleep(st)
}
