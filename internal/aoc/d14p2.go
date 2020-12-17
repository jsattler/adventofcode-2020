package aoc

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Day14Part2() {
	input, err := ioutil.ReadFile("assets/input-14")

	if err != nil {
		log.Fatal(err)
	}

	splitByLine := strings.Split(string(input), "\n")

	solution := inspectProgramV2(splitByLine)
	fmt.Printf("Solution day 14 part 2: %d\n", solution)

}

func inspectProgramV2(input []string) int {
	memory := make(map[int]int)
	floats := make([]int, 0)
	maskBits := 0
	for _, line := range input {

		if strings.HasPrefix(line, "mask") {
			maskBits = 0
			floats = make([]int, 0)
			maskStr := strings.ReplaceAll(line, "mask = ", "")
			index := 0
			for i := len(maskStr) - 1; i >= 0; i-- {
				if maskStr[i] == 'X' {
					// save the index of all floats
					floats = append(floats, index)
					index++
					continue
				}

				val, err := strconv.Atoi(string(maskStr[i]))
				if err != nil {
					log.Fatal(err)
				}
				maskBits = maskBits | (val << index)
				index++
			}

		} else {
			memAddr := 0
			val := 0
			_, err := fmt.Sscanf(line, "mem[%d] = %d", &memAddr, &val)
			if err != nil {
				log.Fatal(err)
			}
			memAddrs := maskAddress(memAddr, floats, maskBits)
			for addr, _ := range memAddrs {
				memory[addr] = val
			}
		}
	}
	sum := 0
	for _, val := range memory {
		sum += val
	}
	return sum
}

func maskAddress(addr int, floats []int, maskBits int) map[int]bool {
	addrs := make(map[int]bool)

	addr |= maskBits

	for _, fi := range floats {
		mask := ^(1 << fi)
		addr &= mask
	}

	addrs[addr] = true
	for _, floatIndex := range floats {
		for ad, _ := range addrs {
			ad = ad | (1 << floatIndex)
			addrs[ad] = true
		}
	}
	return addrs
}
