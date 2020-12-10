package aoc

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const (
	preamble = 25
)


// Find first number that is not the sum of two of previous 25 numbers
func Day9Part1() {

	input, err := ioutil.ReadFile("assets/input-9")
	if err != nil {
		log.Fatal(err)
	}

	numbers := strings.Split(string(input), "\n")

	// mapping from current preamble to previous sums
	cachedSums := make([][]int, len(numbers))

	// fill array
	for i, x := range numbers {
		xi := stringToInt(x)
		cachedSums[i] = make([]int, len(numbers))
		for j := i ; j < len(numbers); j++{
			y := numbers[j]
			yi := stringToInt(y)
			cachedSums[i][j] = xi + yi
		}
	}

	// iterate again over numbers starting at preamble index
	for i := preamble; i < len(numbers); i++ {
		sum := stringToInt(numbers[i])
		isPresent := false
		for j := i - preamble; j < len(cachedSums); j++ {
			for k := 0; k < len(cachedSums[j]); k++ {
				if cachedSums[j][k] == sum {
					isPresent = true
				}
			}
		}

		if !isPresent {
			fmt.Printf("Solution day 9 part 1: %d\n", sum)
			break
		}
	}

}

func stringToInt(str string) int {
	val, err := strconv.Atoi(str)

	if err != nil {
		log.Fatal("failed to convert string to int")
	}
	return val
}