package main

import (
	"fmt"
	"github.com/j-sattler/adventofcode-2020/internal/aocio"
	"log"
	"strconv"
)

func Day01Part1() {
	strArr := aocio.ReadFileSplit("../../assets/input-01", "\n")
	// iterate over array
	for i, s1 := range strArr[:len(strArr)-1] {

		first, err := strconv.Atoi(s1)

		if err != nil {
			log.Fatal(err)
		}

		// iterate over array start from i + 1
		for _, s2 := range strArr[i+1:] {

			second, err := strconv.Atoi(s2)

			if err != nil {
				log.Fatal(err)
			}

			if first+second == 2020 {
				fmt.Printf("solution day 1 part 1: %d\n", first*second)
				return
			}
		}
	}
	fmt.Println("Did not find matching values")
}
