package aoc

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func Day11Part2() {

	occupiedSeats := 0
	input, err := ioutil.ReadFile("assets/input-11")
	if err != nil {
		log.Fatal(err)
	}

	var stateChange = true
	currentState := strings.Split(string(input), "\n")

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
					_, occCount := visitNeighborsInView(i, j, currentState)
					if occCount == 0 {
						r := []rune(nextState[i])
						r[j] = occupied
						nextState[i] = string(r)
						stateChange = true
					}
				case occupied:
					// check occupied rule
					_, occCount := visitNeighborsInView(i, j, currentState)
					if occCount >= 5 {
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

	fmt.Println("Solution day 11 part 2", occupiedSeats)

}

// needs to be refactored into smaller functions
func visitNeighborsInView(x, y int, matrix []string) (emptyCount, occCount int) {
	emptyCount = 0
	occCount = 0

	startRow := x
	stopRow := x

	if x > 0 {
		startRow--
	}

	if x < len(matrix)-1 {
		stopRow++
	}

	for i := startRow; i <= stopRow; i++ {
		startCol := y
		stopCol := y

		if y > 0 {
			startCol--
		}

		if y < len(matrix[i])-1 {
			stopCol++
		}

		for j := startCol; j <= stopCol; j++ {

			if x == i && y == j {
				// own seat
				continue
			}

			if matrix[i][j] == empty {
				emptyCount++
			}
			if matrix[i][j] == occupied {
				occCount++
			}

			// we found a floor and need to check what is in this direction
			if matrix[i][j] == floor {
				moveX := i - x
				moveY := j - y

				n := i + moveX
				m := j + moveY

				for {

					// check if one of the coordinates is out of range
					if !(n >= 0 && n < len(matrix)) || !(m >= 0 && m < len(matrix[n])) {
						// we reached end of matrix
						break
					}

					if matrix[n][m] == empty {
						// we found an empty seat so we can stop
						emptyCount++
						break
					}

					if matrix[n][m] == occupied {
						// we found an occupied seat in view so we can stop
						occCount++
						break
					}

					if matrix[n][m] == floor {
						// we found another floor, so we need to continue
						n += moveX
						m += moveY
						continue
					}

				}

			}
		}
	}

	return emptyCount, occCount
}

func printRows(rows []string) {
	for _, row := range rows {
		fmt.Println(row)
	}
}
