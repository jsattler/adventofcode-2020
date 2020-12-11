package aoc

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)


func Day10Part1() {

	input, err := ioutil.ReadFile("assets/input-10")
	if err != nil {
		log.Fatal(err)
	}

	joltsStr := strings.Split(string(input), "\n")

	jolts := stringsToInts(joltsStr)

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

	fmt.Printf("Solution day 10 part 1: %d\n", countThree * countOne)

}

