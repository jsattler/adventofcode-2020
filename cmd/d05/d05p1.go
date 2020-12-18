package main

import (
	"fmt"
	"github.com/j-sattler/adventofcode-2020/internal/aocio"
	"math"
)

func Day05Part1() {
	passes := aocio.ReadFileSplit("../../assets/input-05", "\n")
	highest := math.MinInt32

	for _, pass := range passes {
		row := 0
		col := 0
		for _, v := range pass {
			switch v {
			case 'B': // shift row 1 and set to 1
				row <<= 1
				row |= 1
			case 'F': // shift row 1 keep 0
				row <<= 1
			case 'R': // shift col 1 set to 1
				col <<= 1
				col |= 1
			case 'L': // shift col 1 keep 0
				col <<= 1
			}
		}

		if newHigh := row*8 + col; newHigh > highest {
			highest = newHigh
		}
	}

	fmt.Printf("solution day 5 part 1: %d\n", highest)

}
