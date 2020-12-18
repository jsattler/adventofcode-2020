package main

import (
	"fmt"
	"github.com/j-sattler/adventofcode-2020/internal/aocio"
	"log"
	"strconv"
	"strings"
)

func Day15Part1() {
	initStr := aocio.ReadFileSplit("../../assets/input-15", ",")
	numbers := parseNumbers(initStr)

	for i := len(numbers); i < 2020; i++ {
		current := 0
		prev := numbers[i-1]

		if index, present := lastIndexOf(prev, numbers[:i-1]); present {
			current = i - (index + 1) // previous turn - last turn before that
		}
		numbers = append(numbers, current)
	}

	fmt.Printf("Solution day 15 part 1: %d\n", numbers[len(numbers)-1])
}

func lastIndexOf(value int, slice []int) (index int, present bool) {
	for i, v := range slice {
		if v == value {
			index = i
			present = true
		}
	}
	return
}

func parseNumbers(in []string) (nums []int) {
	for _, str := range in {
		str = strings.ReplaceAll(str, "\n", "")
		num, err := strconv.Atoi(str)

		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)
	}
	return
}
