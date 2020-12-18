package main

import (
	"fmt"
	"github.com/j-sattler/adventofcode-2020/internal/aocio"
	"strings"
)

func Day04Part1() {
	strArr := aocio.ReadFileSplit("../../assets/input-04", "\n\n")
	var counter int
	fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for _, str := range strArr {

		isValid := true

		for _, field := range fields {
			if !strings.Contains(str, field) {
				// passport does not contain field
				isValid = false
				break
			}
		}

		if isValid {
			counter++
		}

	}
	fmt.Printf("solution day 4 part 1: %d\n", counter)
}
