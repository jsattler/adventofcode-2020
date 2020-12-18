package main

import (
	"fmt"
	"github.com/j-sattler/adventofcode-2020/internal/aocio"
	"strings"
)

func Day06Part2() {
	groups := aocio.ReadFileSplit("../../assets/input-06", "\n\n")

	var counter int
	for _, group := range groups {
		chars := make(map[rune]int)
		groupSize := strings.Count(group, "\n") + 1
		group := strings.ReplaceAll(group, "\n", "")
		// iterate over each rune within the group
		for _, char := range group {
			chars[char]++
		}

		for _, v := range chars {
			if v == groupSize {
				counter++
			}
		}
	}

	fmt.Printf("Solution day 6 part 2: %d\n", counter)

}
