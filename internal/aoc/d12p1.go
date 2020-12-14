package aoc

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const (
	moveNorth   = 'N'
	moveSouth   = 'S'
	moveEast    = 'E'
	moveWest    = 'W'
	turnLeft    = 'L'
	turnRight   = 'R'
	moveForward = 'F'

	maxDegrees = 360
	north      = 0
	east       = 90
	south      = 180
	west       = 270
)

type Ship struct {
	currentDirection int
	southNorth       int // south is negative north is positive
	westEast         int // west is negative east is positive
	waypoint         Waypoint
}

type Action struct {
	action int32
	value  int
}

func Day12Part1() {

	input, err := ioutil.ReadFile("assets/input-12")
	if err != nil {
		log.Fatal(err)
	}
	inputStr := strings.Split(string(input), "\n")
	actions := parseActions(inputStr)

	ship := Ship{currentDirection: east}

	for _, action := range actions {
		ship.execute(action)
	}

	fmt.Printf("Solution day 12 part 1: %d\n", intAbs(ship.westEast)+intAbs(ship.southNorth))
}

func (s *Ship) execute(a Action) {
	switch a.action {
	case moveNorth:
		s.southNorth += a.value
	case moveEast:
		s.westEast += a.value
	case moveSouth:
		s.southNorth -= a.value
	case moveWest:
		s.westEast -= a.value
	case turnLeft:
		s.currentDirection = mod(s.currentDirection-a.value, maxDegrees)
	case turnRight:
		s.currentDirection = mod(s.currentDirection+a.value, maxDegrees)
	case moveForward:
		switch s.currentDirection {
		case north:
			s.southNorth += a.value
		case east:
			s.westEast += a.value
		case south:
			s.southNorth -= a.value
		case west:
			s.westEast -= a.value
		}
	}
}

func parseActions(inputStr []string) []Action {
	actions := make([]Action, len(inputStr))

	for i, action := range inputStr {
		a := Action{}
		a.action = rune(action[0])
		val, err := strconv.Atoi(action[1:])

		if err != nil {
			log.Fatal(err)
		}

		a.value = val
		actions[i] = a
	}
	return actions
}

func intAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func mod(a, b int) int {
	return (a%b + b) % b
}
