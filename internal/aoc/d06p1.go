package aoc

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func Day6Part2() {
	var counter int
	content, err := ioutil.ReadFile("assets/input-6")
	if err != nil {
		log.Fatal(err)
	}

	groups:= strings.Split(string(content), "\n\n")

	for _, group := range groups {
		chars := make(map[rune]bool)
		group := strings.ReplaceAll(group, "\n", "")
		// iterate over each rune within the group
		for _, char := range group {
			chars[char] = true
		}

		counter += len(chars)
	}


	fmt.Printf("solution day 6 part 1: %d\n", counter)

}

