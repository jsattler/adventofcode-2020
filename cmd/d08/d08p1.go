package main

import (
	"fmt"
	"github.com/j-sattler/adventofcode-2020/internal/aocio"
	"log"
	"strconv"
	"strings"
)

const (
	acc = "acc" // accumulate - increase/decrease accumulator with given amount
	jmp = "jmp" // jump to a new instruction (+1 is next)
	nop = "nop" // no operation, will do nothing and continue with next
)

func Day08Part1() {
	instructions := aocio.ReadFileSplit("../../assets/input-08", "\n")
	accumulator := 0
	visited := make(map[int]bool, len(instructions))

	for i := 0; i <= len(instructions); {

		if _, ok := visited[i]; ok {
			break
		}

		op, arg := parseInstruction(instructions[i])

		visited[i] = true

		switch op {
		case acc:
			accumulator += arg
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

func parseInstruction(instruction string) (string, int) {
	ins := strings.Split(instruction, " ")
	arg, err := strconv.Atoi(ins[1])
	if err != nil {
		log.Fatal(err)
	}
	return ins[0], arg
}
