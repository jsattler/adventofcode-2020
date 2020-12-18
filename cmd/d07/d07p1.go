package main

import (
	"fmt"
	"github.com/j-sattler/adventofcode-2020/internal/aocio"
	"strings"
)

const shinyGold = "shiny gold"

func Day07Part1() {
	rules := aocio.ReadFileSplit("../../assets/input-07", "\n")

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
