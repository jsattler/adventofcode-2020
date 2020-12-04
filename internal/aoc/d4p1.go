package aoc

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)


func Day4Part1(){
	var counter int
	content, err := ioutil.ReadFile("assets/input-4")
	if err != nil { log.Fatal(err) }
	fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	strArr := strings.Split(string(content), "\n\n")

	for _, str := range strArr {

		isValid := true

		for _, field := range fields {
			if !strings.Contains(str, field) {
				// passport does not contain field
				isValid = false
				break
			}
		}

		if isValid { counter++ }

	}
	fmt.Printf("solution day 4 part 1: %d\n", counter)
}
