package aoc

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Day1Part1(){
	content, err := ioutil.ReadFile("assets/input-1")

	if err != nil { log.Fatal(err) }

	// split written values into string array
	strArr := strings.Split(string(content), "\n")

	// iterate over array
	for i, s1 := range strArr[:len(strArr)-1] {

		first, err := strconv.Atoi(s1)

		if err != nil { log.Fatal(err) }

		// iterate over array start from i + 1
		for _, s2 := range strArr[i+1:] {

			second, err := strconv.Atoi(s2)

			if err != nil { log.Fatal(err) }

			if first + second == 2020 {
				fmt.Printf("Found matching values (%d, %d) product: %d\n", first, second, first * second)
				return
			}
		}
	}
	fmt.Println("Did not find matching values")
}
