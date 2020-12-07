package aoc

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const shinyGold = "shiny gold"

var counter = 0

func Day7Part1() {

	fmt.Println(counter)
	content, err := ioutil.ReadFile("assets/input-7")
	if err != nil {
		log.Fatal(err)
	}

	rules := strings.Split(string(content), "\n")

	res := make([]string, 0)
	res = findValidBags(shinyGold, rules, res)

	// we have collected all rules that contain shiny gold but do not start with it.
	// now we have to look if the parent bag is contained

	fmt.Println(len(res))

}

// 1. Check what lines contain bag as a child
// 2. Take the parent bag of each line that was found and repeat the process
// 3. Repeat until ..

func findValidBags(valid string, rules []string, res []string) []string {

	fmt.Println("checking for", valid)
	// iterate over all bag rules
	for _, rule := range rules {
		fmt.Println("checking rule", rule)
		isChecked := false

		for _, r := range res {
			if strings.HasPrefix(rule, r) {
				isChecked = true
			}
		}

		if isChecked {
			continue
		}

		// check if rule contains 'shiny gold' and does not start with shiny gold
		if strings.Contains(rule, valid) && !strings.HasPrefix(rule, valid) {
			// add bag to the valid bags
			parent := strings.Split(rule, " bags contain ")[0]
			fmt.Println("Found bag that contains shiny gold", parent)
			res = append(findValidBags(parent, rules, res), parent)
		}
	}
	return res

}
