package main

import (
	"fmt"
	"github.com/j-sattler/adventofcode-2020/internal/aocio"
)

func Day08Part2() {

	instructions := aocio.ReadFileSplit("../../assets/input-08", "\n")
	isRunning := true

	for opIndex := 0; isRunning; opIndex++ {
		var accumulator = 0
		visited := make(map[int]bool, len(instructions))
		isSwitched := false
		for i := 0; i <= len(instructions); {
			if _, ok := visited[i]; ok {
				break
			}

			op, arg := parseInstruction(instructions[i])
			if i == opIndex && (op == nop || op == jmp) && !isSwitched {
				opIndex = i
				if op == nop {
					op = jmp
				} else {
					op = nop
				}
				isSwitched = true
			}

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

			if i >= len(instructions) {
				isRunning = false
				fmt.Printf("Solution day 8 part 2: %d\n", accumulator)
				break
			}

		}

	}

}
