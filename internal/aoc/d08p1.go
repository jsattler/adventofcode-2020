package aoc

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var accumulator = 0

const (
	acc = "acc" // accumulate - increase/decrease accumulator with given amount
	jmp = "jmp" // jump to a new instruction (+1 is next)
	nop = "nop" // no operation, will do nothing and continue with next
)

func Day8Part1() {

	input, err := ioutil.ReadFile("assets/input-8")
	if err != nil {
		log.Fatal(err)
	}

	// each line consists of operation and an argument (signed number)
	instructions := strings.Split(string(input), "\n")
	visited := make(map[int]bool, len(instructions))

	for i := 0; i <= len(instructions); {

		if _, ok := visited[i]; ok {
			break
		}

		op, arg := parseInstruction(instructions[i])

		visited[i] = true

		switch op {
		case acc:
			accumulator	+= arg
			i++
		case jmp:
			i += arg
		case nop:
			i++
		default:
			break
		}
	}

	fmt.Printf("Solution day 8 part 1: %d\n", accumulator)

}

func parseInstruction(instruction string) (string, int)  {
	ins := strings.Split(instruction, " ")
	arg, err := strconv.Atoi(ins[1])
	if err != nil {
		log.Fatal("Failed to convert string to int")
	}
	return ins[0], arg
}