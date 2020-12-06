package aoc

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func Day6Part2() {
	var counter int
	content, err := ioutil.ReadFile("assets/input-6")
	if err != nil {
		log.Fatal(err)
	}

	groups:= strings.Split(string(content), "\n\n")

	for _, group := range groups {
		chars := make(map[rune]int)
		groupSize :=  strings.Count(group, "\n") + 1
		group := strings.ReplaceAll(group, "\n", "")
		fmt.Println(group, groupSize)
		// iterate over each rune within the group
		for _, char := range group {
			chars[char]++
			fmt.Printf("Found entries %d for rune %s\n", chars[char], string(char))
		}

		for _, v := range chars {
			if v == groupSize{
				counter++
			}
		}
	}


	fmt.Printf("solution day 6 part 2: %d\n", counter)

}

