package main

import (
	"fmt"
	"github.com/j-sattler/adventofcode-2020/internal/aocio"
)

func Day02Part2() {
	strArr := aocio.ReadFileSplit("../../assets/input-02", "\n")
	counter := 0
	for _, s := range strArr {
		policy := parsePolicy(s)
		if policy.isValidPart2() {
			counter++
		}
	}
	fmt.Printf("solution day 2 part 2: %d\n", counter)
}

func (p *Policy) isValidPart2() bool {
	if p.min-1 < 0 || p.min > len(p.pwd) || p.max-1 < 0 || p.max > len(p.pwd) {
		return false
	}
	first := string(p.pwd[p.min-1])
	second := string(p.pwd[p.max-1])
	return (first == p.char) != (second == p.char)
}
