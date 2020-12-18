package main

import (
	"fmt"
	"github.com/j-sattler/adventofcode-2020/internal/aocio"
	"log"
	"strconv"
	"strings"
)

type Policy struct {
	min, max int
	char     string
	pwd      string
}

func Day02Part1() {
	strArr := aocio.ReadFileSplit("../../assets/input-02", "\n")
	counter := 0
	for _, s := range strArr {
		policy := parsePolicy(s)
		if policy.isValid() {
			counter++
		}
	}
	fmt.Printf("solution day 2 part 1: %d\n", counter)
}
func parsePolicy(pStr string) *Policy {
	splitBySpace := strings.Split(pStr, " ")
	minMax := strings.Split(splitBySpace[0], "-")
	char := strings.Split(splitBySpace[1], ":")[0]
	passwd := splitBySpace[2]

	min, err := strconv.Atoi(minMax[0])
	if err != nil {
		log.Fatal(err)
	}

	max, err := strconv.Atoi(minMax[1])
	if err != nil {
		log.Fatal(err)
	}

	return &Policy{
		min:  min,
		max:  max,
		char: char,
		pwd:  passwd,
	}
}

func (p *Policy) isValid() bool {
	count := strings.Count(p.pwd, p.char)
	return count >= p.min && count <= p.max
}
