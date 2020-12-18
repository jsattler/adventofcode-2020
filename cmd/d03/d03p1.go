package main

import (
	"fmt"
	"github.com/j-sattler/adventofcode-2020/internal/aocio"
)

func Day03Part1() {
	strArr := aocio.ReadFileSplit("../../assets/input-03", "\n")

	var treeCounter int
	mod := len(strArr[0])
	col := 0

	// iterate over each row
	for row := 0; row < len(strArr)-1; {
		col += 3
		row++
		// if out of bound continue at front
		if string(strArr[row][col%mod]) == "#" {
			treeCounter++
		}
	}

	fmt.Printf("solution day 3 part 1: %d\n", treeCounter)

}
