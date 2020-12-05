package aoc

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

func Day5Part1() {
	content, err := ioutil.ReadFile("assets/input-5")
	if err != nil {
		log.Fatal(err)
	}

	passes := strings.Split(string(content), "\n")

	highest := math.MinInt32

	for _ , pass := range passes{
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

		if newHigh := row * 8 + col; newHigh > highest {
			highest = newHigh
		}
	}

	fmt.Printf("solution day 5 part 2: %d\n", highest)

}

