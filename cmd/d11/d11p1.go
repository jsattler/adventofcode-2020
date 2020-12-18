package main

import (
	"fmt"
	"github.com/j-sattler/adventofcode-2020/internal/aocio"
	"strings"
)

const (
	floor    = '.'
	occupied = '#'
	empty    = 'L'
)

func Day11Part1() {
	currentState := aocio.ReadFileSplit("../../assets/input-11", "\n")
	occupiedSeats := 0
	stateChange := true
	nextState := make([]string, len(currentState))

	copy(nextState, currentState)
	for stateChange {
		stateChange = false

		for i, row := range currentState {
			// within this iteration currentState does not get modified

			for j, seat := range row {

				switch seat {
				case empty:
					// check empty rule i, j
					_, occCount := visitNeighbors(i, j, currentState)
					if occCount == 0 {
						r := []rune(nextState[i])
						r[j] = occupied
						nextState[i] = string(r)
						stateChange = true
					}
				case occupied:
					// check occupied rule
					_, occCount := visitNeighbors(i, j, currentState)
					if occCount-1 >= 4 {
						r := []rune(nextState[i])
						r[j] = empty
						nextState[i] = string(r)
						stateChange = true
					}
				case floor:
					continue
				}

			}
		}
		copy(currentState, nextState)
	}

	for _, row := range nextState {
		occupiedSeats += strings.Count(row, string(occupied))
	}

	fmt.Println("Solution day 11 part 1", occupiedSeats)

}

// visits each neighbor of the given coordinate and returns number of occupied and free seats
func visitNeighbors(i, j int, matrix []string) (emptyCount, occCount int) {
	emptyCount = 0
	occCount = 0

	startRow := i
	stopRow := i

	if i > 0 {
		startRow--
	}

	if i < len(matrix)-1 {
		stopRow++
	}

	for _, row := range matrix[startRow : stopRow+1] {
		startCol := j
		stopCol := j

		if j > 0 {
			startCol--
		}

		if j < len(row)-1 {
			stopCol++
		}

		for _, col := range row[startCol : stopCol+1] {

			if col == empty {
				emptyCount++
			}
			if col == occupied {
				occCount++
			}
		}
	}

	return emptyCount, occCount
}
