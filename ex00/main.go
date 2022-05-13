package main

import (
	"fmt"
	"os"
	"piscine"
)

func main() {

	core := piscine.NewCore()
	fmt.Println(core.Minos)

	n := piscine.SliceLen(os.Args)
	if n < 2 {
		return
	}
	path := os.Args[1]
	content, _ := piscine.Read(path)
	blocks, _ := piscine.Divide(content)
	fmt.Println(blocks)
	fmt.Println(piscine.Validate(blocks))
}
