package aoc

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func Day3Part1(){
	content, err := ioutil.ReadFile("assets/input-3")
	var treeCounter int
	if err != nil { log.Fatal(err) }

	// split input by new line
	strArr := strings.Split(string(content), "\n")

	mod := len(strArr[0])
	col := 0

	// iterate over each row
	for row:= 0; row < len(strArr) - 1; {
		col += 3
		row++
		// if out of bound continue at front
		if string(strArr[row][col % mod]) == "#" {
			treeCounter++
		}
	}

	fmt.Printf("solution day 3 part 1: %d\n", treeCounter)

}