package main

import (
	"flag"
	"os"
	"piscine"
)

func main() {
	core := piscine.NewCore()
	if piscine.GFlag != 0 {
		piscine.Generate(core, piscine.GFlag)
		return
	}
	n := piscine.SliceLen(os.Args)
	if n != 2 {
		piscine.ShowUsage()
		return
	}
	path := os.Args[flag.NFlag()+1]
	content, _ := piscine.Read(path)
	blocks, _ := piscine.Divide(content)
	if !piscine.Validate(core, blocks) {
		return
	}
	core.GivenMinos = piscine.Symbolize(core, blocks)
	// fmt.Println(core)

	piscine.Solve(core)
}
