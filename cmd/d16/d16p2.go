package main

import (
	"fmt"
	"github.com/j-sattler/adventofcode-2020/internal/aocio"
	"strings"
)

func Day16Part2() {
	input := aocio.ReadFileSplit("../../assets/input-16", "\n\n")
	policies := parseFieldPolicies(input[0])
	myTicket := parseTickets(input[1])
	nearByTickets := parseTickets(input[2])
	validTickets := discardInvalid(nearByTickets, policies)
	res := determineFields(validTickets, policies)
	fmt.Println(res)
	product := 1
	for k, v := range res {
		if strings.Contains(k, "departure") {
			product *= myTicket[0][v]
		}
	}
	fmt.Printf("Solution day 16 part 2: %d\n", product)
}

func determineFields(validTickets []Ticket, policies []FieldPolicy) (result map[string]int) {
	fieldLen := len(validTickets[0])
	result = make(map[string]int)
	validityMatrix := make(map[int]map[string]bool)
	remainingPolicies := make(map[string]bool)
	// first step is to collect all possible fields for one position
	for i := 0; i < fieldLen; i++ {
		validityMatrix[i] = make(map[string]bool)
		for _, p := range policies {
			remainingPolicies[p.field] = true
			isValid := true
			for _, t := range validTickets {
				if !isBetween(p.intervals[0], t[i]) && !isBetween(p.intervals[1], t[i]) {
					isValid = false
				}
			}
			if isValid {
				validityMatrix[i][p.field] = isValid
			}
		}
	}
	// second step is to determine which position fits which field
	// remove a field from all positions in the matrix, once it was assigned to a position
	for len(remainingPolicies) > 0 { // until all policies are assigned to fields
		found := false
		remove := ""
		for pos, validFields := range validityMatrix {
			for field := range validFields {
				if len(validFields) == 1 {
					found = true
					remove = field
					result[field] = pos
					delete(remainingPolicies, field)
					break
				}
			}
			if found {
				break
			}
		}
		if found {
			for _, validFields := range validityMatrix {
				delete(validFields, remove)
			}
		}
	}
	return
}

func discardInvalid(tickets []Ticket, policies []FieldPolicy) (valid []Ticket) {
	for _, t := range tickets {
		isValid := false
		for _, v := range t {
			isValid = false
			for _, p := range policies {
				if isBetween(p.intervals[0], v) || isBetween(p.intervals[1], v) {
					isValid = true
					break
				}
			}
			if !isValid {
				break
			}
		}
		if isValid {
			valid = append(valid, t)
		}
	}
	return
}
