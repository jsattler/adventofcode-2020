package main

import (
	"fmt"
	"github.com/j-sattler/adventofcode-2020/internal/aocio"
	"strings"
)

func Day06Part1() {
	groups := aocio.ReadFileSplit("../../assets/input-06", "\n\n")
	var counter int
	for _, group := range groups {
		chars := make(map[rune]bool)
		group := strings.ReplaceAll(group, "\n", "")
		// iterate over each rune within the group
		for _, char := range group {
			chars[char] = true
		}

		counter += len(chars)
	}

	fmt.Printf("Solution day 6 part 1: %d\n", counter)

}
