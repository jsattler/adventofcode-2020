package main

import (
	"fmt"
	"github.com/j-sattler/adventofcode-2020/internal/aocio"
	"log"
	"strconv"
)

func Day01Part2() {

	strArr := aocio.ReadFileSplit("../../assets/input-01", "\n")
	// iterate over array
	for i, s1 := range strArr[:len(strArr)-1] {

		first, err := strconv.Atoi(s1)

		if err != nil {
			log.Fatal(err)
		}
		if first > 2020 {
			continue
		}
		// iterate over array start from i + 1
		for j, s2 := range strArr[i+1 : len(strArr)-1] {

			second, err := strconv.Atoi(s2)

			if err != nil {
				log.Fatal(err)
			}

			if first+second > 2020 {
				continue
			}
			// iterate over array start from j + 1
			for _, s3 := range strArr[j+1:] {
				third, err := strconv.Atoi(s3)

				if err != nil {
					log.Fatal(err)
				}

				if first+second+third == 2020 {
					fmt.Printf("solution day 1 part 2: %d\n", first*second*third)
					return
				}
			}
		}
	}
	fmt.Println("Did not find matching values")
}
