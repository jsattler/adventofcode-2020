package aoc

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type policy struct {
	min, max int
	char string
	pwd string
}

func Day2Part1(){

	content, err := ioutil.ReadFile("assets/input-2")
	var counter int

	if err != nil { log.Fatal(err) }

	// split input by new line
	strArr := strings.Split(string(content), "\n")

	for _, s := range strArr {
		policy := parsePolicy(s)
		if policy.isValid() { counter++ }
	}
	fmt.Printf("%d passwords are valid\n", counter)
}
func parsePolicy(pStr string) *policy {
	splitBySpace := strings.Split(pStr, " ")
	minMax := strings.Split(splitBySpace[0], "-")
	char := strings.Split(splitBySpace[1], ":")[0]
	passwd := splitBySpace[2]

	min, err:= strconv.Atoi(minMax[0])
	if err != nil { log.Fatal(err) }

	max, err := strconv.Atoi(minMax[1])
	if err != nil { log.Fatal(err) }

	return &policy{
		min:  min,
		max:  max,
		char: char,
		pwd:  passwd,
	}
}

func (p *policy) isValid() bool {
	count := strings.Count(p.pwd, p.char)
	return count >= p.min && count <= p.max
}
