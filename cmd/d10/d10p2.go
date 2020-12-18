package main

import (
	"fmt"
	"github.com/j-sattler/adventofcode-2020/internal/aocio"
	"github.com/j-sattler/adventofcode-2020/internal/aocparse"
	"sort"
)

var cache map[int]int64

func Day10Part2() {
	joltsStr := aocio.ReadFileSplit("../../assets/input-10", "\n")
	jolts := aocparse.ParseInts(joltsStr)
	cache = make(map[int]int64, len(jolts))

	sort.Ints(jolts)

	result := adapt(0, jolts)

	fmt.Printf("Solution day 10 part 2: %d\n", result)

}

func adapt(jolt int, jolts []int) int64 {

	pathCounter := int64(0)

	if len(jolts) == 0 {
		return 1
	}

	if v, ok := cache[jolt]; ok {
		return v
	}

	if jolts[0]-jolt > 3 {
		return 0
	}

	for i, v := range jolts {
		if v > jolt+3 {
			break
		}

		pathCounter += adapt(v, jolts[i+1:])
	}
	cache[jolt] = pathCounter
	return pathCounter

}
