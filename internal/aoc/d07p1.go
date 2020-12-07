package aoc

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const shinyGold = "shiny gold"

func Day7Part1() {

	content, err := ioutil.ReadFile("assets/input-7")
	if err != nil {
		log.Fatal(err)
	}

	rules := strings.Split(string(content), "\n")

	res := make([]string, 0)
	res = findValidBags(shinyGold, rules, res)

	fmt.Printf("Solution day 7 part 1: %d\n", len(res))

}

func findValidBags(valid string, rules []string, res []string) []string {

	// iterate over all bag rules
	for _, rule := range rules {
		isChecked := false

		for _, r := range res {
			if strings.HasPrefix(rule, r) {
				isChecked = true
			}
		}

		if isChecked {
			continue
		}

		if strings.Contains(rule, valid) && !strings.HasPrefix(rule, valid) {
			// add bag to the valid bags
			parent := strings.Split(rule, " bags contain ")[0]
			res = append(findValidBags(parent, rules, res), parent)
		}
	}
	return res

}
