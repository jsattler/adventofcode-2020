package main

import (
	"fmt"
	"github.com/j-sattler/adventofcode-2020/internal/aocio"
	"github.com/j-sattler/adventofcode-2020/internal/aocparse"
	"log"
	"strings"
)

type Interval struct {
	low  int
	high int
}

type FieldPolicy struct {
	field     string
	intervals []Interval
}

type Ticket []int

func Day16Part1() {
	input := aocio.ReadFileSplit("../../assets/input-16", "\n\n")
	policies := parseFieldPolicies(input[0])
	nearByTickets := parseTickets(input[2])
	invalid := findInvalidValues(nearByTickets, policies)
	sum := 0
	for _, el := range invalid {
		sum += el
	}
	fmt.Printf("Solution day 16 part 1: %d\n", sum)
}

func isBetween(interval Interval, value int) bool {
	return value >= interval.low && value <= interval.high
}
func findInvalidValues(tickets []Ticket, policies []FieldPolicy) (invalid []int) {
	for _, t := range tickets {
		for _, v := range t {
			isValid := false
			for _, p := range policies {
				if isBetween(p.intervals[0], v) || isBetween(p.intervals[1], v) {
					isValid = true
					break
				}
			}
			if !isValid {
				invalid = append(invalid, v)
			}
		}
	}
	return
}

func parseTickets(in string) (tickets []Ticket) {
	lines := strings.Split(in, "\n")
	for _, l := range lines[1:] {
		if l == "" {
			continue
		}
		values := aocparse.SplitInts(l, ",")
		tickets = append(tickets, values)
	}
	return
}

func parseFieldPolicies(in string) (policies []FieldPolicy) {
	lines := strings.Split(in, "\n")
	for _, l := range lines {
		var firstLow, firstHigh, secondLow, secondHigh int
		splitLine := strings.Split(l, ": ")
		l = strings.ReplaceAll(splitLine[1], "-", " ")
		_, err := fmt.Sscanf(l, "%d %d or %d %d", &firstLow, &firstHigh, &secondLow, &secondHigh)
		if err != nil {
			log.Fatal(err)
		}

		policies = append(policies, FieldPolicy{
			field: splitLine[0],
			intervals: []Interval{
				Interval{low: firstLow, high: firstHigh},
				Interval{low: secondLow, high: secondHigh},
			},
		})
	}
	return
}
