package aoc

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

const invalidNumber = 133015568

// Find first number that is not the sum of two of previous 25 numbers
func Day9Part2() {

	input, err := ioutil.ReadFile("assets/input-9")
	if err != nil {
		log.Fatal(err)
	}

	numbers := strings.Split(string(input), "\n")

	found := false
	indexI := 0
	indexJ := 0

	// fill array
	for i, x := range numbers {
		xi := stringToInt(x)
		sum := xi

		for j := i+1 ; j < len(numbers); j++{
			y := numbers[j]
			yi := stringToInt(y)
			sum += yi
			if sum == invalidNumber {
				found = true
				indexI = i
				indexJ = j
			}
		}

		if found {
			break
		}
	}

	ints := stringsToInts(numbers)
	fmt.Printf("Solution day 9 part 2: %d\n", sumMinMaxRange(indexI, indexJ, ints))

}

func stringsToInts(strArr []string) []int{
	ints := make([]int, len(strArr))
	for i, str := range strArr{
		newInt := stringToInt(str)
		ints[i] = newInt
	}
	return ints
}

// Sum the lowest and highest number within a given range of the array
func sumMinMaxRange(low, high int, arr []int) int {
	min := math.MaxInt32
	max := math.MinInt32

	for ; low < high ; {
		if newMin := arr[low]; newMin < min  {
			min = newMin
		}
		if newMax := arr[low]; newMax > max {
			max = newMax
		}
		low++
	}
	return min + max
}