package main

import (
	"fmt"
	"github.com/j-sattler/adventofcode-2020/internal/aocio"
	"github.com/j-sattler/adventofcode-2020/internal/aocparse"
	"sort"
)

func Day10Part1() {
	joltsStr := aocio.ReadFileSplit("../../assets/input-10", "\n")
	jolts := aocparse.ParseInts(joltsStr)

	sort.Ints(jolts)

	countThree := 0
	countOne := 0
	last := 0

	for _, jolt := range jolts {
		if diff := jolt - last; diff <= 3 {
			if diff == 3 {
				countThree++
			}
			if diff == 1 {
				countOne++
			}
			last = jolt
		}
	}

	countThree++

	fmt.Printf("Solution day 10 part 1: %d\n", countThree*countOne)

}
