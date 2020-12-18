package main

import (
	"fmt"
	"github.com/j-sattler/adventofcode-2020/internal/aocio"
)

type Move struct {
	right, down int
}

func Day03Part2() {
	strArr := aocio.ReadFileSplit("../../assets/input-03", "\n")

	moves := [5]Move{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	treeCounter := 0
	mod := len(strArr[0])
	prod := 1

	for _, v := range moves {
		col := 0
		// iterate over each row
		for row := 0; row < len(strArr)-1; {
			col += v.right
			row += v.down
			if row > len(strArr) {
				break
			}
			// if out of bound continue at front
			if string(strArr[row][col%mod]) == "#" {
				treeCounter++
			}
		}
		prod *= treeCounter
		treeCounter = 0
	}
	fmt.Printf("solution day 3 part 2: %d\n", prod)
}
