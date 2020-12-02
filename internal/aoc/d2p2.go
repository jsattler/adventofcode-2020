package aoc

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func Day2Part2(){
	content, err := ioutil.ReadFile("assets/input-2")
	var counter int

	if err != nil { log.Fatal(err) }

	// split input by new line
	strArr := strings.Split(string(content), "\n")

	for _, s := range strArr {
		policy := parsePolicy(s)
		if policy.isValidPart2() { counter++ }
	}
	fmt.Printf("solution day 2 part 2: %d\n", counter)
}

func (p *policy) isValidPart2() bool {
	if p.min - 1 < 0 || p.min > len(p.pwd) || p.max - 1 < 0 || p.max > len(p.pwd) {
		return false
	}
	first := string(p.pwd[p.min - 1])
	second := string(p.pwd[p.max - 1])
	return (first == p.char) != (second == p.char)
}
