package aoc

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Move struct {
	right, down int
}

func Day3Part2(){
	content, err := ioutil.ReadFile("assets/input-3")

	var treeCounter int
	moves := [5]Move {
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	if err != nil { log.Fatal(err) }

	// split input by new line
	strArr := strings.Split(string(content), "\n")

	mod := len(strArr[0])
	prod := 1

	for _, v := range moves {
		col := 0
		// iterate over each row
		for row := 0; row < len(strArr) - 1; {
			col += v.right
			row += v.down
			if row > len(strArr) {
				break
			}
			// if out of bound continue at front
			if string(strArr[row][col % mod]) == "#" {
				treeCounter++
			}
		}
		prod *= treeCounter
		treeCounter = 0
	}
	fmt.Printf("solution day 3 part 2: %d\n", prod)
}