package main

import (
	"fmt"
	"github.com/j-sattler/adventofcode-2020/internal/aocio"
	"log"
	"strconv"
	"strings"
)

func Day14Part1() {
	splitByLine := aocio.ReadFileSplit("../../assets/input-14", "\n")
	res := inspectProgram(splitByLine)
	fmt.Printf("Solution day 14 part 1: %d\n", res)
}

func inspectProgram(input []string) int {
	memory := make(map[int]int)
	var maskIndex map[int]int
	for _, line := range input {
		if strings.HasPrefix(line, "mask") {
			maskIndex = make(map[int]int)
			maskStr := strings.ReplaceAll(line, "mask = ", "")
			index := 0
			for i := len(maskStr) - 1; i >= 0; i-- {
				if maskStr[i] == 'X' {
					index++
					continue
				}
				val, err := strconv.Atoi(string(maskStr[i]))
				if err != nil {
					log.Fatal(err)
				}
				maskIndex[index] = val
				index++
			}

		} else {
			memAddr := 0
			val := 0
			_, err := fmt.Sscanf(line, "mem[%d] = %d", &memAddr, &val)
			if err != nil {
				log.Fatal(err)
			}

			for n, x := range maskIndex {
				// clear the nth bit then set to 1 or 0
				val = setBitNToX(val, n, x)
			}

			memory[memAddr] = val
		}
	}
	sum := 0
	for _, val := range memory {
		sum += val
	}
	return sum
}

func setBitNToX(val, n, x int) int {
	return (val & (^(1 << n))) | (x << n)
}
